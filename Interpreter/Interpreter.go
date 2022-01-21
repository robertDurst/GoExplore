package interpreter

import (
	"bufio"
	"fmt"
	"os"
)

func getNextInput() string {
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return text
}

func RunGoExplore() {
	fmt.Println("Starting GoExplore...")

	for {
		line := getNextInput()
		tokens := tokenize(line)
		finalAnswer := evaluate(tokens)
		fmt.Println(finalAnswer)
	}
}
