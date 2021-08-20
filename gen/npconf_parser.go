// Code generated from /Users/umi/Desktop/dfoplanner/npconf/npconf.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // npconf

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 80, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4, 
	8, 9, 8, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3, 3, 3, 
	3, 7, 3, 26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 3, 3, 3, 7, 3, 33, 10, 
	3, 12, 3, 14, 3, 36, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 42, 10, 3, 12, 
	3, 14, 3, 45, 11, 3, 3, 3, 3, 3, 5, 3, 49, 10, 3, 3, 4, 3, 4, 3, 4, 3, 
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 
	7, 3, 7, 5, 7, 68, 10, 7, 3, 8, 3, 8, 3, 8, 5, 8, 73, 10, 8, 7, 8, 75, 
	10, 8, 12, 8, 14, 8, 78, 11, 8, 3, 8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 
	2, 3, 3, 2, 7, 8, 2, 82, 2, 20, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 50, 3, 
	2, 2, 2, 8, 55, 3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 67, 3, 2, 2, 2, 14, 
	69, 3, 2, 2, 2, 16, 19, 5, 4, 3, 2, 17, 19, 7, 10, 2, 2, 18, 16, 3, 2, 
	2, 2, 18, 17, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 
	3, 2, 2, 2, 21, 3, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 27, 5, 6, 4, 2, 
	24, 26, 5, 12, 7, 2, 25, 24, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 
	2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 49, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 
	34, 5, 6, 4, 2, 31, 33, 5, 4, 3, 2, 32, 31, 3, 2, 2, 2, 33, 36, 3, 2, 2, 
	2, 34, 32, 3, 2, 2, 2, 34, 35, 3, 2, 2, 2, 35, 37, 3, 2, 2, 2, 36, 34, 
	3, 2, 2, 2, 37, 38, 5, 8, 5, 2, 38, 49, 3, 2, 2, 2, 39, 43, 5, 6, 4, 2, 
	40, 42, 5, 12, 7, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 
	2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 
	47, 5, 8, 5, 2, 47, 49, 3, 2, 2, 2, 48, 23, 3, 2, 2, 2, 48, 30, 3, 2, 2, 
	2, 48, 39, 3, 2, 2, 2, 49, 5, 3, 2, 2, 2, 50, 51, 7, 3, 2, 2, 51, 52, 5, 
	10, 6, 2, 52, 53, 7, 4, 2, 2, 53, 54, 7, 10, 2, 2, 54, 7, 3, 2, 2, 2, 55, 
	56, 7, 3, 2, 2, 56, 57, 7, 5, 2, 2, 57, 58, 5, 10, 6, 2, 58, 59, 7, 4, 
	2, 2, 59, 60, 7, 10, 2, 2, 60, 9, 3, 2, 2, 2, 61, 62, 9, 2, 2, 2, 62, 11, 
	3, 2, 2, 2, 63, 64, 5, 14, 8, 2, 64, 65, 7, 10, 2, 2, 65, 68, 3, 2, 2, 
	2, 66, 68, 5, 14, 8, 2, 67, 63, 3, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 13, 
	3, 2, 2, 2, 69, 76, 5, 10, 6, 2, 70, 72, 7, 6, 2, 2, 71, 73, 5, 10, 6, 
	2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 75, 3, 2, 2, 2, 74, 70, 
	3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 
	77, 15, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 11, 18, 20, 27, 34, 43, 48, 67, 
	72, 76,
}
var literalNames = []string{
	"", "'['", "']'", "'/'", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "CHARS", "STRING", "COMMENT", "EOL", "WS",
}

var ruleNames = []string{
	"npconf", "section", "sectionheader", "sectionfooter", "str", "line", "stringlist",
}
type npconfParser struct {
	*antlr.BaseParser
}

// NewnpconfParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *npconfParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewnpconfParser(input antlr.TokenStream) *npconfParser {
	this := new(npconfParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "npconf.g4"

	return this
}


// npconfParser tokens.
const (
	npconfParserEOF = antlr.TokenEOF
	npconfParserT__0 = 1
	npconfParserT__1 = 2
	npconfParserT__2 = 3
	npconfParserT__3 = 4
	npconfParserCHARS = 5
	npconfParserSTRING = 6
	npconfParserCOMMENT = 7
	npconfParserEOL = 8
	npconfParserWS = 9
)

// npconfParser rules.
const (
	npconfParserRULE_npconf = 0
	npconfParserRULE_section = 1
	npconfParserRULE_sectionheader = 2
	npconfParserRULE_sectionfooter = 3
	npconfParserRULE_str = 4
	npconfParserRULE_line = 5
	npconfParserRULE_stringlist = 6
)

// INpconfContext is an interface to support dynamic dispatch.
type INpconfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNpconfContext differentiates from other interfaces.
	IsNpconfContext()
}

type NpconfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNpconfContext() *NpconfContext {
	var p = new(NpconfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = npconfParserRULE_npconf
	return p
}

func (*NpconfContext) IsNpconfContext() {}

func NewNpconfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NpconfContext {
	var p = new(NpconfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = npconfParserRULE_npconf

	return p
}

func (s *NpconfContext) GetParser() antlr.Parser { return s.parser }

func (s *NpconfContext) AllSection() []ISectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISectionContext)(nil)).Elem())
	var tst = make([]ISectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISectionContext)
		}
	}

	return tst
}

func (s *NpconfContext) Section(i int) ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *NpconfContext) AllEOL() []antlr.TerminalNode {
	return s.GetTokens(npconfParserEOL)
}

func (s *NpconfContext) EOL(i int) antlr.TerminalNode {
	return s.GetToken(npconfParserEOL, i)
}

func (s *NpconfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NpconfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NpconfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.EnterNpconf(s)
	}
}

func (s *NpconfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.ExitNpconf(s)
	}
}




func (p *npconfParser) Npconf() (localctx INpconfContext) {
	localctx = NewNpconfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, npconfParserRULE_npconf)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == npconfParserT__0 || _la == npconfParserEOL {
		p.SetState(16)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case npconfParserT__0:
			{
				p.SetState(14)
				p.Section()
			}


		case npconfParserEOL:
			{
				p.SetState(15)
				p.Match(npconfParserEOL)
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = npconfParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = npconfParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) Sectionheader() ISectionheaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionheaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionheaderContext)
}

func (s *SectionContext) AllLine() []ILineContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineContext)(nil)).Elem())
	var tst = make([]ILineContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineContext)
		}
	}

	return tst
}

func (s *SectionContext) Line(i int) ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *SectionContext) Sectionfooter() ISectionfooterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionfooterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionfooterContext)
}

func (s *SectionContext) AllSection() []ISectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISectionContext)(nil)).Elem())
	var tst = make([]ISectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISectionContext)
		}
	}

	return tst
}

func (s *SectionContext) Section(i int) ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.ExitSection(s)
	}
}




func (p *npconfParser) Section() (localctx ISectionContext) {
	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, npconfParserRULE_section)
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

	var _alt int

	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Sectionheader()
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == npconfParserCHARS || _la == npconfParserSTRING {
			{
				p.SetState(22)
				p.Line()
			}


			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Sectionheader()
		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(29)
					p.Section()
				}


			}
			p.SetState(34)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		{
			p.SetState(35)
			p.Sectionfooter()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(37)
			p.Sectionheader()
		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == npconfParserCHARS || _la == npconfParserSTRING {
			{
				p.SetState(38)
				p.Line()
			}


			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(44)
			p.Sectionfooter()
		}

	}


	return localctx
}


// ISectionheaderContext is an interface to support dynamic dispatch.
type ISectionheaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionheaderContext differentiates from other interfaces.
	IsSectionheaderContext()
}

type SectionheaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionheaderContext() *SectionheaderContext {
	var p = new(SectionheaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = npconfParserRULE_sectionheader
	return p
}

func (*SectionheaderContext) IsSectionheaderContext() {}

func NewSectionheaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionheaderContext {
	var p = new(SectionheaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = npconfParserRULE_sectionheader

	return p
}

func (s *SectionheaderContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionheaderContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *SectionheaderContext) EOL() antlr.TerminalNode {
	return s.GetToken(npconfParserEOL, 0)
}

func (s *SectionheaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionheaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SectionheaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.EnterSectionheader(s)
	}
}

func (s *SectionheaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.ExitSectionheader(s)
	}
}




func (p *npconfParser) Sectionheader() (localctx ISectionheaderContext) {
	localctx = NewSectionheaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, npconfParserRULE_sectionheader)

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
		p.SetState(48)
		p.Match(npconfParserT__0)
	}
	{
		p.SetState(49)
		p.Str()
	}
	{
		p.SetState(50)
		p.Match(npconfParserT__1)
	}
	{
		p.SetState(51)
		p.Match(npconfParserEOL)
	}



	return localctx
}


// ISectionfooterContext is an interface to support dynamic dispatch.
type ISectionfooterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionfooterContext differentiates from other interfaces.
	IsSectionfooterContext()
}

type SectionfooterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionfooterContext() *SectionfooterContext {
	var p = new(SectionfooterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = npconfParserRULE_sectionfooter
	return p
}

func (*SectionfooterContext) IsSectionfooterContext() {}

func NewSectionfooterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionfooterContext {
	var p = new(SectionfooterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = npconfParserRULE_sectionfooter

	return p
}

func (s *SectionfooterContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionfooterContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *SectionfooterContext) EOL() antlr.TerminalNode {
	return s.GetToken(npconfParserEOL, 0)
}

func (s *SectionfooterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionfooterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SectionfooterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.EnterSectionfooter(s)
	}
}

func (s *SectionfooterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.ExitSectionfooter(s)
	}
}




func (p *npconfParser) Sectionfooter() (localctx ISectionfooterContext) {
	localctx = NewSectionfooterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, npconfParserRULE_sectionfooter)

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
		p.SetState(53)
		p.Match(npconfParserT__0)
	}
	{
		p.SetState(54)
		p.Match(npconfParserT__2)
	}
	{
		p.SetState(55)
		p.Str()
	}
	{
		p.SetState(56)
		p.Match(npconfParserT__1)
	}
	{
		p.SetState(57)
		p.Match(npconfParserEOL)
	}



	return localctx
}


// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = npconfParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = npconfParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) CHARS() antlr.TerminalNode {
	return s.GetToken(npconfParserCHARS, 0)
}

func (s *StrContext) STRING() antlr.TerminalNode {
	return s.GetToken(npconfParserSTRING, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.ExitStr(s)
	}
}




func (p *npconfParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, npconfParserRULE_str)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		_la = p.GetTokenStream().LA(1)

		if !(_la == npconfParserCHARS || _la == npconfParserSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = npconfParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = npconfParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Stringlist() IStringlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringlistContext)
}

func (s *LineContext) EOL() antlr.TerminalNode {
	return s.GetToken(npconfParserEOL, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.ExitLine(s)
	}
}




func (p *npconfParser) Line() (localctx ILineContext) {
	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, npconfParserRULE_line)

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

	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Stringlist()
		}
		{
			p.SetState(62)
			p.Match(npconfParserEOL)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(64)
			p.Stringlist()
		}

	}


	return localctx
}


// IStringlistContext is an interface to support dynamic dispatch.
type IStringlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringlistContext differentiates from other interfaces.
	IsStringlistContext()
}

type StringlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringlistContext() *StringlistContext {
	var p = new(StringlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = npconfParserRULE_stringlist
	return p
}

func (*StringlistContext) IsStringlistContext() {}

func NewStringlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringlistContext {
	var p = new(StringlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = npconfParserRULE_stringlist

	return p
}

func (s *StringlistContext) GetParser() antlr.Parser { return s.parser }

func (s *StringlistContext) AllStr() []IStrContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStrContext)(nil)).Elem())
	var tst = make([]IStrContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStrContext)
		}
	}

	return tst
}

func (s *StringlistContext) Str(i int) IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
}

func (s *StringlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StringlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.EnterStringlist(s)
	}
}

func (s *StringlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(npconfListener); ok {
		listenerT.ExitStringlist(s)
	}
}




func (p *npconfParser) Stringlist() (localctx IStringlistContext) {
	localctx = NewStringlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, npconfParserRULE_stringlist)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Str()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == npconfParserT__3 {
		{
			p.SetState(68)
			p.Match(npconfParserT__3)
		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)


		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(69)
				p.Str()
			}


		}


		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


