// Generated from /home/leodido/workspaces/go/src/github.com/leodido/go-urn/grammar/Urn.g4 by ANTLR 4.7.

package grammar
import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 5, 89, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 3, 
	3, 3, 5, 3, 38, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 
	3, 6, 6, 6, 49, 10, 6, 13, 6, 14, 6, 50, 3, 7, 3, 7, 6, 7, 55, 10, 7, 13, 
	7, 14, 7, 56, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 5, 10, 66, 10, 
	10, 3, 11, 3, 11, 5, 11, 70, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 
	5, 13, 77, 10, 13, 3, 14, 3, 14, 5, 14, 81, 10, 14, 3, 15, 3, 15, 3, 16, 
	3, 16, 3, 16, 5, 16, 88, 10, 16, 2, 2, 17, 3, 3, 5, 4, 7, 5, 9, 2, 11, 
	2, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2, 31, 2, 
	3, 2, 11, 4, 2, 87, 87, 119, 119, 4, 2, 84, 84, 116, 116, 4, 2, 80, 80, 
	112, 112, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 4, 2, 42, 43, 45, 46, 11, 
	2, 35, 35, 38, 38, 41, 41, 44, 44, 48, 48, 61, 61, 63, 63, 66, 66, 97, 
	97, 4, 2, 67, 72, 99, 104, 6, 2, 37, 37, 39, 39, 49, 49, 65, 65, 2, 87, 
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 
	5, 37, 3, 2, 2, 2, 7, 39, 3, 2, 2, 2, 9, 41, 3, 2, 2, 2, 11, 45, 3, 2, 
	2, 2, 13, 54, 3, 2, 2, 2, 15, 58, 3, 2, 2, 2, 17, 62, 3, 2, 2, 2, 19, 65, 
	3, 2, 2, 2, 21, 69, 3, 2, 2, 2, 23, 71, 3, 2, 2, 2, 25, 76, 3, 2, 2, 2, 
	27, 80, 3, 2, 2, 2, 29, 82, 3, 2, 2, 2, 31, 87, 3, 2, 2, 2, 33, 34, 5, 
	9, 5, 2, 34, 4, 3, 2, 2, 2, 35, 38, 5, 11, 6, 2, 36, 38, 5, 13, 7, 2, 37, 
	35, 3, 2, 2, 2, 37, 36, 3, 2, 2, 2, 38, 6, 3, 2, 2, 2, 39, 40, 7, 60, 2, 
	2, 40, 8, 3, 2, 2, 2, 41, 42, 9, 2, 2, 2, 42, 43, 9, 3, 2, 2, 43, 44, 9, 
	4, 2, 2, 44, 10, 3, 2, 2, 2, 45, 48, 5, 21, 11, 2, 46, 49, 5, 21, 11, 2, 
	47, 49, 5, 23, 12, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 
	2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 12, 3, 2, 2, 2, 52, 
	55, 5, 31, 16, 2, 53, 55, 5, 15, 8, 2, 54, 52, 3, 2, 2, 2, 54, 53, 3, 2, 
	2, 2, 55, 56, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 14, 
	3, 2, 2, 2, 58, 59, 7, 39, 2, 2, 59, 60, 5, 27, 14, 2, 60, 61, 5, 27, 14, 
	2, 61, 16, 3, 2, 2, 2, 62, 63, 9, 5, 2, 2, 63, 18, 3, 2, 2, 2, 64, 66, 
	9, 6, 2, 2, 65, 64, 3, 2, 2, 2, 66, 20, 3, 2, 2, 2, 67, 70, 5, 17, 9, 2, 
	68, 70, 5, 19, 10, 2, 69, 67, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 22, 3, 
	2, 2, 2, 71, 72, 7, 47, 2, 2, 72, 24, 3, 2, 2, 2, 73, 77, 9, 7, 2, 2, 74, 
	77, 5, 23, 12, 2, 75, 77, 9, 8, 2, 2, 76, 73, 3, 2, 2, 2, 76, 74, 3, 2, 
	2, 2, 76, 75, 3, 2, 2, 2, 77, 26, 3, 2, 2, 2, 78, 81, 5, 17, 9, 2, 79, 
	81, 9, 9, 2, 2, 80, 78, 3, 2, 2, 2, 80, 79, 3, 2, 2, 2, 81, 28, 3, 2, 2, 
	2, 82, 83, 9, 10, 2, 2, 83, 30, 3, 2, 2, 2, 84, 88, 5, 21, 11, 2, 85, 88, 
	5, 25, 13, 2, 86, 88, 5, 29, 15, 2, 87, 84, 3, 2, 2, 2, 87, 85, 3, 2, 2, 
	2, 87, 86, 3, 2, 2, 2, 88, 32, 3, 2, 2, 2, 13, 2, 37, 48, 50, 54, 56, 65, 
	69, 76, 80, 87, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "':'",
}

var lexerSymbolicNames = []string{
	"", "Urn", "Part", "Colon",
}

var lexerRuleNames = []string{
	"Urn", "Part", "Colon", "URN", "IDENTIFIER", "CHARS", "HEX", "NUMBER", 
	"LETTER", "ALPHA_NUMERIC", "HYPHEN", "OTHER", "HEX_DIGIT", "RESERVED", 
	"TRANSLATION",
}

type UrnLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewUrnLexer(input antlr.CharStream) *UrnLexer {

	l := new(UrnLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Urn.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// UrnLexer tokens.
const (
	UrnLexerUrn = 1
	UrnLexerPart = 2
	UrnLexerColon = 3
)

