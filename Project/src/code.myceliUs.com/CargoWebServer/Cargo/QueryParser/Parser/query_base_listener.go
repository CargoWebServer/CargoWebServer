// Code generated from Query.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseQueryListener is a complete listener for a parse tree produced by QueryParser.
type BaseQueryListener struct{}

var _ QueryListener = &BaseQueryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseQueryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseQueryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseQueryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseQueryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseQueryListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseQueryListener) ExitQuery(ctx *QueryContext) {}

// EnterBracketExpression is called when production BracketExpression is entered.
func (s *BaseQueryListener) EnterBracketExpression(ctx *BracketExpressionContext) {}

// ExitBracketExpression is called when production BracketExpression is exited.
func (s *BaseQueryListener) ExitBracketExpression(ctx *BracketExpressionContext) {}

// EnterAndExpression is called when production AndExpression is entered.
func (s *BaseQueryListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production AndExpression is exited.
func (s *BaseQueryListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterPredicateExpression is called when production PredicateExpression is entered.
func (s *BaseQueryListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {}

// ExitPredicateExpression is called when production PredicateExpression is exited.
func (s *BaseQueryListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {}

// EnterOrExpression is called when production OrExpression is entered.
func (s *BaseQueryListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production OrExpression is exited.
func (s *BaseQueryListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseQueryListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseQueryListener) ExitReference(ctx *ReferenceContext) {}

// EnterElement is called when production element is entered.
func (s *BaseQueryListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseQueryListener) ExitElement(ctx *ElementContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseQueryListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseQueryListener) ExitPredicate(ctx *PredicateContext) {}

// EnterOperator is called when production operator is entered.
func (s *BaseQueryListener) EnterOperator(ctx *OperatorContext) {}

// ExitOperator is called when production operator is exited.
func (s *BaseQueryListener) ExitOperator(ctx *OperatorContext) {}

// EnterIntegerValue is called when production IntegerValue is entered.
func (s *BaseQueryListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production IntegerValue is exited.
func (s *BaseQueryListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterDoubleValue is called when production DoubleValue is entered.
func (s *BaseQueryListener) EnterDoubleValue(ctx *DoubleValueContext) {}

// ExitDoubleValue is called when production DoubleValue is exited.
func (s *BaseQueryListener) ExitDoubleValue(ctx *DoubleValueContext) {}

// EnterBooleanValue is called when production BooleanValue is entered.
func (s *BaseQueryListener) EnterBooleanValue(ctx *BooleanValueContext) {}

// ExitBooleanValue is called when production BooleanValue is exited.
func (s *BaseQueryListener) ExitBooleanValue(ctx *BooleanValueContext) {}

// EnterStringValue is called when production StringValue is entered.
func (s *BaseQueryListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production StringValue is exited.
func (s *BaseQueryListener) ExitStringValue(ctx *StringValueContext) {}
