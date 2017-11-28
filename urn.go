package urn

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/leodido/go-urn/grammar"
)

// URN represents an Uniform Resource Name.
//
// The general form represented is:
//
//	urn:<id>:<ss>
//
// Details at https://tools.ietf.org/html/rfc2141
type URN struct {
	ID   string // Namespace identifier
	SS   string // Namespace specific string
	tree string
}

// Error describes an error and the input that caused it.
type Error struct {
	Source string
	Column int
	Detail string
}

func (e *Error) Error() string {
	return "\"" + e.Source + "\" causes syntax error at character " + strconv.Itoa(e.Column) + ", " + e.Detail
}

type errorListener struct {
	*antlr.DefaultErrorListener
	input string
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(&Error{
		Source: l.input,
		Column: column,
		Detail: msg,
	})
}

// Parse is ...
func Parse(u string) (*URN, error) {
	urn, err := parse(u)
	if err != nil {
		return nil, err
	}

	return urn, err
}

func parse(u string) (urn *URN, err error) {
	errl := new(errorListener)
	errl.DefaultErrorListener = new(antlr.DefaultErrorListener)
	errl.input = u

	stream := antlr.NewInputStream(u)
	lexer := grammar.NewUrnLexer(stream)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parser := grammar.NewUrnParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errl)
	// parser.BuildParseTrees = false // (todo) > ?

	defer func() {
		if r := recover(); r != nil {
			urn = nil
			err = r.(error)
		}
	}()

	urn = &URN{}
	parser.AddParseListener(NewListener(urn))
	parser.Urn()

	return
}

// String reassembles the URN into a valid URN string.
//
// This requires both ID and SS fields to be non-empty.
// Otherwise it returns an empty string.
func (u *URN) String() string {
	var res string
	if u.ID != "" && u.SS != "" {
		res = "urn:" + u.ID + ":" + u.SS
	}

	return res
}

// Tree returns a string representation of the resulting concrete syntax tree.
func (u *URN) Tree() string {
	return u.tree
}
