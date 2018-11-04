package QueryParser

import (
	"code.myceliUs.com/CargoWebServer/Cargo/QueryParser/Parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

/**
 * Simply run the query.
 */
func Parse(query string, listener parser.QueryListener) {

	// Setup the input
	is := antlr.NewInputStream(query)

	// Create the Lexer
	lexer := parser.NewQueryLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewQueryParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Query())
}
