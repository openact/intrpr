package lexer

import "intrpr/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// ASCII code for "NUL"
		l.ch = 0
	} else {
		// Get the next character
		l.ch = l.input[l.readPosition]
	}
	// Advance the position
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		// ASCII code for "NUL"
		return 0
	} else {
		// Get the next character
		return l.input[l.readPosition]
	}
}
func isLetter(ch byte) bool {
	// Check if the character is a letter
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	// Save the current position
	position := l.position
	// Read the whole identifier
	for isLetter(l.ch) {
		l.readChar()
	}
	// Return the identifier
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	// Skip the whitespace
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
func newToken(tokenType token.TokenType, ch byte) token.Token {
	// Return a new token
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]

}

func isDigit(ch byte) bool {
	// Check if the character is a digit
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// Read the whole identifier
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// AND --> && ...
			tok.Alias = token.LookupAlias(tok.Literal)
			tok.Literal = tok.Alias
			return tok
		} else {
			// Read the whole number
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
	}

	l.readChar()
	return tok
}
