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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 5, 30, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 
	2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 22, 10, 4, 12, 4, 14, 4, 25, 
	11, 4, 3, 4, 5, 4, 28, 10, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 28, 2, 8, 
	3, 2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 9, 7, 3, 2, 2, 9, 
	10, 7, 5, 2, 2, 10, 11, 5, 4, 3, 2, 11, 12, 7, 5, 2, 2, 12, 13, 5, 6, 4, 
	2, 13, 14, 7, 2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 16, 7, 4, 2, 2, 16, 17, 6, 
	3, 2, 3, 17, 5, 3, 2, 2, 2, 18, 23, 7, 4, 2, 2, 19, 20, 7, 5, 2, 2, 20, 
	22, 7, 4, 2, 2, 21, 19, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3, 2, 2, 
	2, 23, 24, 3, 2, 2, 2, 24, 28, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26, 28, 
	7, 3, 2, 2, 27, 18, 3, 2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 7, 3, 2, 2, 2, 
	4, 23, 27,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "':'",
}
var symbolicNames = []string{
	"", "Urn", "Part", "Colon",
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
	UrnParserUrn = 1
	UrnParserPart = 2
	UrnParserColon = 3
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

func (s *UrnContext) Urn() antlr.TerminalNode {
	return s.GetToken(UrnParserUrn, 0)
}

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
		p.Match(UrnParserUrn)
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

func (s *IDContext) Part() antlr.TerminalNode {
	return s.GetToken(UrnParserPart, 0)
}

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
		p.Match(UrnParserPart)
	}
	p.SetState(14)

	if !(len(localctx.(*IDContext).GetText()) <= 32) {
		panic(antlr.NewFailedPredicateException(p, "len($ctx.GetText()) <= 32", ""))
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

func (s *SSContext) AllPart() []antlr.TerminalNode {
	return s.GetTokens(UrnParserPart)
}

func (s *SSContext) Part(i int) antlr.TerminalNode {
	return s.GetToken(UrnParserPart, i)
}

func (s *SSContext) AllColon() []antlr.TerminalNode {
	return s.GetTokens(UrnParserColon)
}

func (s *SSContext) Colon(i int) antlr.TerminalNode {
	return s.GetToken(UrnParserColon, i)
}

func (s *SSContext) Urn() antlr.TerminalNode {
	return s.GetToken(UrnParserUrn, 0)
}

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
	var _la int


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

	p.SetState(25)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case UrnParserPart:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(16)
			p.Match(UrnParserPart)
		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == UrnParserColon {
			{
				p.SetState(17)
				p.Match(UrnParserColon)
			}
			{
				p.SetState(18)
				p.Match(UrnParserPart)
			}


			p.SetState(23)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case UrnParserUrn:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(UrnParserUrn)
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


func (p *UrnParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
			var t *IDContext = nil
			if localctx != nil { t = localctx.(*IDContext) }
			return p.ID_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *UrnParser) ID_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return len(localctx.(*IDContext).GetText()) <= 32

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

