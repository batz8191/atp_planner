// Code generated from ./parser2/Calc.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "','", "'('", "')'", "'%'", "'-'", "'mile'", "'miles'", "'yard'",
		"'yards'", "'meter'", "'meters'", "'km'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "MUL", "ZONE", "DURSUFFIX",
		"NUMBER", "WHITESPACE",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "MUL", "ZONE", "DURSUFFIX", "NUMBER", "WHITESPACE",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 104, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13,
		89, 8, 13, 1, 14, 1, 14, 1, 15, 4, 15, 94, 8, 15, 11, 15, 12, 15, 95, 1,
		16, 4, 16, 99, 8, 16, 11, 16, 12, 16, 100, 1, 16, 1, 16, 0, 0, 17, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23,
		12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 5, 1, 0, 49, 53, 1, 0,
		97, 99, 3, 0, 104, 104, 109, 109, 115, 115, 1, 0, 48, 57, 3, 0, 9, 10,
		13, 13, 32, 32, 106, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3,
		37, 1, 0, 0, 0, 5, 39, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 43, 1, 0, 0, 0,
		11, 45, 1, 0, 0, 0, 13, 50, 1, 0, 0, 0, 15, 56, 1, 0, 0, 0, 17, 61, 1,
		0, 0, 0, 19, 67, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 80, 1, 0, 0, 0, 25,
		83, 1, 0, 0, 0, 27, 85, 1, 0, 0, 0, 29, 90, 1, 0, 0, 0, 31, 93, 1, 0, 0,
		0, 33, 98, 1, 0, 0, 0, 35, 36, 5, 44, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38,
		5, 40, 0, 0, 38, 4, 1, 0, 0, 0, 39, 40, 5, 41, 0, 0, 40, 6, 1, 0, 0, 0,
		41, 42, 5, 37, 0, 0, 42, 8, 1, 0, 0, 0, 43, 44, 5, 45, 0, 0, 44, 10, 1,
		0, 0, 0, 45, 46, 5, 109, 0, 0, 46, 47, 5, 105, 0, 0, 47, 48, 5, 108, 0,
		0, 48, 49, 5, 101, 0, 0, 49, 12, 1, 0, 0, 0, 50, 51, 5, 109, 0, 0, 51,
		52, 5, 105, 0, 0, 52, 53, 5, 108, 0, 0, 53, 54, 5, 101, 0, 0, 54, 55, 5,
		115, 0, 0, 55, 14, 1, 0, 0, 0, 56, 57, 5, 121, 0, 0, 57, 58, 5, 97, 0,
		0, 58, 59, 5, 114, 0, 0, 59, 60, 5, 100, 0, 0, 60, 16, 1, 0, 0, 0, 61,
		62, 5, 121, 0, 0, 62, 63, 5, 97, 0, 0, 63, 64, 5, 114, 0, 0, 64, 65, 5,
		100, 0, 0, 65, 66, 5, 115, 0, 0, 66, 18, 1, 0, 0, 0, 67, 68, 5, 109, 0,
		0, 68, 69, 5, 101, 0, 0, 69, 70, 5, 116, 0, 0, 70, 71, 5, 101, 0, 0, 71,
		72, 5, 114, 0, 0, 72, 20, 1, 0, 0, 0, 73, 74, 5, 109, 0, 0, 74, 75, 5,
		101, 0, 0, 75, 76, 5, 116, 0, 0, 76, 77, 5, 101, 0, 0, 77, 78, 5, 114,
		0, 0, 78, 79, 5, 115, 0, 0, 79, 22, 1, 0, 0, 0, 80, 81, 5, 107, 0, 0, 81,
		82, 5, 109, 0, 0, 82, 24, 1, 0, 0, 0, 83, 84, 5, 42, 0, 0, 84, 26, 1, 0,
		0, 0, 85, 86, 5, 90, 0, 0, 86, 88, 7, 0, 0, 0, 87, 89, 7, 1, 0, 0, 88,
		87, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 28, 1, 0, 0, 0, 90, 91, 7, 2, 0,
		0, 91, 30, 1, 0, 0, 0, 92, 94, 7, 3, 0, 0, 93, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 32, 1, 0, 0, 0,
		97, 99, 7, 4, 0, 0, 98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1,
		0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 6, 16, 0,
		0, 103, 34, 1, 0, 0, 0, 4, 0, 88, 95, 100, 1, 6, 0, 0,
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

// CalcLexerInit initializes any static state used to implement CalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalcLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerT__0       = 1
	CalcLexerT__1       = 2
	CalcLexerT__2       = 3
	CalcLexerT__3       = 4
	CalcLexerT__4       = 5
	CalcLexerT__5       = 6
	CalcLexerT__6       = 7
	CalcLexerT__7       = 8
	CalcLexerT__8       = 9
	CalcLexerT__9       = 10
	CalcLexerT__10      = 11
	CalcLexerT__11      = 12
	CalcLexerMUL        = 13
	CalcLexerZONE       = 14
	CalcLexerDURSUFFIX  = 15
	CalcLexerNUMBER     = 16
	CalcLexerWHITESPACE = 17
)
