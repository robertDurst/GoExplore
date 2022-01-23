package main

import (
	i "GoExplore/interpreter"
	"os"
)

func main() {
	data, _ := os.ReadFile("./example.lisp")
	i.RunGoExplore(string(data))
}
