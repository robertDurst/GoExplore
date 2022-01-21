package main

import (
	g "GoExplore/gen"
	i "GoExplore/interpreter"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "gen_sed" {
		// regenerate the latest SExpressions
		g.GenerateSExpressionDefinitions([]string{"LParen", "RParen"}, "./Interpreter", "interpreter")
	} else {
		// data, _ := os.ReadFile("./lib.lisp")
		data := `
cons (car ((car (2 3)))) (4)
		`
		i.RunGoExplore(string(data))
	}
}
