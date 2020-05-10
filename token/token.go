package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"

	IF   = "IF"
	THEN = "THEN"
	ELSE = "ELSE"
	FI   = "FI"
	FOR  = "FOR"
	DO   = "DO"
	DONE = "DONE"
	IN   = "IN"

	QUESTION  = "?"
	EXCLMTON  = "!"
	SEMICOLON = ";"
	COLON     = ":"
	PIPE      = "|"
	AMPSND    = "&"
	REDIRIN   = "<"
	REDIROU   = ">"
	REDIRAP   = ">>"
)

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	">>":   REDIRAP,
	"if":   IF,
	"then": THEN,
	"else": ELSE,
	"fi":   FI,
	"for":  FOR,
	"in":   IN,
	"do":   DO,
	"done": DONE,
}

func LookupIdent(ident string) TokenType {
	if tt, ok := keywords[ident]; ok {
		return tt
	}
	return IDENT
}
