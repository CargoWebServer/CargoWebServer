// Code generated from Query.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Query

import "github.com/antlr/antlr4/runtime/Go/antlr"

// QueryListener is a complete listener for a parse tree produced by QueryParser.
type QueryListener interface {
	antlr.ParseTreeListener

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterBracketExpression is called when entering the BracketExpression production.
	EnterBracketExpression(c *BracketExpressionContext)

	// EnterAndExpression is called when entering the AndExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterPredicateExpression is called when entering the PredicateExpression production.
	EnterPredicateExpression(c *PredicateExpressionContext)

	// EnterOrExpression is called when entering the OrExpression production.
	EnterOrExpression(c *OrExpressionContext)

	// EnterReference is called when entering the reference production.
	EnterReference(c *ReferenceContext)

	// EnterElement is called when entering the element production.
	EnterElement(c *ElementContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterOperator is called when entering the operator production.
	EnterOperator(c *OperatorContext)

	// EnterIntegerValue is called when entering the IntegerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterDoubleValue is called when entering the DoubleValue production.
	EnterDoubleValue(c *DoubleValueContext)

	// EnterBooleanValue is called when entering the BooleanValue production.
	EnterBooleanValue(c *BooleanValueContext)

	// EnterStringValue is called when entering the StringValue production.
	EnterStringValue(c *StringValueContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitBracketExpression is called when exiting the BracketExpression production.
	ExitBracketExpression(c *BracketExpressionContext)

	// ExitAndExpression is called when exiting the AndExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitPredicateExpression is called when exiting the PredicateExpression production.
	ExitPredicateExpression(c *PredicateExpressionContext)

	// ExitOrExpression is called when exiting the OrExpression production.
	ExitOrExpression(c *OrExpressionContext)

	// ExitReference is called when exiting the reference production.
	ExitReference(c *ReferenceContext)

	// ExitElement is called when exiting the element production.
	ExitElement(c *ElementContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitOperator is called when exiting the operator production.
	ExitOperator(c *OperatorContext)

	// ExitIntegerValue is called when exiting the IntegerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitDoubleValue is called when exiting the DoubleValue production.
	ExitDoubleValue(c *DoubleValueContext)

	// ExitBooleanValue is called when exiting the BooleanValue production.
	ExitBooleanValue(c *BooleanValueContext)

	// ExitStringValue is called when exiting the StringValue production.
	ExitStringValue(c *StringValueContext)
}
