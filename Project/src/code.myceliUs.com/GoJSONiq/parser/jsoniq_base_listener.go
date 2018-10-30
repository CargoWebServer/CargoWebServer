// Code generated from jsoniq.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // jsoniq

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasejsoniqListener is a complete listener for a parse tree produced by jsoniqParser.
type BasejsoniqListener struct{}

var _ jsoniqListener = &BasejsoniqListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasejsoniqListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasejsoniqListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasejsoniqListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasejsoniqListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasejsoniqListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasejsoniqListener) ExitProgram(ctx *ProgramContext) {}

// EnterMainModule is called when production mainModule is entered.
func (s *BasejsoniqListener) EnterMainModule(ctx *MainModuleContext) {}

// ExitMainModule is called when production mainModule is exited.
func (s *BasejsoniqListener) ExitMainModule(ctx *MainModuleContext) {}

// EnterLibraryModule is called when production libraryModule is entered.
func (s *BasejsoniqListener) EnterLibraryModule(ctx *LibraryModuleContext) {}

// ExitLibraryModule is called when production libraryModule is exited.
func (s *BasejsoniqListener) ExitLibraryModule(ctx *LibraryModuleContext) {}

// EnterProlog is called when production prolog is entered.
func (s *BasejsoniqListener) EnterProlog(ctx *PrologContext) {}

// ExitProlog is called when production prolog is exited.
func (s *BasejsoniqListener) ExitProlog(ctx *PrologContext) {}

// EnterDefaultCollationDecl is called when production defaultCollationDecl is entered.
func (s *BasejsoniqListener) EnterDefaultCollationDecl(ctx *DefaultCollationDeclContext) {}

// ExitDefaultCollationDecl is called when production defaultCollationDecl is exited.
func (s *BasejsoniqListener) ExitDefaultCollationDecl(ctx *DefaultCollationDeclContext) {}

// EnterOrderingModeDecl is called when production orderingModeDecl is entered.
func (s *BasejsoniqListener) EnterOrderingModeDecl(ctx *OrderingModeDeclContext) {}

// ExitOrderingModeDecl is called when production orderingModeDecl is exited.
func (s *BasejsoniqListener) ExitOrderingModeDecl(ctx *OrderingModeDeclContext) {}

// EnterEmptyOrderDecl is called when production emptyOrderDecl is entered.
func (s *BasejsoniqListener) EnterEmptyOrderDecl(ctx *EmptyOrderDeclContext) {}

// ExitEmptyOrderDecl is called when production emptyOrderDecl is exited.
func (s *BasejsoniqListener) ExitEmptyOrderDecl(ctx *EmptyOrderDeclContext) {}

// EnterDecimalFormatDecl is called when production decimalFormatDecl is entered.
func (s *BasejsoniqListener) EnterDecimalFormatDecl(ctx *DecimalFormatDeclContext) {}

// ExitDecimalFormatDecl is called when production decimalFormatDecl is exited.
func (s *BasejsoniqListener) ExitDecimalFormatDecl(ctx *DecimalFormatDeclContext) {}

// EnterDfPropertyName is called when production dfPropertyName is entered.
func (s *BasejsoniqListener) EnterDfPropertyName(ctx *DfPropertyNameContext) {}

// ExitDfPropertyName is called when production dfPropertyName is exited.
func (s *BasejsoniqListener) ExitDfPropertyName(ctx *DfPropertyNameContext) {}

// EnterModuleImport is called when production moduleImport is entered.
func (s *BasejsoniqListener) EnterModuleImport(ctx *ModuleImportContext) {}

// ExitModuleImport is called when production moduleImport is exited.
func (s *BasejsoniqListener) ExitModuleImport(ctx *ModuleImportContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BasejsoniqListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BasejsoniqListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *BasejsoniqListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production functionDecl is exited.
func (s *BasejsoniqListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BasejsoniqListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BasejsoniqListener) ExitParamList(ctx *ParamListContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasejsoniqListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasejsoniqListener) ExitExpr(ctx *ExprContext) {}

// EnterExprSingle is called when production exprSingle is entered.
func (s *BasejsoniqListener) EnterExprSingle(ctx *ExprSingleContext) {}

// ExitExprSingle is called when production exprSingle is exited.
func (s *BasejsoniqListener) ExitExprSingle(ctx *ExprSingleContext) {}

// EnterFlowrExpr is called when production flowrExpr is entered.
func (s *BasejsoniqListener) EnterFlowrExpr(ctx *FlowrExprContext) {}

// ExitFlowrExpr is called when production flowrExpr is exited.
func (s *BasejsoniqListener) ExitFlowrExpr(ctx *FlowrExprContext) {}

// EnterForClause is called when production forClause is entered.
func (s *BasejsoniqListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *BasejsoniqListener) ExitForClause(ctx *ForClauseContext) {}

// EnterLetClause is called when production letClause is entered.
func (s *BasejsoniqListener) EnterLetClause(ctx *LetClauseContext) {}

// ExitLetClause is called when production letClause is exited.
func (s *BasejsoniqListener) ExitLetClause(ctx *LetClauseContext) {}

// EnterCountClause is called when production countClause is entered.
func (s *BasejsoniqListener) EnterCountClause(ctx *CountClauseContext) {}

// ExitCountClause is called when production countClause is exited.
func (s *BasejsoniqListener) ExitCountClause(ctx *CountClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BasejsoniqListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BasejsoniqListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BasejsoniqListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BasejsoniqListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BasejsoniqListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BasejsoniqListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterQuantifiedExpr is called when production quantifiedExpr is entered.
func (s *BasejsoniqListener) EnterQuantifiedExpr(ctx *QuantifiedExprContext) {}

// ExitQuantifiedExpr is called when production quantifiedExpr is exited.
func (s *BasejsoniqListener) ExitQuantifiedExpr(ctx *QuantifiedExprContext) {}

// EnterSwitchExpr is called when production switchExpr is entered.
func (s *BasejsoniqListener) EnterSwitchExpr(ctx *SwitchExprContext) {}

// ExitSwitchExpr is called when production switchExpr is exited.
func (s *BasejsoniqListener) ExitSwitchExpr(ctx *SwitchExprContext) {}

// EnterSwitchCaseClause is called when production switchCaseClause is entered.
func (s *BasejsoniqListener) EnterSwitchCaseClause(ctx *SwitchCaseClauseContext) {}

// ExitSwitchCaseClause is called when production switchCaseClause is exited.
func (s *BasejsoniqListener) ExitSwitchCaseClause(ctx *SwitchCaseClauseContext) {}

// EnterTypeswitchExpr is called when production typeswitchExpr is entered.
func (s *BasejsoniqListener) EnterTypeswitchExpr(ctx *TypeswitchExprContext) {}

// ExitTypeswitchExpr is called when production typeswitchExpr is exited.
func (s *BasejsoniqListener) ExitTypeswitchExpr(ctx *TypeswitchExprContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BasejsoniqListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BasejsoniqListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BasejsoniqListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BasejsoniqListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterTryCatchExpr is called when production tryCatchExpr is entered.
func (s *BasejsoniqListener) EnterTryCatchExpr(ctx *TryCatchExprContext) {}

// ExitTryCatchExpr is called when production tryCatchExpr is exited.
func (s *BasejsoniqListener) ExitTryCatchExpr(ctx *TryCatchExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BasejsoniqListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BasejsoniqListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BasejsoniqListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BasejsoniqListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BasejsoniqListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BasejsoniqListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterComparisonExpr2 is called when production comparisonExpr2 is entered.
func (s *BasejsoniqListener) EnterComparisonExpr2(ctx *ComparisonExpr2Context) {}

// ExitComparisonExpr2 is called when production comparisonExpr2 is exited.
func (s *BasejsoniqListener) ExitComparisonExpr2(ctx *ComparisonExpr2Context) {}

// EnterComparisonExpr is called when production comparisonExpr is entered.
func (s *BasejsoniqListener) EnterComparisonExpr(ctx *ComparisonExprContext) {}

// ExitComparisonExpr is called when production comparisonExpr is exited.
func (s *BasejsoniqListener) ExitComparisonExpr(ctx *ComparisonExprContext) {}

// EnterStringConcatExpr is called when production stringConcatExpr is entered.
func (s *BasejsoniqListener) EnterStringConcatExpr(ctx *StringConcatExprContext) {}

// ExitStringConcatExpr is called when production stringConcatExpr is exited.
func (s *BasejsoniqListener) ExitStringConcatExpr(ctx *StringConcatExprContext) {}

// EnterRangeExpr is called when production rangeExpr is entered.
func (s *BasejsoniqListener) EnterRangeExpr(ctx *RangeExprContext) {}

// ExitRangeExpr is called when production rangeExpr is exited.
func (s *BasejsoniqListener) ExitRangeExpr(ctx *RangeExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BasejsoniqListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BasejsoniqListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterMultiplicativeExpr is called when production multiplicativeExpr is entered.
func (s *BasejsoniqListener) EnterMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// ExitMultiplicativeExpr is called when production multiplicativeExpr is exited.
func (s *BasejsoniqListener) ExitMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// EnterInstanceofExpr is called when production instanceofExpr is entered.
func (s *BasejsoniqListener) EnterInstanceofExpr(ctx *InstanceofExprContext) {}

// ExitInstanceofExpr is called when production instanceofExpr is exited.
func (s *BasejsoniqListener) ExitInstanceofExpr(ctx *InstanceofExprContext) {}

// EnterTreatExpr is called when production treatExpr is entered.
func (s *BasejsoniqListener) EnterTreatExpr(ctx *TreatExprContext) {}

// ExitTreatExpr is called when production treatExpr is exited.
func (s *BasejsoniqListener) ExitTreatExpr(ctx *TreatExprContext) {}

// EnterCastableExpr is called when production castableExpr is entered.
func (s *BasejsoniqListener) EnterCastableExpr(ctx *CastableExprContext) {}

// ExitCastableExpr is called when production castableExpr is exited.
func (s *BasejsoniqListener) ExitCastableExpr(ctx *CastableExprContext) {}

// EnterCastExpr is called when production castExpr is entered.
func (s *BasejsoniqListener) EnterCastExpr(ctx *CastExprContext) {}

// ExitCastExpr is called when production castExpr is exited.
func (s *BasejsoniqListener) ExitCastExpr(ctx *CastExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BasejsoniqListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BasejsoniqListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterSimpleMapExpr is called when production simpleMapExpr is entered.
func (s *BasejsoniqListener) EnterSimpleMapExpr(ctx *SimpleMapExprContext) {}

// ExitSimpleMapExpr is called when production simpleMapExpr is exited.
func (s *BasejsoniqListener) ExitSimpleMapExpr(ctx *SimpleMapExprContext) {}

// EnterPostfixExpr is called when production postfixExpr is entered.
func (s *BasejsoniqListener) EnterPostfixExpr(ctx *PostfixExprContext) {}

// ExitPostfixExpr is called when production postfixExpr is exited.
func (s *BasejsoniqListener) ExitPostfixExpr(ctx *PostfixExprContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BasejsoniqListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BasejsoniqListener) ExitPredicate(ctx *PredicateContext) {}

// EnterObjectLookup is called when production objectLookup is entered.
func (s *BasejsoniqListener) EnterObjectLookup(ctx *ObjectLookupContext) {}

// ExitObjectLookup is called when production objectLookup is exited.
func (s *BasejsoniqListener) ExitObjectLookup(ctx *ObjectLookupContext) {}

// EnterArrayLookup is called when production arrayLookup is entered.
func (s *BasejsoniqListener) EnterArrayLookup(ctx *ArrayLookupContext) {}

// ExitArrayLookup is called when production arrayLookup is exited.
func (s *BasejsoniqListener) ExitArrayLookup(ctx *ArrayLookupContext) {}

// EnterArrayUnboxing is called when production arrayUnboxing is entered.
func (s *BasejsoniqListener) EnterArrayUnboxing(ctx *ArrayUnboxingContext) {}

// ExitArrayUnboxing is called when production arrayUnboxing is exited.
func (s *BasejsoniqListener) ExitArrayUnboxing(ctx *ArrayUnboxingContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BasejsoniqListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BasejsoniqListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterVarRef is called when production varRef is entered.
func (s *BasejsoniqListener) EnterVarRef(ctx *VarRefContext) {}

// ExitVarRef is called when production varRef is exited.
func (s *BasejsoniqListener) ExitVarRef(ctx *VarRefContext) {}

// EnterParenthesizedExpr is called when production parenthesizedExpr is entered.
func (s *BasejsoniqListener) EnterParenthesizedExpr(ctx *ParenthesizedExprContext) {}

// ExitParenthesizedExpr is called when production parenthesizedExpr is exited.
func (s *BasejsoniqListener) ExitParenthesizedExpr(ctx *ParenthesizedExprContext) {}

// EnterContextItemExpr is called when production contextItemExpr is entered.
func (s *BasejsoniqListener) EnterContextItemExpr(ctx *ContextItemExprContext) {}

// ExitContextItemExpr is called when production contextItemExpr is exited.
func (s *BasejsoniqListener) ExitContextItemExpr(ctx *ContextItemExprContext) {}

// EnterOrderedExpr is called when production orderedExpr is entered.
func (s *BasejsoniqListener) EnterOrderedExpr(ctx *OrderedExprContext) {}

// ExitOrderedExpr is called when production orderedExpr is exited.
func (s *BasejsoniqListener) ExitOrderedExpr(ctx *OrderedExprContext) {}

// EnterUnorderedExpr is called when production unorderedExpr is entered.
func (s *BasejsoniqListener) EnterUnorderedExpr(ctx *UnorderedExprContext) {}

// ExitUnorderedExpr is called when production unorderedExpr is exited.
func (s *BasejsoniqListener) ExitUnorderedExpr(ctx *UnorderedExprContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BasejsoniqListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BasejsoniqListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BasejsoniqListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BasejsoniqListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasejsoniqListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasejsoniqListener) ExitArgument(ctx *ArgumentContext) {}

// EnterObjectConstructor is called when production objectConstructor is entered.
func (s *BasejsoniqListener) EnterObjectConstructor(ctx *ObjectConstructorContext) {}

// ExitObjectConstructor is called when production objectConstructor is exited.
func (s *BasejsoniqListener) ExitObjectConstructor(ctx *ObjectConstructorContext) {}

// EnterPairConstructor is called when production pairConstructor is entered.
func (s *BasejsoniqListener) EnterPairConstructor(ctx *PairConstructorContext) {}

// ExitPairConstructor is called when production pairConstructor is exited.
func (s *BasejsoniqListener) ExitPairConstructor(ctx *PairConstructorContext) {}

// EnterArrayConstructor is called when production arrayConstructor is entered.
func (s *BasejsoniqListener) EnterArrayConstructor(ctx *ArrayConstructorContext) {}

// ExitArrayConstructor is called when production arrayConstructor is exited.
func (s *BasejsoniqListener) ExitArrayConstructor(ctx *ArrayConstructorContext) {}

// EnterSequenceType is called when production sequenceType is entered.
func (s *BasejsoniqListener) EnterSequenceType(ctx *SequenceTypeContext) {}

// ExitSequenceType is called when production sequenceType is exited.
func (s *BasejsoniqListener) ExitSequenceType(ctx *SequenceTypeContext) {}

// EnterItemType is called when production itemType is entered.
func (s *BasejsoniqListener) EnterItemType(ctx *ItemTypeContext) {}

// ExitItemType is called when production itemType is exited.
func (s *BasejsoniqListener) ExitItemType(ctx *ItemTypeContext) {}

// EnterJsonItemTest is called when production jsonItemTest is entered.
func (s *BasejsoniqListener) EnterJsonItemTest(ctx *JsonItemTestContext) {}

// ExitJsonItemTest is called when production jsonItemTest is exited.
func (s *BasejsoniqListener) ExitJsonItemTest(ctx *JsonItemTestContext) {}

// EnterAtomicType is called when production atomicType is entered.
func (s *BasejsoniqListener) EnterAtomicType(ctx *AtomicTypeContext) {}

// ExitAtomicType is called when production atomicType is exited.
func (s *BasejsoniqListener) ExitAtomicType(ctx *AtomicTypeContext) {}

// EnterUriLiteral is called when production uriLiteral is entered.
func (s *BasejsoniqListener) EnterUriLiteral(ctx *UriLiteralContext) {}

// ExitUriLiteral is called when production uriLiteral is exited.
func (s *BasejsoniqListener) ExitUriLiteral(ctx *UriLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BasejsoniqListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BasejsoniqListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BasejsoniqListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BasejsoniqListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BasejsoniqListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BasejsoniqListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterNullLiteral is called when production nullLiteral is entered.
func (s *BasejsoniqListener) EnterNullLiteral(ctx *NullLiteralContext) {}

// ExitNullLiteral is called when production nullLiteral is exited.
func (s *BasejsoniqListener) ExitNullLiteral(ctx *NullLiteralContext) {}
