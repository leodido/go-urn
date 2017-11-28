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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 7, 90, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 
	3, 2, 3, 3, 3, 3, 5, 3, 40, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 6, 8, 55, 10, 8, 13, 8, 14, 8, 
	56, 3, 9, 3, 9, 6, 9, 61, 10, 9, 13, 9, 14, 9, 62, 3, 10, 3, 10, 3, 10, 
	3, 10, 3, 11, 3, 11, 3, 12, 5, 12, 72, 10, 12, 3, 13, 3, 13, 5, 13, 76, 
	10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 82, 10, 15, 3, 16, 3, 16, 3, 
	17, 3, 17, 3, 17, 5, 17, 89, 10, 17, 2, 2, 18, 3, 3, 5, 4, 7, 5, 9, 6, 
	11, 7, 13, 2, 15, 2, 17, 2, 19, 2, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2, 31, 
	2, 33, 2, 3, 2, 10, 4, 2, 87, 87, 119, 119, 4, 2, 84, 84, 116, 116, 4, 
	2, 80, 80, 112, 112, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 10, 2, 35, 35, 
	38, 38, 41, 46, 48, 48, 61, 61, 63, 63, 66, 66, 97, 97, 4, 2, 67, 72, 99, 
	104, 6, 2, 37, 37, 39, 39, 49, 49, 65, 65, 2, 87, 2, 3, 3, 2, 2, 2, 2, 
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 3, 
	35, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2, 7, 41, 3, 2, 2, 2, 9, 43, 3, 2, 2, 2, 
	11, 45, 3, 2, 2, 2, 13, 47, 3, 2, 2, 2, 15, 51, 3, 2, 2, 2, 17, 60, 3, 
	2, 2, 2, 19, 64, 3, 2, 2, 2, 21, 68, 3, 2, 2, 2, 23, 71, 3, 2, 2, 2, 25, 
	75, 3, 2, 2, 2, 27, 77, 3, 2, 2, 2, 29, 81, 3, 2, 2, 2, 31, 83, 3, 2, 2, 
	2, 33, 88, 3, 2, 2, 2, 35, 36, 5, 13, 7, 2, 36, 4, 3, 2, 2, 2, 37, 40, 
	5, 15, 8, 2, 38, 40, 5, 17, 9, 2, 39, 37, 3, 2, 2, 2, 39, 38, 3, 2, 2, 
	2, 40, 6, 3, 2, 2, 2, 41, 42, 7, 60, 2, 2, 42, 8, 3, 2, 2, 2, 43, 44, 7, 
	47, 2, 2, 44, 10, 3, 2, 2, 2, 45, 46, 7, 34, 2, 2, 46, 12, 3, 2, 2, 2, 
	47, 48, 9, 2, 2, 2, 48, 49, 9, 3, 2, 2, 49, 50, 9, 4, 2, 2, 50, 14, 3, 
	2, 2, 2, 51, 54, 5, 25, 13, 2, 52, 55, 5, 25, 13, 2, 53, 55, 5, 9, 5, 2, 
	54, 52, 3, 2, 2, 2, 54, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 54, 3, 
	2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 16, 3, 2, 2, 2, 58, 61, 5, 33, 17, 2, 
	59, 61, 5, 19, 10, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 62, 3, 
	2, 2, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 18, 3, 2, 2, 2, 64, 
	65, 7, 39, 2, 2, 65, 66, 5, 29, 15, 2, 66, 67, 5, 29, 15, 2, 67, 20, 3, 
	2, 2, 2, 68, 69, 9, 5, 2, 2, 69, 22, 3, 2, 2, 2, 70, 72, 9, 6, 2, 2, 71, 
	70, 3, 2, 2, 2, 72, 24, 3, 2, 2, 2, 73, 76, 5, 21, 11, 2, 74, 76, 5, 23, 
	12, 2, 75, 73, 3, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 26, 3, 2, 2, 2, 77, 
	78, 9, 7, 2, 2, 78, 28, 3, 2, 2, 2, 79, 82, 5, 21, 11, 2, 80, 82, 9, 8, 
	2, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 30, 3, 2, 2, 2, 83, 84, 
	9, 9, 2, 2, 84, 32, 3, 2, 2, 2, 85, 89, 5, 27, 14, 2, 86, 89, 5, 31, 16, 
	2, 87, 89, 5, 25, 13, 2, 88, 85, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 87, 
	3, 2, 2, 2, 89, 34, 3, 2, 2, 2, 12, 2, 39, 54, 56, 60, 62, 71, 75, 81, 
	88, 2,
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
	"", "", "", "':'", "'-'", "' '",
}

var lexerSymbolicNames = []string{
	"", "Urn", "Part", "Colon", "Hyphen", "Whitespace",
}

var lexerRuleNames = []string{
	"Urn", "Part", "Colon", "Hyphen", "Whitespace", "URN", "IDENTIFIER", "CHARS", 
	"HEX", "NUMBER", "LETTER", "ALPHA_NUMERIC", "OTHER", "HEX_DIGIT", "RESERVED", 
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
	UrnLexerHyphen = 4
	UrnLexerWhitespace = 5
)

