package src

const (
	PLUS = iota
	MINUS
	SLASH
	STAR
	RIGHT_PAREN
	LEFT_PAREN
	NUM_LITERAL
	POWER
	EQUAL_EQUAL
	TRUE
	FALSE
	GREATER
	LESS
	GREATER_EQUAL
	LESS_EQUAL
)

type token struct {
	TokenId int
	Lexeme string
	Value any
}


