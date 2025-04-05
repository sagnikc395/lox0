package main

// keywords part of the shape of the language's grammar
// so the parser often has to code like ->
// if next token is while then do

//parser could categorize tokens from the raw lexeme by comparing the strings
// but slow and kind of ugly. Instead we recgonize a lexeme, we also remember which kind of lexeme
// it represents

// tokentype represents lexical token types
type TokenType = int

const (
	//single-character tokens

	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	//one or more character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	//Literals.
	IDENTIFIER
	STRING
	NUMBER

	//KEYWORDS.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	//OTHERS
	EOF
)
