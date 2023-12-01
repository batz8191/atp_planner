// Code generated from ./parser2/Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Calc

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

import "os"
import "time"

var Result Interval

func Atoi(location, v string) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		fmt.Printf("Could not parse %s [%s]: %v\n", location, v, err)
		os.Exit(1)
	}
	return i
}
func ParseDuration(v string) time.Duration {
	d, err := time.ParseDuration(v)
	if err != nil {
		fmt.Printf("Could not parse duration [%s]: %v\n", v, err)
		os.Exit(1)
	}
	return d
}

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CalcParser struct {
	*antlr.BaseParser
}

var CalcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcParserInit() {
	staticData := &CalcParserStaticData
	staticData.LiteralNames = []string{
		"", "','", "'('", "')'", "'%'", "'-'", "'mile'", "'miles'", "'yard'",
		"'yards'", "'meter'", "'meters'", "'km'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "MUL", "ZONE", "DURSUFFIX",
		"NUMBER", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"start", "expression", "repeat", "atom", "zone", "range", "dur", "dist",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 17, 79, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 27, 8, 1, 10, 1, 12, 1, 30, 9, 1, 1, 1,
		1, 1, 3, 1, 34, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 3, 3, 47, 8, 3, 1, 3, 1, 3, 3, 3, 51, 8, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 63, 8, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 4, 6, 69, 8, 6, 11, 6, 12, 6, 70, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 1, 1, 0, 6, 12,
		77, 0, 16, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 4, 35, 1, 0, 0, 0, 6, 42, 1,
		0, 0, 0, 8, 54, 1, 0, 0, 0, 10, 57, 1, 0, 0, 0, 12, 68, 1, 0, 0, 0, 14,
		74, 1, 0, 0, 0, 16, 17, 3, 2, 1, 0, 17, 18, 5, 0, 0, 1, 18, 19, 6, 0, -1,
		0, 19, 1, 1, 0, 0, 0, 20, 21, 3, 4, 2, 0, 21, 22, 6, 1, -1, 0, 22, 34,
		1, 0, 0, 0, 23, 28, 3, 6, 3, 0, 24, 25, 5, 1, 0, 0, 25, 27, 3, 2, 1, 0,
		26, 24, 1, 0, 0, 0, 27, 30, 1, 0, 0, 0, 28, 26, 1, 0, 0, 0, 28, 29, 1,
		0, 0, 0, 29, 31, 1, 0, 0, 0, 30, 28, 1, 0, 0, 0, 31, 32, 6, 1, -1, 0, 32,
		34, 1, 0, 0, 0, 33, 20, 1, 0, 0, 0, 33, 23, 1, 0, 0, 0, 34, 3, 1, 0, 0,
		0, 35, 36, 5, 16, 0, 0, 36, 37, 5, 13, 0, 0, 37, 38, 5, 2, 0, 0, 38, 39,
		3, 2, 1, 0, 39, 40, 5, 3, 0, 0, 40, 41, 6, 2, -1, 0, 41, 5, 1, 0, 0, 0,
		42, 46, 6, 3, -1, 0, 43, 47, 3, 12, 6, 0, 44, 47, 3, 14, 7, 0, 45, 47,
		5, 16, 0, 0, 46, 43, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 45, 1, 0, 0, 0,
		47, 50, 1, 0, 0, 0, 48, 51, 3, 8, 4, 0, 49, 51, 3, 10, 5, 0, 50, 48, 1,
		0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 6, 3, -1, 0, 53,
		7, 1, 0, 0, 0, 54, 55, 5, 14, 0, 0, 55, 56, 6, 4, -1, 0, 56, 9, 1, 0, 0,
		0, 57, 58, 5, 16, 0, 0, 58, 62, 5, 4, 0, 0, 59, 60, 5, 5, 0, 0, 60, 61,
		5, 16, 0, 0, 61, 63, 5, 4, 0, 0, 62, 59, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0,
		63, 64, 1, 0, 0, 0, 64, 65, 6, 5, -1, 0, 65, 11, 1, 0, 0, 0, 66, 67, 5,
		16, 0, 0, 67, 69, 5, 15, 0, 0, 68, 66, 1, 0, 0, 0, 69, 70, 1, 0, 0, 0,
		70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 6,
		6, -1, 0, 73, 13, 1, 0, 0, 0, 74, 75, 5, 16, 0, 0, 75, 76, 7, 0, 0, 0,
		76, 77, 6, 7, -1, 0, 77, 15, 1, 0, 0, 0, 6, 28, 33, 46, 50, 62, 70,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalcParserInit initializes any static state used to implement CalcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcParserInit() {
	staticData := &CalcParserStaticData
	staticData.once.Do(calcParserInit)
}

// NewCalcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalcParser(input antlr.TokenStream) *CalcParser {
	CalcParserInit()
	this := new(CalcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Calc.g4"

	return this
}

// CalcParser tokens.
const (
	CalcParserEOF        = antlr.TokenEOF
	CalcParserT__0       = 1
	CalcParserT__1       = 2
	CalcParserT__2       = 3
	CalcParserT__3       = 4
	CalcParserT__4       = 5
	CalcParserT__5       = 6
	CalcParserT__6       = 7
	CalcParserT__7       = 8
	CalcParserT__8       = 9
	CalcParserT__9       = 10
	CalcParserT__10      = 11
	CalcParserT__11      = 12
	CalcParserMUL        = 13
	CalcParserZONE       = 14
	CalcParserDURSUFFIX  = 15
	CalcParserNUMBER     = 16
	CalcParserWHITESPACE = 17
)

// CalcParser rules.
const (
	CalcParserRULE_start      = 0
	CalcParserRULE_expression = 1
	CalcParserRULE_repeat     = 2
	CalcParserRULE_atom       = 3
	CalcParserRULE_zone       = 4
	CalcParserRULE_range      = 5
	CalcParserRULE_dur        = 6
	CalcParserRULE_dist       = 7
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() IExpressionContext

	// SetE sets the e rule contexts.
	SetE(IExpressionContext)

	// Getter signatures
	EOF() antlr.TerminalNode
	Expression() IExpressionContext

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	e      IExpressionContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) GetE() IExpressionContext { return s.e }

func (s *StartContext) SetE(v IExpressionContext) { s.e = v }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcParserEOF, 0)
}

func (s *StartContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *CalcParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcParserRULE_start)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)

		var _x = p.Expression()

		localctx.(*StartContext).e = _x
	}
	{
		p.SetState(17)
		p.Match(CalcParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	Result = localctx.GetE().GetRet()

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetR returns the r rule contexts.
	GetR() IRepeatContext

	// GetA returns the a rule contexts.
	GetA() IAtomContext

	// GetE returns the e rule contexts.
	GetE() IExpressionContext

	// SetR sets the r rule contexts.
	SetR(IRepeatContext)

	// SetA sets the a rule contexts.
	SetA(IAtomContext)

	// SetE sets the e rule contexts.
	SetE(IExpressionContext)

	// GetRet returns the ret attribute.
	GetRet() Interval

	// SetRet sets the ret attribute.
	SetRet(Interval)

	// Getter signatures
	Repeat() IRepeatContext
	Atom() IAtomContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ret    Interval
	r      IRepeatContext
	a      IAtomContext
	e      IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetR() IRepeatContext { return s.r }

func (s *ExpressionContext) GetA() IAtomContext { return s.a }

func (s *ExpressionContext) GetE() IExpressionContext { return s.e }

func (s *ExpressionContext) SetR(v IRepeatContext) { s.r = v }

func (s *ExpressionContext) SetA(v IAtomContext) { s.a = v }

func (s *ExpressionContext) SetE(v IExpressionContext) { s.e = v }

func (s *ExpressionContext) GetRet() Interval { return s.ret }

func (s *ExpressionContext) SetRet(v Interval) { s.ret = v }

func (s *ExpressionContext) Repeat() IRepeatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRepeatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRepeatContext)
}

func (s *ExpressionContext) Atom() IAtomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CalcParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalcParserRULE_expression)
	var _alt int

	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)

			var _x = p.Repeat()

			localctx.(*ExpressionContext).r = _x
		}

		localctx.(*ExpressionContext).ret = localctx.GetR().GetRet()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)

			var _x = p.Atom()

			localctx.(*ExpressionContext).a = _x
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(24)
					p.Match(CalcParserT__0)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(25)

					var _x = p.Expression()

					localctx.(*ExpressionContext).e = _x
				}

			}
			p.SetState(30)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

		all := localctx.AllExpression()
		if len(all) == 0 {
			localctx.(*ExpressionContext).ret = localctx.GetA().GetRet()
		} else {
			u := &UnionInterval{}
			u.interval = []Interval{localctx.GetA().GetRet()}
			for _, e := range localctx.AllExpression() {
				u.interval = append(u.interval, e.GetRet())
			}
			localctx.(*ExpressionContext).ret = u
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRepeatContext is an interface to support dynamic dispatch.
type IRepeatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IExpressionContext

	// SetE sets the e rule contexts.
	SetE(IExpressionContext)

	// GetRet returns the ret attribute.
	GetRet() *RepeatInterval

	// SetRet sets the ret attribute.
	SetRet(*RepeatInterval)

	// Getter signatures
	MUL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Expression() IExpressionContext

	// IsRepeatContext differentiates from other interfaces.
	IsRepeatContext()
}

type RepeatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ret    *RepeatInterval
	n      antlr.Token
	e      IExpressionContext
}

func NewEmptyRepeatContext() *RepeatContext {
	var p = new(RepeatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_repeat
	return p
}

func InitEmptyRepeatContext(p *RepeatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_repeat
}

func (*RepeatContext) IsRepeatContext() {}

func NewRepeatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepeatContext {
	var p = new(RepeatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_repeat

	return p
}

func (s *RepeatContext) GetParser() antlr.Parser { return s.parser }

func (s *RepeatContext) GetN() antlr.Token { return s.n }

func (s *RepeatContext) SetN(v antlr.Token) { s.n = v }

func (s *RepeatContext) GetE() IExpressionContext { return s.e }

func (s *RepeatContext) SetE(v IExpressionContext) { s.e = v }

func (s *RepeatContext) GetRet() *RepeatInterval { return s.ret }

func (s *RepeatContext) SetRet(v *RepeatInterval) { s.ret = v }

func (s *RepeatContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalcParserMUL, 0)
}

func (s *RepeatContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *RepeatContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RepeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepeatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRepeat(s)
	}
}

func (s *RepeatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRepeat(s)
	}
}

func (p *CalcParser) Repeat() (localctx IRepeatContext) {
	localctx = NewRepeatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_repeat)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)

		var _m = p.Match(CalcParserNUMBER)

		localctx.(*RepeatContext).n = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.Match(CalcParserMUL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.Match(CalcParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(38)

		var _x = p.Expression()

		localctx.(*RepeatContext).e = _x
	}
	{
		p.SetState(39)
		p.Match(CalcParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	localctx.(*RepeatContext).ret = &RepeatInterval{
		count:    Atoi("count", localctx.(*RepeatContext).GetN().GetText()),
		interval: localctx.GetE().GetRet(),
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// GetD returns the d rule contexts.
	GetD() IDurContext

	// GetI returns the i rule contexts.
	GetI() IDistContext

	// GetZ returns the z rule contexts.
	GetZ() IZoneContext

	// GetR returns the r rule contexts.
	GetR() IRangeContext

	// SetD sets the d rule contexts.
	SetD(IDurContext)

	// SetI sets the i rule contexts.
	SetI(IDistContext)

	// SetZ sets the z rule contexts.
	SetZ(IZoneContext)

	// SetR sets the r rule contexts.
	SetR(IRangeContext)

	// GetRet returns the ret attribute.
	GetRet() Interval

	// SetRet sets the ret attribute.
	SetRet(Interval)

	// Getter signatures
	Dur() IDurContext
	Dist() IDistContext
	NUMBER() antlr.TerminalNode
	Zone() IZoneContext
	Range_() IRangeContext

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ret    Interval
	d      IDurContext
	i      IDistContext
	n      antlr.Token
	z      IZoneContext
	r      IRangeContext
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_atom
	return p
}

func InitEmptyAtomContext(p *AtomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_atom
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) GetN() antlr.Token { return s.n }

func (s *AtomContext) SetN(v antlr.Token) { s.n = v }

func (s *AtomContext) GetD() IDurContext { return s.d }

func (s *AtomContext) GetI() IDistContext { return s.i }

func (s *AtomContext) GetZ() IZoneContext { return s.z }

func (s *AtomContext) GetR() IRangeContext { return s.r }

func (s *AtomContext) SetD(v IDurContext) { s.d = v }

func (s *AtomContext) SetI(v IDistContext) { s.i = v }

func (s *AtomContext) SetZ(v IZoneContext) { s.z = v }

func (s *AtomContext) SetR(v IRangeContext) { s.r = v }

func (s *AtomContext) GetRet() Interval { return s.ret }

func (s *AtomContext) SetRet(v Interval) { s.ret = v }

func (s *AtomContext) Dur() IDurContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDurContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDurContext)
}

func (s *AtomContext) Dist() IDistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistContext)
}

func (s *AtomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *AtomContext) Zone() IZoneContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZoneContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZoneContext)
}

func (s *AtomContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *CalcParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcParserRULE_atom)
	p.EnterOuterAlt(localctx, 1)
	var pow Power
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(43)

			var _x = p.Dur()

			localctx.(*AtomContext).d = _x
		}

	case 2:
		{
			p.SetState(44)

			var _x = p.Dist()

			localctx.(*AtomContext).i = _x
		}

	case 3:
		{
			p.SetState(45)

			var _m = p.Match(CalcParserNUMBER)

			localctx.(*AtomContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcParserZONE:
		{
			p.SetState(48)

			var _x = p.Zone()

			localctx.(*AtomContext).z = _x
		}

	case CalcParserNUMBER:
		{
			p.SetState(49)

			var _x = p.Range_()

			localctx.(*AtomContext).r = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

	if localctx.GetZ() != nil {
		pow = localctx.GetZ().GetRet()
	} else if localctx.GetR() != nil {
		pow = localctx.GetR().GetRet()
	}
	if localctx.GetD() != nil {
		localctx.(*AtomContext).ret = &DurationInterval{
			dur: localctx.GetD().GetRet(),
			pow: pow,
		}
	} else if localctx.GetI() != nil {
		localctx.(*AtomContext).ret = &DistanceInterval{
			dist: localctx.GetI().GetRet().Distance,
			unit: localctx.GetI().GetRet().Unit,
			pow:  pow,
		}
	} else if localctx.GetN() != nil {
		localctx.(*AtomContext).ret = &DistanceInterval{
			dist: Atoi("constant", localctx.(*AtomContext).GetN().GetText()),
			unit: "yards",
			pow:  pow,
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IZoneContext is an interface to support dynamic dispatch.
type IZoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetZ returns the z token.
	GetZ() antlr.Token

	// SetZ sets the z token.
	SetZ(antlr.Token)

	// GetRet returns the ret attribute.
	GetRet() *ZonePower

	// SetRet sets the ret attribute.
	SetRet(*ZonePower)

	// Getter signatures
	ZONE() antlr.TerminalNode

	// IsZoneContext differentiates from other interfaces.
	IsZoneContext()
}

type ZoneContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ret    *ZonePower
	z      antlr.Token
}

func NewEmptyZoneContext() *ZoneContext {
	var p = new(ZoneContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_zone
	return p
}

func InitEmptyZoneContext(p *ZoneContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_zone
}

func (*ZoneContext) IsZoneContext() {}

func NewZoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZoneContext {
	var p = new(ZoneContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_zone

	return p
}

func (s *ZoneContext) GetParser() antlr.Parser { return s.parser }

func (s *ZoneContext) GetZ() antlr.Token { return s.z }

func (s *ZoneContext) SetZ(v antlr.Token) { s.z = v }

func (s *ZoneContext) GetRet() *ZonePower { return s.ret }

func (s *ZoneContext) SetRet(v *ZonePower) { s.ret = v }

func (s *ZoneContext) ZONE() antlr.TerminalNode {
	return s.GetToken(CalcParserZONE, 0)
}

func (s *ZoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterZone(s)
	}
}

func (s *ZoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitZone(s)
	}
}

func (p *CalcParser) Zone() (localctx IZoneContext) {
	localctx = NewZoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcParserRULE_zone)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(54)

		var _m = p.Match(CalcParserZONE)

		localctx.(*ZoneContext).z = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	switch localctx.(*ZoneContext).GetZ().GetText() {
	case "Z1":
	case "Z2":
	case "Z3":
	case "Z4":
	case "Z5a":
	case "Z5b":
	case "Z5c":
	default:
		fmt.Printf("Unknown zone: %s\n", localctx.(*ZoneContext).GetZ())
		os.Exit(1)
	}
	localctx.(*ZoneContext).ret = &ZonePower{localctx.(*ZoneContext).GetZ().GetText()}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetL returns the l token.
	GetL() antlr.Token

	// GetH returns the h token.
	GetH() antlr.Token

	// SetL sets the l token.
	SetL(antlr.Token)

	// SetH sets the h token.
	SetH(antlr.Token)

	// GetRet returns the ret attribute.
	GetRet() *RangePower

	// SetRet sets the ret attribute.
	SetRet(*RangePower)

	// Getter signatures
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode

	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ret    *RangePower
	l      antlr.Token
	h      antlr.Token
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) GetL() antlr.Token { return s.l }

func (s *RangeContext) GetH() antlr.Token { return s.h }

func (s *RangeContext) SetL(v antlr.Token) { s.l = v }

func (s *RangeContext) SetH(v antlr.Token) { s.h = v }

func (s *RangeContext) GetRet() *RangePower { return s.ret }

func (s *RangeContext) SetRet(v *RangePower) { s.ret = v }

func (s *RangeContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(CalcParserNUMBER)
}

func (s *RangeContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, i)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitRange(s)
	}
}

func (p *CalcParser) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalcParserRULE_range)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)

		var _m = p.Match(CalcParserNUMBER)

		localctx.(*RangeContext).l = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(CalcParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CalcParserT__4 {
		{
			p.SetState(59)
			p.Match(CalcParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)

			var _m = p.Match(CalcParserNUMBER)

			localctx.(*RangeContext).h = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Match(CalcParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

	localctx.(*RangeContext).ret = &RangePower{
		low: Atoi("range low", localctx.(*RangeContext).GetL().GetText()),
	}
	if localctx.GetH() != nil {
		localctx.(*RangeContext).ret.high = Atoi("range high", localctx.(*RangeContext).GetH().GetText())
	} else {
		localctx.(*RangeContext).ret.high = localctx.(*RangeContext).ret.low
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDurContext is an interface to support dynamic dispatch.
type IDurContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRet returns the ret attribute.
	GetRet() time.Duration

	// SetRet sets the ret attribute.
	SetRet(time.Duration)

	// Getter signatures
	AllNUMBER() []antlr.TerminalNode
	NUMBER(i int) antlr.TerminalNode
	AllDURSUFFIX() []antlr.TerminalNode
	DURSUFFIX(i int) antlr.TerminalNode

	// IsDurContext differentiates from other interfaces.
	IsDurContext()
}

type DurContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ret    time.Duration
}

func NewEmptyDurContext() *DurContext {
	var p = new(DurContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_dur
	return p
}

func InitEmptyDurContext(p *DurContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_dur
}

func (*DurContext) IsDurContext() {}

func NewDurContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurContext {
	var p = new(DurContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_dur

	return p
}

func (s *DurContext) GetParser() antlr.Parser { return s.parser }

func (s *DurContext) GetRet() time.Duration { return s.ret }

func (s *DurContext) SetRet(v time.Duration) { s.ret = v }

func (s *DurContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(CalcParserNUMBER)
}

func (s *DurContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, i)
}

func (s *DurContext) AllDURSUFFIX() []antlr.TerminalNode {
	return s.GetTokens(CalcParserDURSUFFIX)
}

func (s *DurContext) DURSUFFIX(i int) antlr.TerminalNode {
	return s.GetToken(CalcParserDURSUFFIX, i)
}

func (s *DurContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDur(s)
	}
}

func (s *DurContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDur(s)
	}
}

func (p *CalcParser) Dur() (localctx IDurContext) {
	localctx = NewDurContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalcParserRULE_dur)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(66)
				p.Match(CalcParserNUMBER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(67)
				p.Match(CalcParserDURSUFFIX)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

	localctx.(*DurContext).ret = ParseDuration(localctx.GetText())

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDistContext is an interface to support dynamic dispatch.
type IDistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetN returns the n token.
	GetN() antlr.Token

	// GetU returns the u token.
	GetU() antlr.Token

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetU sets the u token.
	SetU(antlr.Token)

	// GetRet returns the ret attribute.
	GetRet() *Distance

	// SetRet sets the ret attribute.
	SetRet(*Distance)

	// Getter signatures
	NUMBER() antlr.TerminalNode

	// IsDistContext differentiates from other interfaces.
	IsDistContext()
}

type DistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ret    *Distance
	n      antlr.Token
	u      antlr.Token
}

func NewEmptyDistContext() *DistContext {
	var p = new(DistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_dist
	return p
}

func InitEmptyDistContext(p *DistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_dist
}

func (*DistContext) IsDistContext() {}

func NewDistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistContext {
	var p = new(DistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_dist

	return p
}

func (s *DistContext) GetParser() antlr.Parser { return s.parser }

func (s *DistContext) GetN() antlr.Token { return s.n }

func (s *DistContext) GetU() antlr.Token { return s.u }

func (s *DistContext) SetN(v antlr.Token) { s.n = v }

func (s *DistContext) SetU(v antlr.Token) { s.u = v }

func (s *DistContext) GetRet() *Distance { return s.ret }

func (s *DistContext) SetRet(v *Distance) { s.ret = v }

func (s *DistContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(CalcParserNUMBER, 0)
}

func (s *DistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDist(s)
	}
}

func (s *DistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDist(s)
	}
}

func (p *CalcParser) Dist() (localctx IDistContext) {
	localctx = NewDistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CalcParserRULE_dist)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)

		var _m = p.Match(CalcParserNUMBER)

		localctx.(*DistContext).n = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(75)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*DistContext).u = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8128) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*DistContext).u = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	localctx.(*DistContext).ret = &Distance{
		Distance: Atoi("distance", localctx.(*DistContext).GetN().GetText()),
		Unit:     localctx.(*DistContext).GetU().GetText(),
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
