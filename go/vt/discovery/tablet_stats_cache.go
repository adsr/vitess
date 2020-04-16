/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package discovery

import (
	"sync"

	"golang.org/x/net/context"
	"vitess.io/vitess/go/vt/log"
	querypb "vitess.io/vitess/go/vt/proto/query"
	topodatapb "vitess.io/vitess/go/vt/proto/topodata"
	"vitess.io/vitess/go/vt/topo"
	"vitess.io/vitess/go/vt/topo/topoproto"
)

// TabletStatsCache is a HealthCheckStatsListener that keeps both the
// current list of available TabletStats, and a serving list:
// - for master tablets, only the current master is kept.
// - for non-master tablets, we filter the list using FilterStatsByReplicationLag.
// It keeps entries for all tablets in the cell(s) it's configured to serve for,
// and for the master independently of which cell it's in.
// Note the healthy tablet computation is done when we receive a tablet
// update only, not at serving time.
// Also note the cache may not have the last entry received by the tablet.
// For instance, if a tablet was healthy, and is still healthy, we do not
// keep its new update.
type TabletStatsCache struct {
	// cell is the cell we are keeping all tablets for.
	// Note we keep track of all master tablets in all cells.
	cell string
	// mu protects the following fields. It does not protect individual
	// entries in the entries map.
	mu sync.RWMutex
	// entries maps from keyspace/shard/tabletType to our cache.
	entries map[string]map[string]map[topodatapb.TabletType]*tabletStatsCacheEntry
	// cellAliases is a cache of cell aliases
	cellAliases map[string]string
}

// tabletStatsCacheEntry is the per keyspace/shard/tabletType
// entry of the in-memory map for TabletStatsCache.
type tabletStatsCacheEntry struct {
	// mu protects the rest of this structure.
	mu sync.RWMutex
	// all has the valid tablets, indexed by TabletToMapKey(ts.Tablet),
	// as it is the index used by HealthCheck.
	all map[string]*TabletStats
	// healthy only has the healthy ones.
	healthy []*TabletStats
}

func (e *tabletStatsCacheEntry) updateHealthyMapForMaster(ts *TabletStats) {
	if ts.Up {
		// We have an Up master.
		if len(e.healthy) == 0 {
			// We have a new Up server, just remember it.
			e.healthy = append(e.healthy, ts)
			return
		}

		// We already have one up server, see if we
		// need to replace it.
		if ts.MasterTermStartTime < e.healthy[0].MasterTermStartTime {
			log.Warningf("not marking healthy master %s as Up for %s because its externally reparented timestamp is smaller than the highest known timestamp from previous MASTERs %s: %d < %d ",
				topoproto.TabletAliasString(ts.Tablet.Alias),
				topoproto.KeyspaceShardString(ts.Target.Keyspace, ts.Target.Shard),
				topoproto.TabletAliasString(e.healthy[0].Tablet.Alias),
				ts.MasterTermStartTime,
				e.healthy[0].MasterTermStartTime)
			return
		}

		// Just replace it.
		e.healthy[0] = ts
		return
	}

	// We have a Down master, remove it only if it's exactly the same.
	if len(e.healthy) != 0 {
		if ts.Key == e.healthy[0].Key {
			// Same guy, remove it.
			e.healthy = nil
		}
	}
}

// NewTabletStatsCache creates a TabletStatsCache, and registers
// it as HealthCheckStatsListener of the provided healthcheck.
// Note we do the registration in this code to guarantee we call
// SetListener with sendDownEvents=true, as we need these events
// to maintain the integrity of our cache.
func NewTabletStatsCache(hc HealthCheck, ts *topo.Server, cell string) *TabletStatsCache {
	return newTabletStatsCache(cell)
}

func newTabletStatsCache(localCell string) *TabletStatsCache {
	tc := &TabletStatsCache{
		entries:     make(map[string]map[string]map[topodatapb.TabletType]*tabletStatsCacheEntry),
		cell:        localCell,
		cellAliases: make(map[string]string),
	}
	return tc
}

// UpdateStats is used to ...
func (tc *TabletStatsCache) UpdateStats(ts *TabletStats, topoServer *topo.Server) {
	if ts.Target.TabletType != topodatapb.TabletType_MASTER &&
		ts.Tablet.Alias.Cell != tc.cell &&
		tc.getAliasByCell(ts.Tablet.Alias.Cell, topoServer) != tc.getAliasByCell(tc.cell, topoServer) {
		// this is for a non-master tablet in a different cell and a different alias, drop it
		return
	}

	e := tc.getOrCreateEntry(ts.Target)
	e.mu.Lock()
	defer e.mu.Unlock()

	// Update our full map.
	trivialNonMasterUpdate := false
	if existing, ok := e.all[ts.Key]; ok {
		if ts.Up {
			// We have an existing entry, and a new entry.
			// Remember if they are both good (most common case).
			trivialNonMasterUpdate = existing.LastError == nil && existing.Serving && ts.LastError == nil &&
				ts.Serving && ts.Target.TabletType != topodatapb.TabletType_MASTER && existing.TrivialStatsUpdate(ts)

			// We already have the entry, update the
			// values if necessary.  (will update both
			// 'all' and 'healthy' as they use pointers).
			if !trivialNonMasterUpdate {
				*existing = *ts
			}
		} else {
			// We have an entry which we shouldn't. Remove it.
			delete(e.all, ts.Key)
		}
	} else {
		if ts.Up {
			// Add the entry.
			e.all[ts.Key] = ts
		} else {
			// We were told to remove an entry which we
			// didn't have anyway, nothing should happen.
			return
		}
	}

	// Update our healthy list.
	var allArray []*TabletStats
	if ts.Target.TabletType == topodatapb.TabletType_MASTER {
		// The healthy list is different for TabletType_MASTER: we
		// only keep the most recent one.
		e.updateHealthyMapForMaster(ts)
	} else {
		// For non-master, if it is a trivial update,
		// we just skip everything else. We don't even update the
		// aggregate stats.
		if trivialNonMasterUpdate {
			return
		}

		// Now we need to do some work. Recompute our healthy list.
		allArray = make([]*TabletStats, 0, len(e.all))
		for _, s := range e.all {
			allArray = append(allArray, s)
		}
		e.healthy = FilterStatsByReplicationLag(allArray)
	}
}

// getEntry returns an existing TabletStatsCacheEntry in the cache, or nil
// if the entry does not exist. It only takes a Read lock on mu.
func (tc *TabletStatsCache) getEntry(keyspace, shard string, tabletType topodatapb.TabletType) *tabletStatsCacheEntry {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	if s, ok := tc.entries[keyspace]; ok {
		if t, ok := s[shard]; ok {
			if e, ok := t[tabletType]; ok {
				return e
			}
		}
	}
	return nil
}

// getOrCreateEntry returns an existing TabletStatsCacheEntry from the cache,
// or creates it if it doesn't exist.
func (tc *TabletStatsCache) getOrCreateEntry(target *querypb.Target) *tabletStatsCacheEntry {
	// Fast path (most common path too): Read-lock, return the entry.
	if e := tc.getEntry(target.Keyspace, target.Shard, target.TabletType); e != nil {
		return e
	}

	// Slow path: Lock, will probably have to add the entry at some level.
	tc.mu.Lock()
	defer tc.mu.Unlock()

	s, ok := tc.entries[target.Keyspace]
	if !ok {
		s = make(map[string]map[topodatapb.TabletType]*tabletStatsCacheEntry)
		tc.entries[target.Keyspace] = s
	}
	t, ok := s[target.Shard]
	if !ok {
		t = make(map[topodatapb.TabletType]*tabletStatsCacheEntry)
		s[target.Shard] = t
	}
	e, ok := t[target.TabletType]
	if !ok {
		e = &tabletStatsCacheEntry{
			all: make(map[string]*TabletStats),
		}
		t[target.TabletType] = e
	}
	return e
}

func (tc *TabletStatsCache) getAliasByCell(cell string, topoServer *topo.Server) string {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	if alias, ok := tc.cellAliases[cell]; ok {
		return alias
	}

	alias := topo.GetAliasByCell(context.Background(), topoServer, cell)
	tc.cellAliases[cell] = alias

	return alias
}

// GetTabletStats returns the full list of available targets.
// The returned array is owned by the caller.
func (tc *TabletStatsCache) GetTabletStats(keyspace, shard string, tabletType topodatapb.TabletType) []TabletStats {
	e := tc.getEntry(keyspace, shard, tabletType)
	if e == nil {
		return nil
	}

	e.mu.RLock()
	defer e.mu.RUnlock()
	// ok to make a copy here
	result := make([]TabletStats, 0, len(e.all))
	for _, s := range e.all {
		result = append(result, *s)
	}
	return result
}

// GetHealthyTabletStats returns only the healthy targets.
// The returned array is owned by the caller.
// For TabletType_MASTER, this will only return at most one entry,
// the most recent tablet of type master.
func (tc *TabletStatsCache) GetHealthyTabletStats(keyspace, shard string, tabletType topodatapb.TabletType) []TabletStats {
	e := tc.getEntry(keyspace, shard, tabletType)
	if e == nil {
		return nil
	}

	e.mu.RLock()
	defer e.mu.RUnlock()
	// ok to make a copy here
	result := make([]TabletStats, len(e.healthy))
	for i, ts := range e.healthy {
		result[i] = *ts
	}
	return result
}

// ResetForTesting is for use in tests only.
func (tc *TabletStatsCache) ResetForTesting() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	tc.entries = make(map[string]map[string]map[topodatapb.TabletType]*tabletStatsCacheEntry)
}
