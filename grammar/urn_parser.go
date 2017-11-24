// Generated from /home/leodido/workspaces/go/src/github.com/leodido/go-urn/grammar/Urn.g4 by ANTLR 4.7.

package grammar // Urn
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)





// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 20, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 16, 2, 8, 3, 
	2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 17, 3, 2, 2, 2, 8, 9, 7, 3, 2, 2, 9, 10, 
	7, 6, 2, 2, 10, 11, 5, 4, 3, 2, 11, 12, 7, 6, 2, 2, 12, 13, 5, 6, 4, 2, 
	13, 14, 7, 2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 16, 7, 4, 2, 2, 16, 5, 3, 2, 
	2, 2, 17, 18, 7, 5, 2, 2, 18, 7, 3, 2, 2, 2, 2,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'urn'", "'NID'", "'NSS'", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "Colon",
}

var ruleNames = []string{
	"urn", "iD", "sS",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type UrnParser struct {
	*antlr.BaseParser
}

func NewUrnParser(input antlr.TokenStream) *UrnParser {
	this := new(UrnParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Urn.g4"

	return this
}





// UrnParser tokens.
const (
	UrnParserEOF = antlr.TokenEOF
	UrnParserT__0 = 1
	UrnParserT__1 = 2
	UrnParserT__2 = 3
	UrnParserColon = 4
)

// UrnParser rules.
const (
	UrnParserRULE_urn = 0
	UrnParserRULE_iD = 1
	UrnParserRULE_sS = 2
)

// IUrnContext is an interface to support dynamic dispatch.
type IUrnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUrnContext differentiates from other interfaces.
	IsUrnContext()
}

type UrnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrnContext() *UrnContext {
	var p = new(UrnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UrnParserRULE_urn
	return p
}

func (*UrnContext) IsUrnContext() {}

func NewUrnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrnContext {
	var p = new(UrnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UrnParserRULE_urn

	return p
}

func (s *UrnContext) GetParser() antlr.Parser { return s.parser }

func (s *UrnContext) AllColon() []antlr.TerminalNode {
	return s.GetTokens(UrnParserColon)
}

func (s *UrnContext) Colon(i int) antlr.TerminalNode {
	return s.GetToken(UrnParserColon, i)
}

func (s *UrnContext) ID() IIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIDContext)
}

func (s *UrnContext) SS() ISSContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISSContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISSContext)
}

func (s *UrnContext) EOF() antlr.TerminalNode {
	return s.GetToken(UrnParserEOF, 0)
}

func (s *UrnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *UrnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UrnListener); ok {
		listenerT.EnterUrn(s)
	}
}

func (s *UrnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UrnListener); ok {
		listenerT.ExitUrn(s)
	}
}




func (p *UrnParser) Urn() (localctx IUrnContext) {
	localctx = NewUrnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UrnParserRULE_urn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.Match(UrnParserT__0)
	}
	{
		p.SetState(7)
		p.Match(UrnParserColon)
	}
	{
		p.SetState(8)
		p.ID()
	}
	{
		p.SetState(9)
		p.Match(UrnParserColon)
	}
	{
		p.SetState(10)
		p.SS()
	}
	{
		p.SetState(11)
		p.Match(UrnParserEOF)
	}



	return localctx
}


// IIDContext is an interface to support dynamic dispatch.
type IIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIDContext differentiates from other interfaces.
	IsIDContext()
}

type IDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIDContext() *IDContext {
	var p = new(IDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UrnParserRULE_iD
	return p
}

func (*IDContext) IsIDContext() {}

func NewIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IDContext {
	var p = new(IDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UrnParserRULE_iD

	return p
}

func (s *IDContext) GetParser() antlr.Parser { return s.parser }
func (s *IDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *IDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UrnListener); ok {
		listenerT.EnterID(s)
	}
}

func (s *IDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UrnListener); ok {
		listenerT.ExitID(s)
	}
}




func (p *UrnParser) ID() (localctx IIDContext) {
	localctx = NewIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UrnParserRULE_iD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Match(UrnParserT__1)
	}



	return localctx
}


// ISSContext is an interface to support dynamic dispatch.
type ISSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSSContext differentiates from other interfaces.
	IsSSContext()
}

type SSContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySSContext() *SSContext {
	var p = new(SSContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UrnParserRULE_sS
	return p
}

func (*SSContext) IsSSContext() {}

func NewSSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SSContext {
	var p = new(SSContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UrnParserRULE_sS

	return p
}

func (s *SSContext) GetParser() antlr.Parser { return s.parser }
func (s *SSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UrnListener); ok {
		listenerT.EnterSS(s)
	}
}

func (s *SSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UrnListener); ok {
		listenerT.ExitSS(s)
	}
}




func (p *UrnParser) SS() (localctx ISSContext) {
	localctx = NewSSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, UrnParserRULE_sS)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.Match(UrnParserT__2)
	}



	return localctx
}


