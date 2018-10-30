// Code generated from jsoniq.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // jsoniq

import "github.com/antlr/antlr4/runtime/Go/antlr"

// jsoniqListener is a complete listener for a parse tree produced by jsoniqParser.
type jsoniqListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterMainModule is called when entering the mainModule production.
	EnterMainModule(c *MainModuleContext)

	// EnterLibraryModule is called when entering the libraryModule production.
	EnterLibraryModule(c *LibraryModuleContext)

	// EnterProlog is called when entering the prolog production.
	EnterProlog(c *PrologContext)

	// EnterDefaultCollationDecl is called when entering the defaultCollationDecl production.
	EnterDefaultCollationDecl(c *DefaultCollationDeclContext)

	// EnterOrderingModeDecl is called when entering the orderingModeDecl production.
	EnterOrderingModeDecl(c *OrderingModeDeclContext)

	// EnterEmptyOrderDecl is called when entering the emptyOrderDecl production.
	EnterEmptyOrderDecl(c *EmptyOrderDeclContext)

	// EnterDecimalFormatDecl is called when entering the decimalFormatDecl production.
	EnterDecimalFormatDecl(c *DecimalFormatDeclContext)

	// EnterDfPropertyName is called when entering the dfPropertyName production.
	EnterDfPropertyName(c *DfPropertyNameContext)

	// EnterModuleImport is called when entering the moduleImport production.
	EnterModuleImport(c *ModuleImportContext)

	// EnterVarDecl is called when entering the varDecl production.
	EnterVarDecl(c *VarDeclContext)

	// EnterFunctionDecl is called when entering the functionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterExprSingle is called when entering the exprSingle production.
	EnterExprSingle(c *ExprSingleContext)

	// EnterFlowrExpr is called when entering the flowrExpr production.
	EnterFlowrExpr(c *FlowrExprContext)

	// EnterForClause is called when entering the forClause production.
	EnterForClause(c *ForClauseContext)

	// EnterLetClause is called when entering the letClause production.
	EnterLetClause(c *LetClauseContext)

	// EnterCountClause is called when entering the countClause production.
	EnterCountClause(c *CountClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterQuantifiedExpr is called when entering the quantifiedExpr production.
	EnterQuantifiedExpr(c *QuantifiedExprContext)

	// EnterSwitchExpr is called when entering the switchExpr production.
	EnterSwitchExpr(c *SwitchExprContext)

	// EnterSwitchCaseClause is called when entering the switchCaseClause production.
	EnterSwitchCaseClause(c *SwitchCaseClauseContext)

	// EnterTypeswitchExpr is called when entering the typeswitchExpr production.
	EnterTypeswitchExpr(c *TypeswitchExprContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterTryCatchExpr is called when entering the tryCatchExpr production.
	EnterTryCatchExpr(c *TryCatchExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterComparisonExpr2 is called when entering the comparisonExpr2 production.
	EnterComparisonExpr2(c *ComparisonExpr2Context)

	// EnterComparisonExpr is called when entering the comparisonExpr production.
	EnterComparisonExpr(c *ComparisonExprContext)

	// EnterStringConcatExpr is called when entering the stringConcatExpr production.
	EnterStringConcatExpr(c *StringConcatExprContext)

	// EnterRangeExpr is called when entering the rangeExpr production.
	EnterRangeExpr(c *RangeExprContext)

	// EnterAdditiveExpr is called when entering the additiveExpr production.
	EnterAdditiveExpr(c *AdditiveExprContext)

	// EnterMultiplicativeExpr is called when entering the multiplicativeExpr production.
	EnterMultiplicativeExpr(c *MultiplicativeExprContext)

	// EnterInstanceofExpr is called when entering the instanceofExpr production.
	EnterInstanceofExpr(c *InstanceofExprContext)

	// EnterTreatExpr is called when entering the treatExpr production.
	EnterTreatExpr(c *TreatExprContext)

	// EnterCastableExpr is called when entering the castableExpr production.
	EnterCastableExpr(c *CastableExprContext)

	// EnterCastExpr is called when entering the castExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterUnaryExpr is called when entering the unaryExpr production.
	EnterUnaryExpr(c *UnaryExprContext)

	// EnterSimpleMapExpr is called when entering the simpleMapExpr production.
	EnterSimpleMapExpr(c *SimpleMapExprContext)

	// EnterPostfixExpr is called when entering the postfixExpr production.
	EnterPostfixExpr(c *PostfixExprContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterObjectLookup is called when entering the objectLookup production.
	EnterObjectLookup(c *ObjectLookupContext)

	// EnterArrayLookup is called when entering the arrayLookup production.
	EnterArrayLookup(c *ArrayLookupContext)

	// EnterArrayUnboxing is called when entering the arrayUnboxing production.
	EnterArrayUnboxing(c *ArrayUnboxingContext)

	// EnterPrimaryExpr is called when entering the primaryExpr production.
	EnterPrimaryExpr(c *PrimaryExprContext)

	// EnterVarRef is called when entering the varRef production.
	EnterVarRef(c *VarRefContext)

	// EnterParenthesizedExpr is called when entering the parenthesizedExpr production.
	EnterParenthesizedExpr(c *ParenthesizedExprContext)

	// EnterContextItemExpr is called when entering the contextItemExpr production.
	EnterContextItemExpr(c *ContextItemExprContext)

	// EnterOrderedExpr is called when entering the orderedExpr production.
	EnterOrderedExpr(c *OrderedExprContext)

	// EnterUnorderedExpr is called when entering the unorderedExpr production.
	EnterUnorderedExpr(c *UnorderedExprContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterObjectConstructor is called when entering the objectConstructor production.
	EnterObjectConstructor(c *ObjectConstructorContext)

	// EnterPairConstructor is called when entering the pairConstructor production.
	EnterPairConstructor(c *PairConstructorContext)

	// EnterArrayConstructor is called when entering the arrayConstructor production.
	EnterArrayConstructor(c *ArrayConstructorContext)

	// EnterSequenceType is called when entering the sequenceType production.
	EnterSequenceType(c *SequenceTypeContext)

	// EnterItemType is called when entering the itemType production.
	EnterItemType(c *ItemTypeContext)

	// EnterJsonItemTest is called when entering the jsonItemTest production.
	EnterJsonItemTest(c *JsonItemTestContext)

	// EnterAtomicType is called when entering the atomicType production.
	EnterAtomicType(c *AtomicTypeContext)

	// EnterUriLiteral is called when entering the uriLiteral production.
	EnterUriLiteral(c *UriLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterBooleanLiteral is called when entering the booleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterNullLiteral is called when entering the nullLiteral production.
	EnterNullLiteral(c *NullLiteralContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitMainModule is called when exiting the mainModule production.
	ExitMainModule(c *MainModuleContext)

	// ExitLibraryModule is called when exiting the libraryModule production.
	ExitLibraryModule(c *LibraryModuleContext)

	// ExitProlog is called when exiting the prolog production.
	ExitProlog(c *PrologContext)

	// ExitDefaultCollationDecl is called when exiting the defaultCollationDecl production.
	ExitDefaultCollationDecl(c *DefaultCollationDeclContext)

	// ExitOrderingModeDecl is called when exiting the orderingModeDecl production.
	ExitOrderingModeDecl(c *OrderingModeDeclContext)

	// ExitEmptyOrderDecl is called when exiting the emptyOrderDecl production.
	ExitEmptyOrderDecl(c *EmptyOrderDeclContext)

	// ExitDecimalFormatDecl is called when exiting the decimalFormatDecl production.
	ExitDecimalFormatDecl(c *DecimalFormatDeclContext)

	// ExitDfPropertyName is called when exiting the dfPropertyName production.
	ExitDfPropertyName(c *DfPropertyNameContext)

	// ExitModuleImport is called when exiting the moduleImport production.
	ExitModuleImport(c *ModuleImportContext)

	// ExitVarDecl is called when exiting the varDecl production.
	ExitVarDecl(c *VarDeclContext)

	// ExitFunctionDecl is called when exiting the functionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitExprSingle is called when exiting the exprSingle production.
	ExitExprSingle(c *ExprSingleContext)

	// ExitFlowrExpr is called when exiting the flowrExpr production.
	ExitFlowrExpr(c *FlowrExprContext)

	// ExitForClause is called when exiting the forClause production.
	ExitForClause(c *ForClauseContext)

	// ExitLetClause is called when exiting the letClause production.
	ExitLetClause(c *LetClauseContext)

	// ExitCountClause is called when exiting the countClause production.
	ExitCountClause(c *CountClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitQuantifiedExpr is called when exiting the quantifiedExpr production.
	ExitQuantifiedExpr(c *QuantifiedExprContext)

	// ExitSwitchExpr is called when exiting the switchExpr production.
	ExitSwitchExpr(c *SwitchExprContext)

	// ExitSwitchCaseClause is called when exiting the switchCaseClause production.
	ExitSwitchCaseClause(c *SwitchCaseClauseContext)

	// ExitTypeswitchExpr is called when exiting the typeswitchExpr production.
	ExitTypeswitchExpr(c *TypeswitchExprContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitTryCatchExpr is called when exiting the tryCatchExpr production.
	ExitTryCatchExpr(c *TryCatchExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitComparisonExpr2 is called when exiting the comparisonExpr2 production.
	ExitComparisonExpr2(c *ComparisonExpr2Context)

	// ExitComparisonExpr is called when exiting the comparisonExpr production.
	ExitComparisonExpr(c *ComparisonExprContext)

	// ExitStringConcatExpr is called when exiting the stringConcatExpr production.
	ExitStringConcatExpr(c *StringConcatExprContext)

	// ExitRangeExpr is called when exiting the rangeExpr production.
	ExitRangeExpr(c *RangeExprContext)

	// ExitAdditiveExpr is called when exiting the additiveExpr production.
	ExitAdditiveExpr(c *AdditiveExprContext)

	// ExitMultiplicativeExpr is called when exiting the multiplicativeExpr production.
	ExitMultiplicativeExpr(c *MultiplicativeExprContext)

	// ExitInstanceofExpr is called when exiting the instanceofExpr production.
	ExitInstanceofExpr(c *InstanceofExprContext)

	// ExitTreatExpr is called when exiting the treatExpr production.
	ExitTreatExpr(c *TreatExprContext)

	// ExitCastableExpr is called when exiting the castableExpr production.
	ExitCastableExpr(c *CastableExprContext)

	// ExitCastExpr is called when exiting the castExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitUnaryExpr is called when exiting the unaryExpr production.
	ExitUnaryExpr(c *UnaryExprContext)

	// ExitSimpleMapExpr is called when exiting the simpleMapExpr production.
	ExitSimpleMapExpr(c *SimpleMapExprContext)

	// ExitPostfixExpr is called when exiting the postfixExpr production.
	ExitPostfixExpr(c *PostfixExprContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitObjectLookup is called when exiting the objectLookup production.
	ExitObjectLookup(c *ObjectLookupContext)

	// ExitArrayLookup is called when exiting the arrayLookup production.
	ExitArrayLookup(c *ArrayLookupContext)

	// ExitArrayUnboxing is called when exiting the arrayUnboxing production.
	ExitArrayUnboxing(c *ArrayUnboxingContext)

	// ExitPrimaryExpr is called when exiting the primaryExpr production.
	ExitPrimaryExpr(c *PrimaryExprContext)

	// ExitVarRef is called when exiting the varRef production.
	ExitVarRef(c *VarRefContext)

	// ExitParenthesizedExpr is called when exiting the parenthesizedExpr production.
	ExitParenthesizedExpr(c *ParenthesizedExprContext)

	// ExitContextItemExpr is called when exiting the contextItemExpr production.
	ExitContextItemExpr(c *ContextItemExprContext)

	// ExitOrderedExpr is called when exiting the orderedExpr production.
	ExitOrderedExpr(c *OrderedExprContext)

	// ExitUnorderedExpr is called when exiting the unorderedExpr production.
	ExitUnorderedExpr(c *UnorderedExprContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitObjectConstructor is called when exiting the objectConstructor production.
	ExitObjectConstructor(c *ObjectConstructorContext)

	// ExitPairConstructor is called when exiting the pairConstructor production.
	ExitPairConstructor(c *PairConstructorContext)

	// ExitArrayConstructor is called when exiting the arrayConstructor production.
	ExitArrayConstructor(c *ArrayConstructorContext)

	// ExitSequenceType is called when exiting the sequenceType production.
	ExitSequenceType(c *SequenceTypeContext)

	// ExitItemType is called when exiting the itemType production.
	ExitItemType(c *ItemTypeContext)

	// ExitJsonItemTest is called when exiting the jsonItemTest production.
	ExitJsonItemTest(c *JsonItemTestContext)

	// ExitAtomicType is called when exiting the atomicType production.
	ExitAtomicType(c *AtomicTypeContext)

	// ExitUriLiteral is called when exiting the uriLiteral production.
	ExitUriLiteral(c *UriLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitBooleanLiteral is called when exiting the booleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitNullLiteral is called when exiting the nullLiteral production.
	ExitNullLiteral(c *NullLiteralContext)
}
