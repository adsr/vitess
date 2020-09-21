// Code generated by visitorgen/main/main.go. DO NOT EDIT.

package sqlparser

//go:generate go run ./visitorgen/main -input=ast.go -output=rewriter.go

import (
	"reflect"
)

type replacerFunc func(newNode, parent SQLNode)

// application carries all the shared data so we can pass it around cheaply.
type application struct {
	pre, post ApplyFunc
	cursor    Cursor
}

func replaceAliasedExprAs(newNode, parent SQLNode) {
	parent.(*AliasedExpr).As = newNode.(ColIdent)
}

func replaceAliasedExprExpr(newNode, parent SQLNode) {
	parent.(*AliasedExpr).Expr = newNode.(Expr)
}

func replaceAliasedTableExprAs(newNode, parent SQLNode) {
	parent.(*AliasedTableExpr).As = newNode.(TableIdent)
}

func replaceAliasedTableExprExpr(newNode, parent SQLNode) {
	parent.(*AliasedTableExpr).Expr = newNode.(SimpleTableExpr)
}

func replaceAliasedTableExprHints(newNode, parent SQLNode) {
	parent.(*AliasedTableExpr).Hints = newNode.(*IndexHints)
}

func replaceAliasedTableExprPartitions(newNode, parent SQLNode) {
	parent.(*AliasedTableExpr).Partitions = newNode.(Partitions)
}

func replaceAndExprLeft(newNode, parent SQLNode) {
	parent.(*AndExpr).Left = newNode.(Expr)
}

func replaceAndExprRight(newNode, parent SQLNode) {
	parent.(*AndExpr).Right = newNode.(Expr)
}

func replaceAutoIncSpecColumn(newNode, parent SQLNode) {
	parent.(*AutoIncSpec).Column = newNode.(ColIdent)
}

func replaceAutoIncSpecSequence(newNode, parent SQLNode) {
	parent.(*AutoIncSpec).Sequence = newNode.(TableName)
}

func replaceBinaryExprLeft(newNode, parent SQLNode) {
	parent.(*BinaryExpr).Left = newNode.(Expr)
}

func replaceBinaryExprRight(newNode, parent SQLNode) {
	parent.(*BinaryExpr).Right = newNode.(Expr)
}

func replaceCaseExprElse(newNode, parent SQLNode) {
	parent.(*CaseExpr).Else = newNode.(Expr)
}

func replaceCaseExprExpr(newNode, parent SQLNode) {
	parent.(*CaseExpr).Expr = newNode.(Expr)
}

type replaceCaseExprWhens int

func (r *replaceCaseExprWhens) replace(newNode, container SQLNode) {
	container.(*CaseExpr).Whens[int(*r)] = newNode.(*When)
}

func (r *replaceCaseExprWhens) inc() {
	*r++
}

func replaceColNameName(newNode, parent SQLNode) {
	parent.(*ColName).Name = newNode.(ColIdent)
}

func replaceColNameQualifier(newNode, parent SQLNode) {
	parent.(*ColName).Qualifier = newNode.(TableName)
}

func replaceCollateExprExpr(newNode, parent SQLNode) {
	parent.(*CollateExpr).Expr = newNode.(Expr)
}

func replaceColumnDefinitionName(newNode, parent SQLNode) {
	parent.(*ColumnDefinition).Name = newNode.(ColIdent)
}

func replaceColumnTypeAutoincrement(newNode, parent SQLNode) {
	parent.(*ColumnType).Autoincrement = newNode.(BoolVal)
}

func replaceColumnTypeComment(newNode, parent SQLNode) {
	parent.(*ColumnType).Comment = newNode.(*Literal)
}

func replaceColumnTypeDefault(newNode, parent SQLNode) {
	parent.(*ColumnType).Default = newNode.(Expr)
}

func replaceColumnTypeLength(newNode, parent SQLNode) {
	parent.(*ColumnType).Length = newNode.(*Literal)
}

func replaceColumnTypeNotNull(newNode, parent SQLNode) {
	parent.(*ColumnType).NotNull = newNode.(BoolVal)
}

func replaceColumnTypeOnUpdate(newNode, parent SQLNode) {
	parent.(*ColumnType).OnUpdate = newNode.(Expr)
}

func replaceColumnTypeScale(newNode, parent SQLNode) {
	parent.(*ColumnType).Scale = newNode.(*Literal)
}

func replaceColumnTypeUnsigned(newNode, parent SQLNode) {
	parent.(*ColumnType).Unsigned = newNode.(BoolVal)
}

func replaceColumnTypeZerofill(newNode, parent SQLNode) {
	parent.(*ColumnType).Zerofill = newNode.(BoolVal)
}

type replaceColumnsItems int

func (r *replaceColumnsItems) replace(newNode, container SQLNode) {
	container.(Columns)[int(*r)] = newNode.(ColIdent)
}

func (r *replaceColumnsItems) inc() {
	*r++
}

func replaceComparisonExprEscape(newNode, parent SQLNode) {
	parent.(*ComparisonExpr).Escape = newNode.(Expr)
}

func replaceComparisonExprLeft(newNode, parent SQLNode) {
	parent.(*ComparisonExpr).Left = newNode.(Expr)
}

func replaceComparisonExprRight(newNode, parent SQLNode) {
	parent.(*ComparisonExpr).Right = newNode.(Expr)
}

func replaceConstraintDefinitionDetails(newNode, parent SQLNode) {
	parent.(*ConstraintDefinition).Details = newNode.(ConstraintInfo)
}

func replaceConvertExprExpr(newNode, parent SQLNode) {
	parent.(*ConvertExpr).Expr = newNode.(Expr)
}

func replaceConvertExprType(newNode, parent SQLNode) {
	parent.(*ConvertExpr).Type = newNode.(*ConvertType)
}

func replaceConvertTypeLength(newNode, parent SQLNode) {
	parent.(*ConvertType).Length = newNode.(*Literal)
}

func replaceConvertTypeScale(newNode, parent SQLNode) {
	parent.(*ConvertType).Scale = newNode.(*Literal)
}

func replaceConvertUsingExprExpr(newNode, parent SQLNode) {
	parent.(*ConvertUsingExpr).Expr = newNode.(Expr)
}

func replaceCurTimeFuncExprFsp(newNode, parent SQLNode) {
	parent.(*CurTimeFuncExpr).Fsp = newNode.(Expr)
}

func replaceCurTimeFuncExprName(newNode, parent SQLNode) {
	parent.(*CurTimeFuncExpr).Name = newNode.(ColIdent)
}

func replaceDDLAutoIncSpec(newNode, parent SQLNode) {
	parent.(*DDL).AutoIncSpec = newNode.(*AutoIncSpec)
}

func replaceDDLFromTables(newNode, parent SQLNode) {
	parent.(*DDL).FromTables = newNode.(TableNames)
}

func replaceDDLOptLike(newNode, parent SQLNode) {
	parent.(*DDL).OptLike = newNode.(*OptLike)
}

func replaceDDLPartitionSpec(newNode, parent SQLNode) {
	parent.(*DDL).PartitionSpec = newNode.(*PartitionSpec)
}

func replaceDDLTable(newNode, parent SQLNode) {
	parent.(*DDL).Table = newNode.(TableName)
}

func replaceDDLTableSpec(newNode, parent SQLNode) {
	parent.(*DDL).TableSpec = newNode.(*TableSpec)
}

func replaceDDLToTables(newNode, parent SQLNode) {
	parent.(*DDL).ToTables = newNode.(TableNames)
}

type replaceDDLVindexCols int

func (r *replaceDDLVindexCols) replace(newNode, container SQLNode) {
	container.(*DDL).VindexCols[int(*r)] = newNode.(ColIdent)
}

func (r *replaceDDLVindexCols) inc() {
	*r++
}

func replaceDDLVindexSpec(newNode, parent SQLNode) {
	parent.(*DDL).VindexSpec = newNode.(*VindexSpec)
}

func replaceDeleteComments(newNode, parent SQLNode) {
	parent.(*Delete).Comments = newNode.(Comments)
}

func replaceDeleteLimit(newNode, parent SQLNode) {
	parent.(*Delete).Limit = newNode.(*Limit)
}

func replaceDeleteOrderBy(newNode, parent SQLNode) {
	parent.(*Delete).OrderBy = newNode.(OrderBy)
}

func replaceDeletePartitions(newNode, parent SQLNode) {
	parent.(*Delete).Partitions = newNode.(Partitions)
}

func replaceDeleteTableExprs(newNode, parent SQLNode) {
	parent.(*Delete).TableExprs = newNode.(TableExprs)
}

func replaceDeleteTargets(newNode, parent SQLNode) {
	parent.(*Delete).Targets = newNode.(TableNames)
}

func replaceDeleteWhere(newNode, parent SQLNode) {
	parent.(*Delete).Where = newNode.(*Where)
}

func replaceExistsExprSubquery(newNode, parent SQLNode) {
	parent.(*ExistsExpr).Subquery = newNode.(*Subquery)
}

func replaceExplainStatement(newNode, parent SQLNode) {
	parent.(*Explain).Statement = newNode.(Statement)
}

type replaceExprsItems int

func (r *replaceExprsItems) replace(newNode, container SQLNode) {
	container.(Exprs)[int(*r)] = newNode.(Expr)
}

func (r *replaceExprsItems) inc() {
	*r++
}

func replaceForeignKeyDefinitionOnDelete(newNode, parent SQLNode) {
	parent.(*ForeignKeyDefinition).OnDelete = newNode.(ReferenceAction)
}

func replaceForeignKeyDefinitionOnUpdate(newNode, parent SQLNode) {
	parent.(*ForeignKeyDefinition).OnUpdate = newNode.(ReferenceAction)
}

func replaceForeignKeyDefinitionReferencedColumns(newNode, parent SQLNode) {
	parent.(*ForeignKeyDefinition).ReferencedColumns = newNode.(Columns)
}

func replaceForeignKeyDefinitionReferencedTable(newNode, parent SQLNode) {
	parent.(*ForeignKeyDefinition).ReferencedTable = newNode.(TableName)
}

func replaceForeignKeyDefinitionSource(newNode, parent SQLNode) {
	parent.(*ForeignKeyDefinition).Source = newNode.(Columns)
}

func replaceFuncExprExprs(newNode, parent SQLNode) {
	parent.(*FuncExpr).Exprs = newNode.(SelectExprs)
}

func replaceFuncExprName(newNode, parent SQLNode) {
	parent.(*FuncExpr).Name = newNode.(ColIdent)
}

func replaceFuncExprQualifier(newNode, parent SQLNode) {
	parent.(*FuncExpr).Qualifier = newNode.(TableIdent)
}

type replaceGroupByItems int

func (r *replaceGroupByItems) replace(newNode, container SQLNode) {
	container.(GroupBy)[int(*r)] = newNode.(Expr)
}

func (r *replaceGroupByItems) inc() {
	*r++
}

func replaceGroupConcatExprExprs(newNode, parent SQLNode) {
	parent.(*GroupConcatExpr).Exprs = newNode.(SelectExprs)
}

func replaceGroupConcatExprLimit(newNode, parent SQLNode) {
	parent.(*GroupConcatExpr).Limit = newNode.(*Limit)
}

func replaceGroupConcatExprOrderBy(newNode, parent SQLNode) {
	parent.(*GroupConcatExpr).OrderBy = newNode.(OrderBy)
}

func replaceIndexDefinitionInfo(newNode, parent SQLNode) {
	parent.(*IndexDefinition).Info = newNode.(*IndexInfo)
}

type replaceIndexHintsIndexes int

func (r *replaceIndexHintsIndexes) replace(newNode, container SQLNode) {
	container.(*IndexHints).Indexes[int(*r)] = newNode.(ColIdent)
}

func (r *replaceIndexHintsIndexes) inc() {
	*r++
}

func replaceIndexInfoName(newNode, parent SQLNode) {
	parent.(*IndexInfo).Name = newNode.(ColIdent)
}

func replaceInsertColumns(newNode, parent SQLNode) {
	parent.(*Insert).Columns = newNode.(Columns)
}

func replaceInsertComments(newNode, parent SQLNode) {
	parent.(*Insert).Comments = newNode.(Comments)
}

func replaceInsertOnDup(newNode, parent SQLNode) {
	parent.(*Insert).OnDup = newNode.(OnDup)
}

func replaceInsertPartitions(newNode, parent SQLNode) {
	parent.(*Insert).Partitions = newNode.(Partitions)
}

func replaceInsertRows(newNode, parent SQLNode) {
	parent.(*Insert).Rows = newNode.(InsertRows)
}

func replaceInsertTable(newNode, parent SQLNode) {
	parent.(*Insert).Table = newNode.(TableName)
}

func replaceIntervalExprExpr(newNode, parent SQLNode) {
	parent.(*IntervalExpr).Expr = newNode.(Expr)
}

func replaceIsExprExpr(newNode, parent SQLNode) {
	parent.(*IsExpr).Expr = newNode.(Expr)
}

func replaceJoinConditionOn(newNode, parent SQLNode) {
	tmp := parent.(JoinCondition)
	tmp.On = newNode.(Expr)
}

func replaceJoinConditionUsing(newNode, parent SQLNode) {
	tmp := parent.(JoinCondition)
	tmp.Using = newNode.(Columns)
}

func replaceJoinTableExprCondition(newNode, parent SQLNode) {
	parent.(*JoinTableExpr).Condition = newNode.(JoinCondition)
}

func replaceJoinTableExprLeftExpr(newNode, parent SQLNode) {
	parent.(*JoinTableExpr).LeftExpr = newNode.(TableExpr)
}

func replaceJoinTableExprRightExpr(newNode, parent SQLNode) {
	parent.(*JoinTableExpr).RightExpr = newNode.(TableExpr)
}

func replaceLimitOffset(newNode, parent SQLNode) {
	parent.(*Limit).Offset = newNode.(Expr)
}

func replaceLimitRowcount(newNode, parent SQLNode) {
	parent.(*Limit).Rowcount = newNode.(Expr)
}

func replaceMatchExprColumns(newNode, parent SQLNode) {
	parent.(*MatchExpr).Columns = newNode.(SelectExprs)
}

func replaceMatchExprExpr(newNode, parent SQLNode) {
	parent.(*MatchExpr).Expr = newNode.(Expr)
}

func replaceNextvalExpr(newNode, parent SQLNode) {
	tmp := parent.(Nextval)
	tmp.Expr = newNode.(Expr)
}

func replaceNotExprExpr(newNode, parent SQLNode) {
	parent.(*NotExpr).Expr = newNode.(Expr)
}

type replaceOnDupItems int

func (r *replaceOnDupItems) replace(newNode, container SQLNode) {
	container.(OnDup)[int(*r)] = newNode.(*UpdateExpr)
}

func (r *replaceOnDupItems) inc() {
	*r++
}

func replaceOptLikeLikeTable(newNode, parent SQLNode) {
	parent.(*OptLike).LikeTable = newNode.(TableName)
}

func replaceOrExprLeft(newNode, parent SQLNode) {
	parent.(*OrExpr).Left = newNode.(Expr)
}

func replaceOrExprRight(newNode, parent SQLNode) {
	parent.(*OrExpr).Right = newNode.(Expr)
}

func replaceOrderExpr(newNode, parent SQLNode) {
	parent.(*Order).Expr = newNode.(Expr)
}

type replaceOrderByItems int

func (r *replaceOrderByItems) replace(newNode, container SQLNode) {
	container.(OrderBy)[int(*r)] = newNode.(*Order)
}

func (r *replaceOrderByItems) inc() {
	*r++
}

func replaceParenSelectSelect(newNode, parent SQLNode) {
	parent.(*ParenSelect).Select = newNode.(SelectStatement)
}

func replaceParenTableExprExprs(newNode, parent SQLNode) {
	parent.(*ParenTableExpr).Exprs = newNode.(TableExprs)
}

func replacePartitionDefinitionLimit(newNode, parent SQLNode) {
	parent.(*PartitionDefinition).Limit = newNode.(Expr)
}

func replacePartitionDefinitionName(newNode, parent SQLNode) {
	parent.(*PartitionDefinition).Name = newNode.(ColIdent)
}

type replacePartitionSpecDefinitions int

func (r *replacePartitionSpecDefinitions) replace(newNode, container SQLNode) {
	container.(*PartitionSpec).Definitions[int(*r)] = newNode.(*PartitionDefinition)
}

func (r *replacePartitionSpecDefinitions) inc() {
	*r++
}

func replacePartitionSpecName(newNode, parent SQLNode) {
	parent.(*PartitionSpec).Name = newNode.(ColIdent)
}

type replacePartitionsItems int

func (r *replacePartitionsItems) replace(newNode, container SQLNode) {
	container.(Partitions)[int(*r)] = newNode.(ColIdent)
}

func (r *replacePartitionsItems) inc() {
	*r++
}

func replaceRangeCondFrom(newNode, parent SQLNode) {
	parent.(*RangeCond).From = newNode.(Expr)
}

func replaceRangeCondLeft(newNode, parent SQLNode) {
	parent.(*RangeCond).Left = newNode.(Expr)
}

func replaceRangeCondTo(newNode, parent SQLNode) {
	parent.(*RangeCond).To = newNode.(Expr)
}

func replaceReleaseName(newNode, parent SQLNode) {
	parent.(*Release).Name = newNode.(ColIdent)
}

func replaceSRollbackName(newNode, parent SQLNode) {
	parent.(*SRollback).Name = newNode.(ColIdent)
}

func replaceSavepointName(newNode, parent SQLNode) {
	parent.(*Savepoint).Name = newNode.(ColIdent)
}

func replaceSelectComments(newNode, parent SQLNode) {
	parent.(*Select).Comments = newNode.(Comments)
}

func replaceSelectFrom(newNode, parent SQLNode) {
	parent.(*Select).From = newNode.(TableExprs)
}

func replaceSelectGroupBy(newNode, parent SQLNode) {
	parent.(*Select).GroupBy = newNode.(GroupBy)
}

func replaceSelectHaving(newNode, parent SQLNode) {
	parent.(*Select).Having = newNode.(*Where)
}

func replaceSelectLimit(newNode, parent SQLNode) {
	parent.(*Select).Limit = newNode.(*Limit)
}

func replaceSelectOrderBy(newNode, parent SQLNode) {
	parent.(*Select).OrderBy = newNode.(OrderBy)
}

func replaceSelectSelectExprs(newNode, parent SQLNode) {
	parent.(*Select).SelectExprs = newNode.(SelectExprs)
}

func replaceSelectWhere(newNode, parent SQLNode) {
	parent.(*Select).Where = newNode.(*Where)
}

type replaceSelectExprsItems int

func (r *replaceSelectExprsItems) replace(newNode, container SQLNode) {
	container.(SelectExprs)[int(*r)] = newNode.(SelectExpr)
}

func (r *replaceSelectExprsItems) inc() {
	*r++
}

func replaceSetComments(newNode, parent SQLNode) {
	parent.(*Set).Comments = newNode.(Comments)
}

func replaceSetExprs(newNode, parent SQLNode) {
	parent.(*Set).Exprs = newNode.(SetExprs)
}

func replaceSetExprExpr(newNode, parent SQLNode) {
	parent.(*SetExpr).Expr = newNode.(Expr)
}

func replaceSetExprName(newNode, parent SQLNode) {
	parent.(*SetExpr).Name = newNode.(ColIdent)
}

type replaceSetExprsItems int

func (r *replaceSetExprsItems) replace(newNode, container SQLNode) {
	container.(SetExprs)[int(*r)] = newNode.(*SetExpr)
}

func (r *replaceSetExprsItems) inc() {
	*r++
}

type replaceSetTransactionCharacteristics int

func (r *replaceSetTransactionCharacteristics) replace(newNode, container SQLNode) {
	container.(*SetTransaction).Characteristics[int(*r)] = newNode.(Characteristic)
}

func (r *replaceSetTransactionCharacteristics) inc() {
	*r++
}

func replaceSetTransactionComments(newNode, parent SQLNode) {
	parent.(*SetTransaction).Comments = newNode.(Comments)
}

func replaceShowOnTable(newNode, parent SQLNode) {
	parent.(*Show).OnTable = newNode.(TableName)
}

func replaceShowShowCollationFilterOpt(newNode, parent SQLNode) {
	parent.(*Show).ShowCollationFilterOpt = newNode.(Expr)
}

func replaceShowTable(newNode, parent SQLNode) {
	parent.(*Show).Table = newNode.(TableName)
}

func replaceShowFilterFilter(newNode, parent SQLNode) {
	parent.(*ShowFilter).Filter = newNode.(Expr)
}

func replaceStarExprTableName(newNode, parent SQLNode) {
	parent.(*StarExpr).TableName = newNode.(TableName)
}

func replaceStreamComments(newNode, parent SQLNode) {
	parent.(*Stream).Comments = newNode.(Comments)
}

func replaceStreamSelectExpr(newNode, parent SQLNode) {
	parent.(*Stream).SelectExpr = newNode.(SelectExpr)
}

func replaceStreamTable(newNode, parent SQLNode) {
	parent.(*Stream).Table = newNode.(TableName)
}

func replaceSubquerySelect(newNode, parent SQLNode) {
	parent.(*Subquery).Select = newNode.(SelectStatement)
}

func replaceSubstrExprFrom(newNode, parent SQLNode) {
	parent.(*SubstrExpr).From = newNode.(Expr)
}

func replaceSubstrExprName(newNode, parent SQLNode) {
	parent.(*SubstrExpr).Name = newNode.(*ColName)
}

func replaceSubstrExprStrVal(newNode, parent SQLNode) {
	parent.(*SubstrExpr).StrVal = newNode.(*Literal)
}

func replaceSubstrExprTo(newNode, parent SQLNode) {
	parent.(*SubstrExpr).To = newNode.(Expr)
}

type replaceTableExprsItems int

func (r *replaceTableExprsItems) replace(newNode, container SQLNode) {
	container.(TableExprs)[int(*r)] = newNode.(TableExpr)
}

func (r *replaceTableExprsItems) inc() {
	*r++
}

func replaceTableNameName(newNode, parent SQLNode) {
	tmp := parent.(TableName)
	tmp.Name = newNode.(TableIdent)
}

func replaceTableNameQualifier(newNode, parent SQLNode) {
	tmp := parent.(TableName)
	tmp.Qualifier = newNode.(TableIdent)
}

type replaceTableNamesItems int

func (r *replaceTableNamesItems) replace(newNode, container SQLNode) {
	container.(TableNames)[int(*r)] = newNode.(TableName)
}

func (r *replaceTableNamesItems) inc() {
	*r++
}

type replaceTableSpecColumns int

func (r *replaceTableSpecColumns) replace(newNode, container SQLNode) {
	container.(*TableSpec).Columns[int(*r)] = newNode.(*ColumnDefinition)
}

func (r *replaceTableSpecColumns) inc() {
	*r++
}

type replaceTableSpecConstraints int

func (r *replaceTableSpecConstraints) replace(newNode, container SQLNode) {
	container.(*TableSpec).Constraints[int(*r)] = newNode.(*ConstraintDefinition)
}

func (r *replaceTableSpecConstraints) inc() {
	*r++
}

type replaceTableSpecIndexes int

func (r *replaceTableSpecIndexes) replace(newNode, container SQLNode) {
	container.(*TableSpec).Indexes[int(*r)] = newNode.(*IndexDefinition)
}

func (r *replaceTableSpecIndexes) inc() {
	*r++
}

func replaceTimestampFuncExprExpr1(newNode, parent SQLNode) {
	parent.(*TimestampFuncExpr).Expr1 = newNode.(Expr)
}

func replaceTimestampFuncExprExpr2(newNode, parent SQLNode) {
	parent.(*TimestampFuncExpr).Expr2 = newNode.(Expr)
}

func replaceUnaryExprExpr(newNode, parent SQLNode) {
	parent.(*UnaryExpr).Expr = newNode.(Expr)
}

func replaceUnionFirstStatement(newNode, parent SQLNode) {
	parent.(*Union).FirstStatement = newNode.(SelectStatement)
}

func replaceUnionLimit(newNode, parent SQLNode) {
	parent.(*Union).Limit = newNode.(*Limit)
}

func replaceUnionOrderBy(newNode, parent SQLNode) {
	parent.(*Union).OrderBy = newNode.(OrderBy)
}

type replaceUnionUnionSelects int

func (r *replaceUnionUnionSelects) replace(newNode, container SQLNode) {
	container.(*Union).UnionSelects[int(*r)] = newNode.(*UnionSelect)
}

func (r *replaceUnionUnionSelects) inc() {
	*r++
}

func replaceUnionSelectStatement(newNode, parent SQLNode) {
	parent.(*UnionSelect).Statement = newNode.(SelectStatement)
}

func replaceUpdateComments(newNode, parent SQLNode) {
	parent.(*Update).Comments = newNode.(Comments)
}

func replaceUpdateExprs(newNode, parent SQLNode) {
	parent.(*Update).Exprs = newNode.(UpdateExprs)
}

func replaceUpdateLimit(newNode, parent SQLNode) {
	parent.(*Update).Limit = newNode.(*Limit)
}

func replaceUpdateOrderBy(newNode, parent SQLNode) {
	parent.(*Update).OrderBy = newNode.(OrderBy)
}

func replaceUpdateTableExprs(newNode, parent SQLNode) {
	parent.(*Update).TableExprs = newNode.(TableExprs)
}

func replaceUpdateWhere(newNode, parent SQLNode) {
	parent.(*Update).Where = newNode.(*Where)
}

func replaceUpdateExprExpr(newNode, parent SQLNode) {
	parent.(*UpdateExpr).Expr = newNode.(Expr)
}

func replaceUpdateExprName(newNode, parent SQLNode) {
	parent.(*UpdateExpr).Name = newNode.(*ColName)
}

type replaceUpdateExprsItems int

func (r *replaceUpdateExprsItems) replace(newNode, container SQLNode) {
	container.(UpdateExprs)[int(*r)] = newNode.(*UpdateExpr)
}

func (r *replaceUpdateExprsItems) inc() {
	*r++
}

func replaceUseDBName(newNode, parent SQLNode) {
	parent.(*Use).DBName = newNode.(TableIdent)
}

func replaceVStreamComments(newNode, parent SQLNode) {
	parent.(*VStream).Comments = newNode.(Comments)
}

func replaceVStreamLimit(newNode, parent SQLNode) {
	parent.(*VStream).Limit = newNode.(*Limit)
}

func replaceVStreamSelectExpr(newNode, parent SQLNode) {
	parent.(*VStream).SelectExpr = newNode.(SelectExpr)
}

func replaceVStreamTable(newNode, parent SQLNode) {
	parent.(*VStream).Table = newNode.(TableName)
}

func replaceVStreamWhere(newNode, parent SQLNode) {
	parent.(*VStream).Where = newNode.(*Where)
}

type replaceValTupleItems int

func (r *replaceValTupleItems) replace(newNode, container SQLNode) {
	container.(ValTuple)[int(*r)] = newNode.(Expr)
}

func (r *replaceValTupleItems) inc() {
	*r++
}

type replaceValuesItems int

func (r *replaceValuesItems) replace(newNode, container SQLNode) {
	container.(Values)[int(*r)] = newNode.(ValTuple)
}

func (r *replaceValuesItems) inc() {
	*r++
}

func replaceValuesFuncExprName(newNode, parent SQLNode) {
	parent.(*ValuesFuncExpr).Name = newNode.(*ColName)
}

func replaceVindexParamKey(newNode, parent SQLNode) {
	tmp := parent.(VindexParam)
	tmp.Key = newNode.(ColIdent)
}

func replaceVindexSpecName(newNode, parent SQLNode) {
	parent.(*VindexSpec).Name = newNode.(ColIdent)
}

type replaceVindexSpecParams int

func (r *replaceVindexSpecParams) replace(newNode, container SQLNode) {
	container.(*VindexSpec).Params[int(*r)] = newNode.(VindexParam)
}

func (r *replaceVindexSpecParams) inc() {
	*r++
}

func replaceVindexSpecType(newNode, parent SQLNode) {
	parent.(*VindexSpec).Type = newNode.(ColIdent)
}

func replaceWhenCond(newNode, parent SQLNode) {
	parent.(*When).Cond = newNode.(Expr)
}

func replaceWhenVal(newNode, parent SQLNode) {
	parent.(*When).Val = newNode.(Expr)
}

func replaceWhereExpr(newNode, parent SQLNode) {
	parent.(*Where).Expr = newNode.(Expr)
}

func replaceXorExprLeft(newNode, parent SQLNode) {
	parent.(*XorExpr).Left = newNode.(Expr)
}

func replaceXorExprRight(newNode, parent SQLNode) {
	parent.(*XorExpr).Right = newNode.(Expr)
}

// apply is where the visiting happens. Here is where we keep the big switch-case that will be used
// to do the actual visiting of SQLNodes
func (a *application) apply(parent, node SQLNode, replacer replacerFunc) {
	if node == nil || isNilValue(node) {
		return
	}

	// avoid heap-allocating a new cursor for each apply call; reuse a.cursor instead
	saved := a.cursor
	a.cursor.replacer = replacer
	a.cursor.node = node
	a.cursor.parent = parent

	if a.pre != nil && !a.pre(&a.cursor) {
		a.cursor = saved
		return
	}

	// walk children
	// (the order of the cases is alphabetical)
	switch n := node.(type) {
	case nil:
	case AccessMode:

	case *AliasedExpr:
		a.apply(node, n.As, replaceAliasedExprAs)
		a.apply(node, n.Expr, replaceAliasedExprExpr)

	case *AliasedTableExpr:
		a.apply(node, n.As, replaceAliasedTableExprAs)
		a.apply(node, n.Expr, replaceAliasedTableExprExpr)
		a.apply(node, n.Hints, replaceAliasedTableExprHints)
		a.apply(node, n.Partitions, replaceAliasedTableExprPartitions)

	case *AndExpr:
		a.apply(node, n.Left, replaceAndExprLeft)
		a.apply(node, n.Right, replaceAndExprRight)

	case Argument:

	case *AutoIncSpec:
		a.apply(node, n.Column, replaceAutoIncSpecColumn)
		a.apply(node, n.Sequence, replaceAutoIncSpecSequence)

	case *Begin:

	case *BinaryExpr:
		a.apply(node, n.Left, replaceBinaryExprLeft)
		a.apply(node, n.Right, replaceBinaryExprRight)

	case BoolVal:

	case *CaseExpr:
		a.apply(node, n.Else, replaceCaseExprElse)
		a.apply(node, n.Expr, replaceCaseExprExpr)
		replacerWhens := replaceCaseExprWhens(0)
		replacerWhensB := &replacerWhens
		for _, item := range n.Whens {
			a.apply(node, item, replacerWhensB.replace)
			replacerWhensB.inc()
		}

	case ColIdent:

	case *ColName:
		a.apply(node, n.Name, replaceColNameName)
		a.apply(node, n.Qualifier, replaceColNameQualifier)

	case *CollateExpr:
		a.apply(node, n.Expr, replaceCollateExprExpr)

	case *ColumnDefinition:
		a.apply(node, n.Name, replaceColumnDefinitionName)

	case *ColumnType:
		a.apply(node, n.Autoincrement, replaceColumnTypeAutoincrement)
		a.apply(node, n.Comment, replaceColumnTypeComment)
		a.apply(node, n.Default, replaceColumnTypeDefault)
		a.apply(node, n.Length, replaceColumnTypeLength)
		a.apply(node, n.NotNull, replaceColumnTypeNotNull)
		a.apply(node, n.OnUpdate, replaceColumnTypeOnUpdate)
		a.apply(node, n.Scale, replaceColumnTypeScale)
		a.apply(node, n.Unsigned, replaceColumnTypeUnsigned)
		a.apply(node, n.Zerofill, replaceColumnTypeZerofill)

	case Columns:
		replacer := replaceColumnsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case Comments:

	case *Commit:

	case *ComparisonExpr:
		a.apply(node, n.Escape, replaceComparisonExprEscape)
		a.apply(node, n.Left, replaceComparisonExprLeft)
		a.apply(node, n.Right, replaceComparisonExprRight)

	case *ConstraintDefinition:
		a.apply(node, n.Details, replaceConstraintDefinitionDetails)

	case *ConvertExpr:
		a.apply(node, n.Expr, replaceConvertExprExpr)
		a.apply(node, n.Type, replaceConvertExprType)

	case *ConvertType:
		a.apply(node, n.Length, replaceConvertTypeLength)
		a.apply(node, n.Scale, replaceConvertTypeScale)

	case *ConvertUsingExpr:
		a.apply(node, n.Expr, replaceConvertUsingExprExpr)

	case *CurTimeFuncExpr:
		a.apply(node, n.Fsp, replaceCurTimeFuncExprFsp)
		a.apply(node, n.Name, replaceCurTimeFuncExprName)

	case *DBDDL:

	case *DDL:
		a.apply(node, n.AutoIncSpec, replaceDDLAutoIncSpec)
		a.apply(node, n.FromTables, replaceDDLFromTables)
		a.apply(node, n.OptLike, replaceDDLOptLike)
		a.apply(node, n.PartitionSpec, replaceDDLPartitionSpec)
		a.apply(node, n.Table, replaceDDLTable)
		a.apply(node, n.TableSpec, replaceDDLTableSpec)
		a.apply(node, n.ToTables, replaceDDLToTables)
		replacerVindexCols := replaceDDLVindexCols(0)
		replacerVindexColsB := &replacerVindexCols
		for _, item := range n.VindexCols {
			a.apply(node, item, replacerVindexColsB.replace)
			replacerVindexColsB.inc()
		}
		a.apply(node, n.VindexSpec, replaceDDLVindexSpec)

	case *Default:

	case *Delete:
		a.apply(node, n.Comments, replaceDeleteComments)
		a.apply(node, n.Limit, replaceDeleteLimit)
		a.apply(node, n.OrderBy, replaceDeleteOrderBy)
		a.apply(node, n.Partitions, replaceDeletePartitions)
		a.apply(node, n.TableExprs, replaceDeleteTableExprs)
		a.apply(node, n.Targets, replaceDeleteTargets)
		a.apply(node, n.Where, replaceDeleteWhere)

	case *ExistsExpr:
		a.apply(node, n.Subquery, replaceExistsExprSubquery)

	case *Explain:
		a.apply(node, n.Statement, replaceExplainStatement)

	case Exprs:
		replacer := replaceExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *ForeignKeyDefinition:
		a.apply(node, n.OnDelete, replaceForeignKeyDefinitionOnDelete)
		a.apply(node, n.OnUpdate, replaceForeignKeyDefinitionOnUpdate)
		a.apply(node, n.ReferencedColumns, replaceForeignKeyDefinitionReferencedColumns)
		a.apply(node, n.ReferencedTable, replaceForeignKeyDefinitionReferencedTable)
		a.apply(node, n.Source, replaceForeignKeyDefinitionSource)

	case *FuncExpr:
		a.apply(node, n.Exprs, replaceFuncExprExprs)
		a.apply(node, n.Name, replaceFuncExprName)
		a.apply(node, n.Qualifier, replaceFuncExprQualifier)

	case GroupBy:
		replacer := replaceGroupByItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *GroupConcatExpr:
		a.apply(node, n.Exprs, replaceGroupConcatExprExprs)
		a.apply(node, n.Limit, replaceGroupConcatExprLimit)
		a.apply(node, n.OrderBy, replaceGroupConcatExprOrderBy)

	case *IndexDefinition:
		a.apply(node, n.Info, replaceIndexDefinitionInfo)

	case *IndexHints:
		replacerIndexes := replaceIndexHintsIndexes(0)
		replacerIndexesB := &replacerIndexes
		for _, item := range n.Indexes {
			a.apply(node, item, replacerIndexesB.replace)
			replacerIndexesB.inc()
		}

	case *IndexInfo:
		a.apply(node, n.Name, replaceIndexInfoName)

	case *Insert:
		a.apply(node, n.Columns, replaceInsertColumns)
		a.apply(node, n.Comments, replaceInsertComments)
		a.apply(node, n.OnDup, replaceInsertOnDup)
		a.apply(node, n.Partitions, replaceInsertPartitions)
		a.apply(node, n.Rows, replaceInsertRows)
		a.apply(node, n.Table, replaceInsertTable)

	case *IntervalExpr:
		a.apply(node, n.Expr, replaceIntervalExprExpr)

	case *IsExpr:
		a.apply(node, n.Expr, replaceIsExprExpr)

	case IsolationLevel:

	case JoinCondition:
		a.apply(node, n.On, replaceJoinConditionOn)
		a.apply(node, n.Using, replaceJoinConditionUsing)

	case *JoinTableExpr:
		a.apply(node, n.Condition, replaceJoinTableExprCondition)
		a.apply(node, n.LeftExpr, replaceJoinTableExprLeftExpr)
		a.apply(node, n.RightExpr, replaceJoinTableExprRightExpr)

	case *Limit:
		a.apply(node, n.Offset, replaceLimitOffset)
		a.apply(node, n.Rowcount, replaceLimitRowcount)

	case ListArg:

	case *Literal:

	case *MatchExpr:
		a.apply(node, n.Columns, replaceMatchExprColumns)
		a.apply(node, n.Expr, replaceMatchExprExpr)

	case Nextval:
		a.apply(node, n.Expr, replaceNextvalExpr)

	case *NotExpr:
		a.apply(node, n.Expr, replaceNotExprExpr)

	case *NullVal:

	case OnDup:
		replacer := replaceOnDupItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *OptLike:
		a.apply(node, n.LikeTable, replaceOptLikeLikeTable)

	case *OrExpr:
		a.apply(node, n.Left, replaceOrExprLeft)
		a.apply(node, n.Right, replaceOrExprRight)

	case *Order:
		a.apply(node, n.Expr, replaceOrderExpr)

	case OrderBy:
		replacer := replaceOrderByItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *OtherAdmin:

	case *OtherRead:

	case *ParenSelect:
		a.apply(node, n.Select, replaceParenSelectSelect)

	case *ParenTableExpr:
		a.apply(node, n.Exprs, replaceParenTableExprExprs)

	case *PartitionDefinition:
		a.apply(node, n.Limit, replacePartitionDefinitionLimit)
		a.apply(node, n.Name, replacePartitionDefinitionName)

	case *PartitionSpec:
		replacerDefinitions := replacePartitionSpecDefinitions(0)
		replacerDefinitionsB := &replacerDefinitions
		for _, item := range n.Definitions {
			a.apply(node, item, replacerDefinitionsB.replace)
			replacerDefinitionsB.inc()
		}
		a.apply(node, n.Name, replacePartitionSpecName)

	case Partitions:
		replacer := replacePartitionsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *RangeCond:
		a.apply(node, n.From, replaceRangeCondFrom)
		a.apply(node, n.Left, replaceRangeCondLeft)
		a.apply(node, n.To, replaceRangeCondTo)

	case ReferenceAction:

	case *Release:
		a.apply(node, n.Name, replaceReleaseName)

	case *Rollback:

	case *SRollback:
		a.apply(node, n.Name, replaceSRollbackName)

	case *Savepoint:
		a.apply(node, n.Name, replaceSavepointName)

	case *Select:
		a.apply(node, n.Comments, replaceSelectComments)
		a.apply(node, n.From, replaceSelectFrom)
		a.apply(node, n.GroupBy, replaceSelectGroupBy)
		a.apply(node, n.Having, replaceSelectHaving)
		a.apply(node, n.Limit, replaceSelectLimit)
		a.apply(node, n.OrderBy, replaceSelectOrderBy)
		a.apply(node, n.SelectExprs, replaceSelectSelectExprs)
		a.apply(node, n.Where, replaceSelectWhere)

	case SelectExprs:
		replacer := replaceSelectExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *Set:
		a.apply(node, n.Comments, replaceSetComments)
		a.apply(node, n.Exprs, replaceSetExprs)

	case *SetExpr:
		a.apply(node, n.Expr, replaceSetExprExpr)
		a.apply(node, n.Name, replaceSetExprName)

	case SetExprs:
		replacer := replaceSetExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *SetTransaction:
		replacerCharacteristics := replaceSetTransactionCharacteristics(0)
		replacerCharacteristicsB := &replacerCharacteristics
		for _, item := range n.Characteristics {
			a.apply(node, item, replacerCharacteristicsB.replace)
			replacerCharacteristicsB.inc()
		}
		a.apply(node, n.Comments, replaceSetTransactionComments)

	case *Show:
		a.apply(node, n.OnTable, replaceShowOnTable)
		a.apply(node, n.ShowCollationFilterOpt, replaceShowShowCollationFilterOpt)
		a.apply(node, n.Table, replaceShowTable)

	case *ShowFilter:
		a.apply(node, n.Filter, replaceShowFilterFilter)

	case *StarExpr:
		a.apply(node, n.TableName, replaceStarExprTableName)

	case *Stream:
		a.apply(node, n.Comments, replaceStreamComments)
		a.apply(node, n.SelectExpr, replaceStreamSelectExpr)
		a.apply(node, n.Table, replaceStreamTable)

	case *Subquery:
		a.apply(node, n.Select, replaceSubquerySelect)

	case *SubstrExpr:
		a.apply(node, n.From, replaceSubstrExprFrom)
		a.apply(node, n.Name, replaceSubstrExprName)
		a.apply(node, n.StrVal, replaceSubstrExprStrVal)
		a.apply(node, n.To, replaceSubstrExprTo)

	case TableExprs:
		replacer := replaceTableExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case TableIdent:

	case TableName:
		a.apply(node, n.Name, replaceTableNameName)
		a.apply(node, n.Qualifier, replaceTableNameQualifier)

	case TableNames:
		replacer := replaceTableNamesItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *TableSpec:
		replacerColumns := replaceTableSpecColumns(0)
		replacerColumnsB := &replacerColumns
		for _, item := range n.Columns {
			a.apply(node, item, replacerColumnsB.replace)
			replacerColumnsB.inc()
		}
		replacerConstraints := replaceTableSpecConstraints(0)
		replacerConstraintsB := &replacerConstraints
		for _, item := range n.Constraints {
			a.apply(node, item, replacerConstraintsB.replace)
			replacerConstraintsB.inc()
		}
		replacerIndexes := replaceTableSpecIndexes(0)
		replacerIndexesB := &replacerIndexes
		for _, item := range n.Indexes {
			a.apply(node, item, replacerIndexesB.replace)
			replacerIndexesB.inc()
		}

	case *TimestampFuncExpr:
		a.apply(node, n.Expr1, replaceTimestampFuncExprExpr1)
		a.apply(node, n.Expr2, replaceTimestampFuncExprExpr2)

	case *UnaryExpr:
		a.apply(node, n.Expr, replaceUnaryExprExpr)

	case *Union:
		a.apply(node, n.FirstStatement, replaceUnionFirstStatement)
		a.apply(node, n.Limit, replaceUnionLimit)
		a.apply(node, n.OrderBy, replaceUnionOrderBy)
		replacerUnionSelects := replaceUnionUnionSelects(0)
		replacerUnionSelectsB := &replacerUnionSelects
		for _, item := range n.UnionSelects {
			a.apply(node, item, replacerUnionSelectsB.replace)
			replacerUnionSelectsB.inc()
		}

	case *UnionSelect:
		a.apply(node, n.Statement, replaceUnionSelectStatement)

	case *Update:
		a.apply(node, n.Comments, replaceUpdateComments)
		a.apply(node, n.Exprs, replaceUpdateExprs)
		a.apply(node, n.Limit, replaceUpdateLimit)
		a.apply(node, n.OrderBy, replaceUpdateOrderBy)
		a.apply(node, n.TableExprs, replaceUpdateTableExprs)
		a.apply(node, n.Where, replaceUpdateWhere)

	case *UpdateExpr:
		a.apply(node, n.Expr, replaceUpdateExprExpr)
		a.apply(node, n.Name, replaceUpdateExprName)

	case UpdateExprs:
		replacer := replaceUpdateExprsItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *Use:
		a.apply(node, n.DBName, replaceUseDBName)

	case *VStream:
		a.apply(node, n.Comments, replaceVStreamComments)
		a.apply(node, n.Limit, replaceVStreamLimit)
		a.apply(node, n.SelectExpr, replaceVStreamSelectExpr)
		a.apply(node, n.Table, replaceVStreamTable)
		a.apply(node, n.Where, replaceVStreamWhere)

	case ValTuple:
		replacer := replaceValTupleItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case Values:
		replacer := replaceValuesItems(0)
		replacerRef := &replacer
		for _, item := range n {
			a.apply(node, item, replacerRef.replace)
			replacerRef.inc()
		}

	case *ValuesFuncExpr:
		a.apply(node, n.Name, replaceValuesFuncExprName)

	case VindexParam:
		a.apply(node, n.Key, replaceVindexParamKey)

	case *VindexSpec:
		a.apply(node, n.Name, replaceVindexSpecName)
		replacerParams := replaceVindexSpecParams(0)
		replacerParamsB := &replacerParams
		for _, item := range n.Params {
			a.apply(node, item, replacerParamsB.replace)
			replacerParamsB.inc()
		}
		a.apply(node, n.Type, replaceVindexSpecType)

	case *When:
		a.apply(node, n.Cond, replaceWhenCond)
		a.apply(node, n.Val, replaceWhenVal)

	case *Where:
		a.apply(node, n.Expr, replaceWhereExpr)

	case *XorExpr:
		a.apply(node, n.Left, replaceXorExprLeft)
		a.apply(node, n.Right, replaceXorExprRight)

	default:
		panic("unknown ast type " + reflect.TypeOf(node).String())
	}

	if a.post != nil && !a.post(&a.cursor) {
		panic(abort)
	}

	a.cursor = saved
}

func isNilValue(i interface{}) bool {
	valueOf := reflect.ValueOf(i)
	kind := valueOf.Kind()
	isNullable := kind == reflect.Ptr || kind == reflect.Array || kind == reflect.Slice
	return isNullable && valueOf.IsNil()
}
