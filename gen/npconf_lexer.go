// Code generated from /Users/umi/Desktop/dfoplanner/npconf/npconf.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 64, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 
	3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6, 6, 32, 10, 6, 13, 6, 14, 6, 33, 3, 7, 
	3, 7, 7, 7, 38, 10, 7, 12, 7, 14, 7, 41, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 
	7, 8, 47, 10, 8, 12, 8, 14, 8, 50, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 
	3, 9, 3, 10, 6, 10, 59, 10, 10, 13, 10, 14, 10, 60, 3, 10, 3, 10, 2, 2, 
	11, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 
	7, 15, 2, 11, 11, 34, 34, 36, 36, 39, 40, 44, 44, 47, 59, 62, 62, 64, 64, 
	66, 92, 94, 94, 97, 97, 99, 125, 127, 127, 4, 2, 12, 12, 36, 36, 4, 2, 
	37, 37, 61, 61, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 2, 67, 2, 3, 
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 
	19, 3, 2, 2, 2, 3, 21, 3, 2, 2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 
	9, 28, 3, 2, 2, 2, 11, 31, 3, 2, 2, 2, 13, 35, 3, 2, 2, 2, 15, 44, 3, 2, 
	2, 2, 17, 55, 3, 2, 2, 2, 19, 58, 3, 2, 2, 2, 21, 22, 7, 93, 2, 2, 22, 
	4, 3, 2, 2, 2, 23, 24, 7, 95, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 93, 2, 
	2, 26, 27, 7, 49, 2, 2, 27, 8, 3, 2, 2, 2, 28, 29, 7, 46, 2, 2, 29, 10, 
	3, 2, 2, 2, 30, 32, 9, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 
	33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 12, 3, 2, 2, 2, 35, 39, 7, 
	36, 2, 2, 36, 38, 10, 3, 2, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 
	39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 
	2, 2, 2, 42, 43, 7, 36, 2, 2, 43, 14, 3, 2, 2, 2, 44, 48, 9, 4, 2, 2, 45, 
	47, 10, 5, 2, 2, 46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2, 2, 48, 46, 3, 2, 
	2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 51, 52, 
	5, 17, 9, 2, 52, 53, 3, 2, 2, 2, 53, 54, 8, 8, 2, 2, 54, 16, 3, 2, 2, 2, 
	55, 56, 9, 5, 2, 2, 56, 18, 3, 2, 2, 2, 57, 59, 9, 6, 2, 2, 58, 57, 3, 
	2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 
	62, 3, 2, 2, 2, 62, 63, 8, 10, 2, 2, 63, 20, 3, 2, 2, 2, 7, 2, 33, 39, 
	48, 60, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'[/'", "','",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "CHARS", "STRING", "COMMENT", "EOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "CHARS", "STRING", "COMMENT", "EOL", "WS",
}
type npconfLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

// NewnpconfLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *npconfLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewnpconfLexer(input antlr.CharStream) *npconfLexer {
	l := new(npconfLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "npconf.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// npconfLexer tokens.
const (
	npconfLexerT__0 = 1
	npconfLexerT__1 = 2
	npconfLexerT__2 = 3
	npconfLexerT__3 = 4
	npconfLexerCHARS = 5
	npconfLexerSTRING = 6
	npconfLexerCOMMENT = 7
	npconfLexerEOL = 8
	npconfLexerWS = 9
)

