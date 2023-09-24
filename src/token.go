package src

const (
	PLUS = iota
	MINUS
	SLASH
	STAR
	RIGHT_PAREN
	LEFT_PAREN
	NUM_LITERAL
)

type token struct {
	TokenId int
	Lexeme string
	Value any
}


