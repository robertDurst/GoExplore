package main

import (
	i "GoExplore/interpreter"
)

func main() {
	// regenerate the latest SExpressions
	// g.GenerateSExpressionDefinitions([]string{"LParen", "RParen", "Dot", "LSquareBracket", "RSquareBracket", "Semicolon"}, "./Interpreter", "interpreter")

	i.RunGoExplore()
}
