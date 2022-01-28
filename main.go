package main

import (
	"bufio"
	"fmt"
	"os"

	GoExplore "github.com/robertDurst/GoExplore/src"
)

func interpretSingleLine(input string, e GoExplore.Evaluator) (GoExplore.Token, error) {

	le := GoExplore.CreateLexarExecutor()
	ls, err := le.Lex(input)
	if err != nil {
		return nil, fmt.Errorf("[Lexar error]: %s", err)
	}

	tk, err := GoExplore.Tokenize(ls)
	if err != nil {
		return nil, fmt.Errorf("[Tokenizer error]: %s", err)
	}

	return e.Eval(tk), nil
}

func getNextInput() string {
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func main() {
	e := GoExplore.CreateEvaluator()
	for {
		line := getNextInput()
		tk, err := interpretSingleLine(line, e)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(tk.PrettyFormat())
		}
	}
}
