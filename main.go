package main

import (
	"bufio"
	"fmt"
	"interpreter/src"
	"os"
	"strings"
)

func main() {
	runPrompt()
}

func runPrompt() {
	var reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			panic(err)
		}
		if line == "exit" {
			break
		} else {
			runLine(line)
		}
	}
}

func runLine(line string) {
	tokens, err := src.Tokenize(line)
	if err != nil {
		fmt.Println(err)
		return
	}
	st, err := src.Parse(tokens)
	if err != nil {
		fmt.Println(err)
		return
	}
	ans, err := src.Interpret(st)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
