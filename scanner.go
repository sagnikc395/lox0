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
