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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 25, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 2, 2, 6, 
	3, 3, 5, 4, 7, 5, 9, 6, 3, 2, 2, 2, 24, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 3, 11, 3, 2, 2, 2, 5, 15, 3, 2, 
	2, 2, 7, 19, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2, 11, 12, 7, 119, 2, 2, 12, 13, 
	7, 116, 2, 2, 13, 14, 7, 112, 2, 2, 14, 4, 3, 2, 2, 2, 15, 16, 7, 80, 2, 
	2, 16, 17, 7, 75, 2, 2, 17, 18, 7, 70, 2, 2, 18, 6, 3, 2, 2, 2, 19, 20, 
	7, 80, 2, 2, 20, 21, 7, 85, 2, 2, 21, 22, 7, 85, 2, 2, 22, 8, 3, 2, 2, 
	2, 23, 24, 7, 60, 2, 2, 24, 10, 3, 2, 2, 2, 3, 2, 2,
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
	"", "'urn'", "'NID'", "'NSS'", "':'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "Colon",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "Colon",
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
	UrnLexerT__0 = 1
	UrnLexerT__1 = 2
	UrnLexerT__2 = 3
	UrnLexerColon = 4
)

