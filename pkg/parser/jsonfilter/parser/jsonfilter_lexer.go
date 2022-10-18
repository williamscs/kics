// Code generated from JSONFilter.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 209,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 7, 25, 142,
	10, 25, 12, 25, 14, 25, 145, 11, 25, 3, 26, 3, 26, 3, 26, 7, 26, 150, 10,
	26, 12, 26, 14, 26, 153, 11, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 5,
	27, 160, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 31, 5, 31, 173, 10, 31, 3, 31, 3, 31, 3, 31, 6, 31, 178,
	10, 31, 13, 31, 14, 31, 179, 5, 31, 182, 10, 31, 3, 31, 5, 31, 185, 10,
	31, 3, 32, 3, 32, 3, 32, 7, 32, 190, 10, 32, 12, 32, 14, 32, 193, 11, 32,
	5, 32, 195, 10, 32, 3, 33, 3, 33, 5, 33, 199, 10, 33, 3, 33, 3, 33, 3,
	34, 6, 34, 204, 10, 34, 13, 34, 14, 34, 205, 3, 34, 3, 34, 2, 2, 35, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 2, 55, 2, 57, 2, 59, 2, 61,
	28, 63, 2, 65, 2, 67, 29, 3, 2, 12, 4, 2, 67, 92, 99, 124, 5, 2, 50, 59,
	67, 92, 99, 124, 10, 2, 36, 36, 49, 49, 94, 94, 100, 100, 104, 104, 112,
	112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 2, 33, 36,
	36, 94, 94, 4, 2, 45, 45, 47, 47, 3, 2, 50, 59, 3, 2, 51, 59, 4, 2, 71,
	71, 103, 103, 5, 2, 11, 12, 15, 15, 34, 34, 2, 214, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2,
	2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2,
	2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2,
	2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3,
	2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 3, 69, 3, 2, 2, 2, 5,
	72, 3, 2, 2, 2, 7, 74, 3, 2, 2, 2, 9, 76, 3, 2, 2, 2, 11, 78, 3, 2, 2,
	2, 13, 80, 3, 2, 2, 2, 15, 82, 3, 2, 2, 2, 17, 84, 3, 2, 2, 2, 19, 86,
	3, 2, 2, 2, 21, 88, 3, 2, 2, 2, 23, 91, 3, 2, 2, 2, 25, 94, 3, 2, 2, 2,
	27, 96, 3, 2, 2, 2, 29, 99, 3, 2, 2, 2, 31, 101, 3, 2, 2, 2, 33, 103, 3,
	2, 2, 2, 35, 106, 3, 2, 2, 2, 37, 109, 3, 2, 2, 2, 39, 112, 3, 2, 2, 2,
	41, 116, 3, 2, 2, 2, 43, 121, 3, 2, 2, 2, 45, 128, 3, 2, 2, 2, 47, 133,
	3, 2, 2, 2, 49, 139, 3, 2, 2, 2, 51, 146, 3, 2, 2, 2, 53, 156, 3, 2, 2,
	2, 55, 161, 3, 2, 2, 2, 57, 167, 3, 2, 2, 2, 59, 169, 3, 2, 2, 2, 61, 172,
	3, 2, 2, 2, 63, 194, 3, 2, 2, 2, 65, 196, 3, 2, 2, 2, 67, 203, 3, 2, 2,
	2, 69, 70, 7, 38, 2, 2, 70, 71, 7, 48, 2, 2, 71, 4, 3, 2, 2, 2, 72, 73,
	7, 44, 2, 2, 73, 6, 3, 2, 2, 2, 74, 75, 7, 125, 2, 2, 75, 8, 3, 2, 2, 2,
	76, 77, 7, 127, 2, 2, 77, 10, 3, 2, 2, 2, 78, 79, 7, 42, 2, 2, 79, 12,
	3, 2, 2, 2, 80, 81, 7, 43, 2, 2, 81, 14, 3, 2, 2, 2, 82, 83, 7, 93, 2,
	2, 83, 16, 3, 2, 2, 2, 84, 85, 7, 95, 2, 2, 85, 18, 3, 2, 2, 2, 86, 87,
	7, 48, 2, 2, 87, 20, 3, 2, 2, 2, 88, 89, 7, 40, 2, 2, 89, 90, 7, 40, 2,
	2, 90, 22, 3, 2, 2, 2, 91, 92, 7, 126, 2, 2, 92, 93, 7, 126, 2, 2, 93,
	24, 3, 2, 2, 2, 94, 95, 7, 63, 2, 2, 95, 26, 3, 2, 2, 2, 96, 97, 7, 35,
	2, 2, 97, 98, 7, 63, 2, 2, 98, 28, 3, 2, 2, 2, 99, 100, 7, 64, 2, 2, 100,
	30, 3, 2, 2, 2, 101, 102, 7, 62, 2, 2, 102, 32, 3, 2, 2, 2, 103, 104, 7,
	64, 2, 2, 104, 105, 7, 63, 2, 2, 105, 34, 3, 2, 2, 2, 106, 107, 7, 62,
	2, 2, 107, 108, 7, 63, 2, 2, 108, 36, 3, 2, 2, 2, 109, 110, 7, 75, 2, 2,
	110, 111, 7, 85, 2, 2, 111, 38, 3, 2, 2, 2, 112, 113, 7, 80, 2, 2, 113,
	114, 7, 81, 2, 2, 114, 115, 7, 86, 2, 2, 115, 40, 3, 2, 2, 2, 116, 117,
	7, 80, 2, 2, 117, 118, 7, 87, 2, 2, 118, 119, 7, 78, 2, 2, 119, 120, 7,
	78, 2, 2, 120, 42, 3, 2, 2, 2, 121, 122, 7, 71, 2, 2, 122, 123, 7, 90,
	2, 2, 123, 124, 7, 75, 2, 2, 124, 125, 7, 85, 2, 2, 125, 126, 7, 86, 2,
	2, 126, 127, 7, 85, 2, 2, 127, 44, 3, 2, 2, 2, 128, 129, 7, 86, 2, 2, 129,
	130, 7, 84, 2, 2, 130, 131, 7, 87, 2, 2, 131, 132, 7, 71, 2, 2, 132, 46,
	3, 2, 2, 2, 133, 134, 7, 72, 2, 2, 134, 135, 7, 67, 2, 2, 135, 136, 7,
	78, 2, 2, 136, 137, 7, 85, 2, 2, 137, 138, 7, 71, 2, 2, 138, 48, 3, 2,
	2, 2, 139, 143, 9, 2, 2, 2, 140, 142, 9, 3, 2, 2, 141, 140, 3, 2, 2, 2,
	142, 145, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144,
	50, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 151, 7, 36, 2, 2, 147, 150,
	5, 53, 27, 2, 148, 150, 5, 59, 30, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3,
	2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2,
	2, 152, 154, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 155, 7, 36, 2, 2, 155,
	52, 3, 2, 2, 2, 156, 159, 7, 94, 2, 2, 157, 160, 9, 4, 2, 2, 158, 160,
	5, 55, 28, 2, 159, 157, 3, 2, 2, 2, 159, 158, 3, 2, 2, 2, 160, 54, 3, 2,
	2, 2, 161, 162, 7, 119, 2, 2, 162, 163, 5, 57, 29, 2, 163, 164, 5, 57,
	29, 2, 164, 165, 5, 57, 29, 2, 165, 166, 5, 57, 29, 2, 166, 56, 3, 2, 2,
	2, 167, 168, 9, 5, 2, 2, 168, 58, 3, 2, 2, 2, 169, 170, 10, 6, 2, 2, 170,
	60, 3, 2, 2, 2, 171, 173, 9, 7, 2, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3,
	2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 181, 5, 63, 32, 2, 175, 177, 7, 48,
	2, 2, 176, 178, 9, 8, 2, 2, 177, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2,
	179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181,
	175, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 184, 3, 2, 2, 2, 183, 185,
	5, 65, 33, 2, 184, 183, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 62, 3, 2,
	2, 2, 186, 195, 7, 50, 2, 2, 187, 191, 9, 9, 2, 2, 188, 190, 9, 8, 2, 2,
	189, 188, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191,
	192, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 186,
	3, 2, 2, 2, 194, 187, 3, 2, 2, 2, 195, 64, 3, 2, 2, 2, 196, 198, 9, 10,
	2, 2, 197, 199, 9, 7, 2, 2, 198, 197, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2,
	199, 200, 3, 2, 2, 2, 200, 201, 5, 63, 32, 2, 201, 66, 3, 2, 2, 2, 202,
	204, 9, 11, 2, 2, 203, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 203,
	3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 8, 34,
	2, 2, 208, 68, 3, 2, 2, 2, 15, 2, 143, 149, 151, 159, 172, 179, 181, 184,
	191, 194, 198, 205, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'$.'", "'*'", "'{'", "'}'", "'('", "')'", "'['", "']'", "'.'", "'&&'",
	"'||'", "'='", "'!='", "'>'", "'<'", "'>='", "'<='", "'IS'", "'NOT'", "'NULL'",
	"'EXISTS'", "'TRUE'", "'FALSE'",
}

var lexerSymbolicNames = []string{
	"", "SEL_START", "STAR", "LCURLY", "RCURLY", "LPAREN", "RPAREN", "LBRACKET",
	"RBRACKET", "DOT", "AND", "OR", "EQUALS", "NOT_EQUALS", "GT", "LT", "GE",
	"LE", "IS", "NOT", "NULL", "EXISTS", "TRUE", "FALSE", "INDENTIFIER", "STRING",
	"NUMBER", "WS",
}

var lexerRuleNames = []string{
	"SEL_START", "STAR", "LCURLY", "RCURLY", "LPAREN", "RPAREN", "LBRACKET",
	"RBRACKET", "DOT", "AND", "OR", "EQUALS", "NOT_EQUALS", "GT", "LT", "GE",
	"LE", "IS", "NOT", "NULL", "EXISTS", "TRUE", "FALSE", "INDENTIFIER", "STRING",
	"ESC", "UNICODE", "HEX", "SAFECODEPOINT", "NUMBER", "INT", "EXP", "WS",
}

type JSONFilterLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewJSONFilterLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *JSONFilterLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewJSONFilterLexer(input antlr.CharStream) *JSONFilterLexer {
	l := new(JSONFilterLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.Deserialize(uint16Toint32(serializedLexerAtn))
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
	l.GrammarFileName = "JSONFilter.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JSONFilterLexer tokens.
const (
	JSONFilterLexerSEL_START   = 1
	JSONFilterLexerSTAR        = 2
	JSONFilterLexerLCURLY      = 3
	JSONFilterLexerRCURLY      = 4
	JSONFilterLexerLPAREN      = 5
	JSONFilterLexerRPAREN      = 6
	JSONFilterLexerLBRACKET    = 7
	JSONFilterLexerRBRACKET    = 8
	JSONFilterLexerDOT         = 9
	JSONFilterLexerAND         = 10
	JSONFilterLexerOR          = 11
	JSONFilterLexerEQUALS      = 12
	JSONFilterLexerNOT_EQUALS  = 13
	JSONFilterLexerGT          = 14
	JSONFilterLexerLT          = 15
	JSONFilterLexerGE          = 16
	JSONFilterLexerLE          = 17
	JSONFilterLexerIS          = 18
	JSONFilterLexerNOT         = 19
	JSONFilterLexerNULL        = 20
	JSONFilterLexerEXISTS      = 21
	JSONFilterLexerTRUE        = 22
	JSONFilterLexerFALSE       = 23
	JSONFilterLexerINDENTIFIER = 24
	JSONFilterLexerSTRING      = 25
	JSONFilterLexerNUMBER      = 26
	JSONFilterLexerWS          = 27
)

// uint16Toint32 converts []uint16 to []int32
func uint16Toint32(atn []uint16) []int32 {
	atn32 := make([]int32, 0)
	for i := range atn {
		atn32 = append(atn32, int32(atn[i]))
	}
	return atn32
}
