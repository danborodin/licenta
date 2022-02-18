package main

type Scanner struct {
	Source  string
	Tokens  []Token
	start   int
	current int
	line    int
}

func scanTokens(scanner Scanner) []Token {
	for !isAtEnd(scanner) {
		scanner.start = scanner.current
		scanToken()
	}

	scanner.Tokens = append(scanner.Tokens, Token{EOF, "", "", scanner.line})
	return scanner.Tokens
}

func isAtEnd(scanner Scanner) bool {
	return scanner.current >= len(scanner.Source)
}
