package jsoniq

import (
	"log"

	"code.myceliUs.com/GoJSONiq/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type jsoniqInterpreter struct {
}

// The
func NewInterpreter() *jsoniqInterpreter {
	interpreter := new(jsoniqInterpreter)
	return interpreter
}

/**
 * Run a JSONiq program.
 * let The let function callback with items to be set.
 */
func (jsoniqInterpreter) Run(code string, callback func(...interface{}) error) {
	is := antlr.NewInputStream(code)
	lexer := parser.NewjsoniqLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewjsoniqParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&jsoniqListener{}, p.Program())

	// The query proecessing is done.
	callback("Done")
}

func (jsoniqInterpreter) RunSingleExpression(code string) {

}

/**
 * The jsoniq listener.
 */
type jsoniqListener struct {
	*parser.BasejsoniqListener
}

////////////////////////////////////////////////////////////////////////////////
// Enter event
////////////////////////////////////////////////////////////////////////////////

// Enter the program...
func (s *jsoniqListener) EnterProgram(ctx *parser.ProgramContext) {
	log.Println("---> enter program!")
}

// EnterMainModule is called when production mainModule is entered.
func (s *jsoniqListener) EnterMainModule(ctx *parser.MainModuleContext) {
	log.Println("---> enter main module!")
}

func (s *jsoniqListener) EnterProlog(ctx *parser.PrologContext) {
	log.Println("---> enter prolog!")
}

func (s *jsoniqListener) EnterExprSingle(c *parser.ExprSingleContext) {
	log.Println("---> enter expr single! ", c.GetText())
}

func (s *jsoniqListener) EnterAtomicType(c *parser.AtomicTypeContext) {
	log.Println("---> enter function call! ", c.GetText())
}

func (s *jsoniqListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	log.Println("---> enter function call! ", ctx.NCName(0))
}

// EnterFunctionDecl is called when production functionDecl is entered.
func (s *jsoniqListener) EnterFunctionDecl(ctx *parser.FunctionDeclContext) {
	log.Println("---> enter function declaration!")
}

// Enter property definition.
func (s *jsoniqListener) EnterForClause(ctx *parser.ForClauseContext) {
	log.Println("---> enter for: ", ctx)
}

////////////////////////////////////////////////////////////////////////////////
// Exit event
////////////////////////////////////////////////////////////////////////////////
func (s *jsoniqListener) ExitProgram(ctx *parser.ProgramContext) {
	log.Println("---> exit program!")
}
