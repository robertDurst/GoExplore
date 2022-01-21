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
		dat, _ := os.ReadFile("./lib.lisp")
		i.RunGoExplore(string(dat))
	}
}
