package urn

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/leodido/go-urn/grammar"
)

func TestListener(t *testing.T) {

	inputs := []string{
	//"URN:1ab:xxx",
	}

	fmt.Println("---")
	for _, input := range inputs {
		stream := antlr.NewInputStream(input)
		lexer := grammar.NewUrnLexer(stream)
		tokens := antlr.NewCommonTokenStream(lexer, 0)

		parser := grammar.NewUrnParser(tokens)
		parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		parser.BuildParseTrees = true

		/*
			urn := URN{}
			parser.AddParseListener(NewListener(&urn))
		*/

		parser.Urn()

		//fmt.Printf("getstruct=> %+v\n", tree.GetStructure())
		//fmt.Println(tree.ToStringTree(nil, parser))

		// fmt.Printf("%s => %+v\n", input, urn)

		fmt.Println("---")
	}
}
