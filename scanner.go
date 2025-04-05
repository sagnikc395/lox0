package main

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
	for !s.IsAtEnd() {
		s.start = s.current
	}
	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
	return s.tokens
}

// helper to check if we have consumed all the characters
func (s Scanner) IsAtEnd() bool {
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
