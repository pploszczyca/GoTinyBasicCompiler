package domain

import "fmt"

type Token struct {
	Type  TokenType
	Value string
}

type TokenType int

const (
	Eof        TokenType = -1
	Cr         TokenType = 0
	Number     TokenType = 1
	Identifier TokenType = 2
	String     TokenType = 3

	// Keywords
	Print  TokenType = 101
	If     TokenType = 102
	Then   TokenType = 103
	Goto   TokenType = 104
	Input  TokenType = 105
	Let    TokenType = 106
	Gosub  TokenType = 107
	Return TokenType = 108
	Clear  TokenType = 109
	List   TokenType = 110
	Run    TokenType = 111
	End    TokenType = 112
	While  TokenType = 113
	Wend   TokenType = 114
	For    TokenType = 115
	To     TokenType = 116
	Next   TokenType = 117

	// Operators
	Plus               TokenType = 201
	Minus              TokenType = 202
	Multiply           TokenType = 203
	Divide             TokenType = 204
	Equal              TokenType = 205
	LessThan           TokenType = 206
	GreaterThan        TokenType = 207
	LessThanOrEqual    TokenType = 208
	GreaterThanOrEqual TokenType = 209
	NotEqual           TokenType = 210

	// Delimiters
	Comma     TokenType = 301
	Semicolon TokenType = 302
	LParen    TokenType = 303
	RParen    TokenType = 304
)

var tokenTypeToString = map[TokenType]string{
	Eof:                "Eof",
	Cr:                 "Cr",
	Number:             "Number",
	Identifier:         "Identifier",
	String:             "String",
	Print:              "Print",
	If:                 "If",
	Then:               "Then",
	Goto:               "Goto",
	Input:              "Input",
	Let:                "Let",
	Gosub:              "Gosub",
	Return:             "Return",
	Clear:              "Clear",
	List:               "List",
	Run:                "Run",
	End:                "End",
	While:              "While",
	Wend:               "Wend",
	For:                "For",
	To:                 "To",
	Next:               "Next",
	Plus:               "Plus",
	Minus:              "Minus",
	Multiply:           "Multiply",
	Divide:             "Divide",
	Equal:              "Equal",
	LessThan:           "LessThan",
	GreaterThan:        "GreaterThan",
	LessThanOrEqual:    "LessThanOrEqual",
	GreaterThanOrEqual: "GreaterThanOrEqual",
	NotEqual:           "NotEqual",
	Comma:              "Comma",
	Semicolon:          "Semicolon",
	LParen:             "LParen",
	RParen:             "RParen",
}

func (t TokenType) String() string {
	if str, ok := tokenTypeToString[t]; ok {
		return str
	}
	return fmt.Sprintf("Unknown TokenType (%d)", t)
}

type TokenIterator struct {
	currentIndex int
	tokens       []Token
}

func NewTokenIterator(tokens []Token) TokenIterator {
	return TokenIterator{
		currentIndex: 0,
		tokens:       tokens,
	}
}

func (ti *TokenIterator) Current() (Token, error) {
	if ti.currentIndex >= len(ti.tokens) {
		return Token{}, fmt.Errorf("tokens index out of range")
	}
	return ti.tokens[ti.currentIndex], nil
}

func (ti *TokenIterator) Next() {
	ti.currentIndex++
}

func (ti *TokenIterator) HasNext() bool {
	return ti.currentIndex < len(ti.tokens)
}
