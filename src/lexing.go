package src

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type scanningSession struct {
	currentIndex int
	tokens       []token
	input        string
}

func compileErrMsg(unknownTokens []string) string {
	var sb = strings.Builder{}
	if len(unknownTokens) == 1 {
		sb.WriteString(fmt.Sprintf("Error: Unknown token: \n\t`%s`", unknownTokens[0]))
	} else {
		sb.WriteString("Error: Unknown tokens: \n")
		for _, uexpTok := range unknownTokens {
			sb.WriteString(fmt.Sprintf("\t`%s`,\n", uexpTok))
		}
	}
	return sb.String()
}

func Tokenize(input string) ([]token, error) {
	ss := scanningSession{0, []token{}, input}
	var unknownTokens = []string{}
	for !ss.atEnd() {
		currToken := ss.current()
		err := false
		if currToken == "+" {
			ss.addToken(PLUS, "+", nil)
		} else if currToken == "-" {
			ss.addToken(MINUS, "-", nil)
		} else if currToken == "/" {
			ss.addToken(SLASH, "/", nil)
		} else if currToken == "*" {
			ss.addToken(STAR, "*", nil)
		} else if currToken == "(" {
			ss.addToken(LEFT_PAREN, "(", nil)
		} else if currToken == ")" {
			ss.addToken(RIGHT_PAREN, ")", nil)
		} else if currToken == "^" {
			ss.addToken(POWER, "^", nil)
		} else if currToken == ">" {
			if ss.peekNext() == "=" {
				ss.currentIndex++
				ss.addToken(GREATER_EQUAL, ">=", nil)
			} else {
				ss.addToken(GREATER, ">", nil)
			}
		} else if currToken == "<" {
			if ss.peekNext() == "=" {
				ss.currentIndex++
				ss.addToken(LESS_EQUAL, "<=", nil)
			} else {
				ss.addToken(LESS, "<", nil)
			}
		} else if currToken == "=" {
			ss.currentIndex++
			if ss.atEnd() || ss.current() != "=" {
				err = true
			} else {
				ss.addToken(EQUAL_EQUAL, "==", nil)
			}
		} else if isAlpha(currToken) {
			iden := ss.Boolean()
			if iden == "true" {
				ss.addToken(TRUE, "true", true)
				continue
			} else if iden == "false" {
				ss.addToken(FALSE, "false", false)
				continue
			}
			err = true
			currToken = iden
		} else if isDigit(currToken) {
			ss.consumeDigit()
			continue
		} else if currToken == " " {
			ss.consumeWhiteSpace()
			continue
		} else {
			err = true
		}

		if err {
			unknownTokens = append(unknownTokens, currToken)
		}

		ss.currentIndex++
	}
	if len(unknownTokens) > 0 {
		return nil, errors.New(compileErrMsg(unknownTokens))
	}
	if len(ss.tokens) == 0 {
		return nil, errors.New("idk")
	}
	return ss.tokens, nil
}

func (ss *scanningSession) consumeWhiteSpace() {
	for !ss.atEnd() {
		if ss.current() != " " {
			break
		}
		ss.currentIndex++
	}
}

func (ss *scanningSession) consumeDigit() {
	var lexeme = ss.current()
	ss.currentIndex++
	for !ss.atEnd() {
		c := ss.current()
		if !isDigit(c) {
			break
		}
		lexeme += c
		ss.currentIndex++
	}
	number, _ := strconv.Atoi(string(lexeme))
	ss.addToken(NUM_LITERAL, string(lexeme), float64(number))
}

func isDigit(c string) bool {
	return c >= string('0') && c <= string('9')
}

func isAlpha(c string) bool {
	t := c[0]
	return t >= 'a' && t <= 'z'
}

func (ss *scanningSession) addToken(tokenId int, lexeme string, val any) {
	ss.tokens = append(ss.tokens, token{tokenId, lexeme, val})
}

func (ss *scanningSession) current() string {
	return string(ss.input[ss.currentIndex])
}

func (ss *scanningSession) atEnd() bool {
	return ss.currentIndex >= len(ss.input)
}

func (ss *scanningSession) Boolean() string {
	var lexeme = ss.current()
	ss.currentIndex++
	for !ss.atEnd() {
		c := ss.current()
		if !isAlpha(c) {
			break
		}
		lexeme += c
		ss.currentIndex++
	}
	return lexeme
}

func (ss *scanningSession) peekNext() string {
	ss.currentIndex++
	if ss.atEnd() {
		return "" 
	}
	t := ss.current()
	ss.currentIndex--
	return t
}
