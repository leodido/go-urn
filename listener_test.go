package urn

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/leodido/go-urn/grammar"
)

func TestX(t *testing.T) {
	inFile, _ := os.Open("testdata/00.urn")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input := scanner.Text()

		stream := antlr.NewInputStream(input)
		lexer := grammar.NewUrnLexer(stream)
		tokens := antlr.NewCommonTokenStream(lexer, 0)

		parser := grammar.NewUrnParser(tokens)
		// parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

		urn := Urn{}
		parser.AddParseListener(NewListener(&urn))

		// parser.BuildParseTrees = true

		parser.Urn()

		fmt.Printf("%+v\n", urn)
	}
}
