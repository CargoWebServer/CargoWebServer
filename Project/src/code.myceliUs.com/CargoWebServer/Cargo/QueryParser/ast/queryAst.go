package ast

import (
	"log"
	"strconv"
	"strings"

	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/token"
	"code.myceliUs.com/Utility"
)

type (
	StmtList []Stmt
	Stmt     string
)

/**
 * The asyntaxique tree use to run the query over the the search engine.
 */
type QueryAst struct {

	// Operator.
	operator string

	// Keep the list of sub-queries.
	subQueries []*QueryAst

	// The comparator...
	comparator string

	// The object id.
	objectId []string

	// The value to find.
	value interface{} // can be one of string float64 bool and int64.
}

/**
 * If the tree contain sub values...
 */
func (this *QueryAst) IsComposite() bool {
	return len(this.operator) > 0
}

/**
 * Return the sub-queries
 */
func (this *QueryAst) GetSubQueries() (q1 *QueryAst, operator string, q2 *QueryAst) {
	return this.subQueries[0], this.operator, this.subQueries[1]
}

/**
 * Return the informations related to an expression.
 */
func (this *QueryAst) GetExpression() (typeName string, fieldName string, comparator string, values interface{}) {
	for i := 0; i < len(this.subQueries[0].objectId)-1; i++ {
		if i == 0 {
			typeName += this.subQueries[0].objectId[i]
		} else {
			typeName += "." + this.subQueries[0].objectId[i]
		}
	}

	fieldName = this.subQueries[0].objectId[len(this.subQueries[0].objectId)-1]

	return typeName, fieldName, this.comparator, this.subQueries[1].value
}

func (this *QueryAst) Print(level int) {
	arrow := "--"
	for i := 0; i < level; i++ {
		arrow += arrow
	}
	arrow += ">"
	log.Println(level, arrow+" {")
	if len(this.comparator) > 0 {
		log.Println(level, arrow+" "+this.comparator)
	} else if len(this.operator) > 0 {
		log.Println(level, arrow+" "+this.operator)
	} else if this.value != nil {
		log.Println(level, arrow+" ", this.value)
	}

	for i := 0; i < len(this.subQueries); i++ {
		this.subQueries[i].Print(level + 1)
	}

	log.Println(level, arrow+" }")
}

/**
 * Create a new asyntaxique parse tree.
 */
func NewQueryAst(val interface{}) (*QueryAst, error) {
	return val.(*QueryAst), nil
}

func NewExpressionAst(val0 interface{}, val1 interface{}, val2 interface{}) (*QueryAst, error) {
	ast := new(QueryAst)
	ast.subQueries = make([]*QueryAst, 0)
	ast.comparator = string(val1.(*token.Token).Lit)

	child0 := val0.(*QueryAst)

	ast.subQueries = append(ast.subQueries, child0)

	child1 := val2.(*QueryAst)
	ast.subQueries = append(ast.subQueries, child1)

	return ast, nil
}

/**
 * Append a new query.
 */
func AppendQueryAst(val0 interface{}, val1 interface{}, val2 interface{}) (*QueryAst, error) {
	ast := new(QueryAst)
	ast.subQueries = make([]*QueryAst, 0)
	ast.operator = string(val1.(*token.Token).Lit)

	child0 := val0.(*QueryAst)
	ast.subQueries = append(ast.subQueries, child0)

	child1 := val2.(*QueryAst)
	ast.subQueries = append(ast.subQueries, child1)

	return ast, nil
}

////////////////////////////////////////////////////////////////////////////////
//  Creation of leaf values
////////////////////////////////////////////////////////////////////////////////
/**
 * integer.
 */
func NewIntegerValue(val interface{}) (*QueryAst, error) {
	intVal, err := strconv.ParseInt(string(val.(*token.Token).Lit), 10, 64)
	ast := new(QueryAst)
	ast.value = intVal
	return ast, err
}

/**
 * float
 */
func NewFloatValue(val interface{}) (*QueryAst, error) {
	floatVal, err := strconv.ParseFloat(string(val.(*token.Token).Lit), 64)
	ast := new(QueryAst)
	ast.value = floatVal
	return ast, err
}

/**
 * string
 */
func NewStringValue(val interface{}) (*QueryAst, error) {
	strVal := string(val.(*token.Token).Lit)

	// If the string dosent contain space...
	// other wize it will be considere a phrase.
	if strings.HasPrefix(strVal, "\"") {
		strVal = strings.TrimPrefix(strVal, "\"")
	}
	if strings.HasSuffix(strVal, "\"") {
		strVal = strings.TrimSuffix(strVal, "\"")
	}

	strVal = Utility.RemoveAccent(strVal)

	ast := new(QueryAst)
	ast.value = strVal
	return ast, nil
}

/**
 * string
 */
func NewNullValue(val interface{}) (*QueryAst, error) {
	// The value must be equal to 'null'...
	ast := new(QueryAst)
	ast.value = nil // nil value.
	return ast, nil
}

/**
 * indentifier
 */
func NewIdValue(val0 interface{}, val1 interface{}) (*QueryAst, error) {
	strVal0 := string(val0.(*token.Token).Lit)
	strVal1 := string(val1.(*token.Token).Lit)
	ast := new(QueryAst)
	ast.objectId = make([]string, 2)
	ast.objectId[0] = strVal0
	ast.objectId[1] = strVal1
	return ast, nil
}

/**
 * Append a new value inside the identifier.
 */
func AppendIdValue(val0 interface{}, val1 interface{}) (*QueryAst, error) {
	ast := val0.(*QueryAst)
	ast.objectId = append(ast.objectId, string(val1.(*token.Token).Lit))

	return ast, nil
}

/**
 * Boolean
 */
func NewBoolValue(val interface{}) (*QueryAst, error) {
	boolVal, err := strconv.ParseBool(string(val.(*token.Token).Lit))
	ast := new(QueryAst)
	ast.value = boolVal
	return ast, err
}

/**
 * Regex
 */
func NewRegexExpr(val interface{}) (*QueryAst, error) {
	strVal := string(val.(*token.Token).Lit)
	ast := new(QueryAst)
	ast.value = strVal

	return ast, nil
}
