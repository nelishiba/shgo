package lexer

import "github.com/negli0/shgo/token"

// Lexer is general lexer type
type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // next to read position
	ch           byte // character under the check
}

// New initialize lexer from input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken return the token with its literal and toke type
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '|':
		tok = newToken(token.PIPE, l.ch)
	case '>':
		tok = newToken(token.REDIROU, l.ch)
	case '<':
		tok = newToken(token.REDIRIN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return ch != ' ' && ch != '\t' && ch != '\r' && ch != '\n' && ch != ';' && ch != '|' && ch != ':' && ch != '>' && ch != '<' && ch != '&' && ch != '?' && ch != '!' && ch != 0
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
