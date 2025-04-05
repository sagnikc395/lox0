package main

import "strconv"

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string, tokens []string) Scanner {
	return Scanner{
		source:  source,
		tokens:  make([]Token, BUFFER_SIZE),
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
	}
	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
	return s.tokens
}

// helper to check if we have consumed all the characters
func (s Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

//in each turn of the loop, we scan a single token.
// imagine if every lexeme were only a single character long.

func (s Scanner) ScanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(PLUS, nil)
	case '*':
		s.addToken(STAR, nil)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL, nil)
		} else {
			s.addToken(BANG, nil)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL, nil)
		} else {
			s.addToken(EQUAL, nil)
		}

	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL, nil)
		} else {
			s.addToken(LESS, nil)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL, nil)
		} else {
			s.addToken(GREATER, nil)
		}

	case '/':
		//for longer lexemes ,division and comments
		if s.match('/') {
			//comment goes until the end of the line
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}

	//scanning meaningless characters
	case ' ':
	case '\r':
	case '\t':
		//ignore whitespaces
		break
	case '\n':
		s.line++

	//string literals
	case '"':
		s.string()

	default:
		if s.isDigit(c) {
			s.number()
		} else {
			LoxError(NewLox(true), s.line, "Unexpected character.")
		}
	}
}

func (s Scanner) advance() byte {
	temp := s.current + 1
	return s.source[temp]
}

func (s *Scanner) addToken(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}

	s.current += 1
	return true
}

// peek one token ahead helper function
func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		//return null byte
		return '\x00'
	}
	return s.source[s.current]
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		LoxError(NewLox(false), s.line, "Unterminated string.")
		return
	}

	//close ".
	s.advance()

	//trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, value)
}

func (s Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	//look out for a fractional part
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		//consume "."
		s.advance()
	}
	for s.isDigit(s.peek()) {
		s.advance()
	}

	repr, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		LoxError(NewLox(false), s.line, "Error converting to float type.")
		return
	}
	s.addToken(NUMBER, repr)
}

func (s Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		//return null string
		return '\x00'
	}
	return s.source[s.current+1]
}
