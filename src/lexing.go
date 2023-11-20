package src

import (
	"errors"
	"strconv"
	"strings"
	"fmt"
)

type scanningSession struct {
	currentIndex int 
	tokens []token
	input string
}

func compileErrMsg(unexpectedTokens []string) string {
	var sb = strings.Builder {}
	if len(unexpectedTokens) == 1 {
		sb.WriteString(fmt.Sprintf("Error: unexpected token: \n\t`%s`\n", unexpectedTokens[0]))
	} else {
		sb.WriteString("Error: Unexpected tokens: \n")
		for _, uexpTok := range unexpectedTokens {
			sb.WriteString(fmt.Sprintf("\t`%s`,\n", uexpTok))
		}
	}
	return sb.String()
}

func Tokenize(input string) ([]token, error) {
	ss := scanningSession {0, []token{}, input}
	var unexpectedTokens = []string {}
	for !ss.atEnd() {
		c := ss.current()
		var err error
		if c == "+" {
			ss.addToken(PLUS, "+", nil)
		} else if c == "-" {
			ss.addToken(MINUS, "-", nil)
		} else if c == "/" {
			ss.addToken(SLASH, "/", nil)
		} else if c == "*" {
			ss.addToken(STAR, "*", nil)
		} else if c == "(" {
			ss.addToken(LEFT_PAREN, "(", nil)
		} else if c == ")" {
			ss.addToken(RIGHT_PAREN, ")", nil)
		} else if isDigit(c) {
			err = ss.consumeDigit()
			if err != nil {
				unexpectedTokens = append(unexpectedTokens, string(err.Error()))
			}
			continue
		} else if c == " " { 
			ss.consumeWhiteSpace()
			continue
		} else {
			unexpectedTokens = append(unexpectedTokens, c)
			// errMsg := fmt.Sprintf("Error: unexpected token: %s\n", c)
			// return nil, errors.New(errMsg)
		}
		// if err != nil {
		// 	return nil, err
		// }
		ss.currentIndex++
	}
	if len(unexpectedTokens) > 0 {
		return nil, errors.New(compileErrMsg(unexpectedTokens))
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

func (ss *scanningSession) consumeDigit() error {
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
	number, err := strconv.Atoi(string(lexeme))
	if err != nil {
		return errors.New("Error: could not convert number")
	}
	ss.addToken(NUM_LITERAL, string(lexeme), number)
	return nil
}

func isDigit(c string) bool {
	return c >= string('0') && c <= string('9')
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
