package main

import "fmt"

// a token class will represent the data about the
// tokentype , lexeme literal type and line number
// helpful for debugging down the line

type Token struct {
	Type   TokenType
	Lexeme string
	//using a empty interface as a replacement for object
	Literal interface{}
	Line    int
}

// constructor for the Token

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int) Token {
	return Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

//String returns the string repr of token

func (t Token) String() string {
	if t.Literal != nil {
		return fmt.Sprintf("%s %s %v", t.Type, t.Lexeme, t.Literal)
	}
	return fmt.Sprintf("%s %s", t.Type, t.Lexeme)
}
