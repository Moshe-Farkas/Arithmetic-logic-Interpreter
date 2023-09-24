package src

import (
	"errors"
	"strconv"
)

type scanningSession struct {
	currentIndex int 
	tokens []token
	input string
}

func Tokenize(input string) ([]token, error) {
	ss := scanningSession {0, []token{}, input}
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
		} else if isDigit(c) {
			err = ss.consumeDigit()
			continue
		} else if c == " " { 
			ss.consumeWhiteSpace()
			continue
		} else {
			return nil, errors.New("Error: unexpected token\n")
		}
		if err != nil {
			return nil, err
		}
		ss.currentIndex++
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
