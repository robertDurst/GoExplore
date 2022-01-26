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

func RunGoExploreInterpreter() {
	fmt.Println("Starting GoExplore...")

	for {
		// line := getNextInput()
		// tokens := tokenize(line)
		// for _, token := range tokens {
		// 	fmt.Printf("[%s]: %s\n", token.GetName(), token)
		// }
	}
}

func RunGoExplore(input string) {
	fmt.Println("======= Source =======")
	fmt.Println(input)
	fmt.Printf("======================\n\n")

	lexicons := lex(input)
	for _, l := range lexicons {
		fmt.Printf("[%s]: %s\n", l.GetType(), l)
	}
	// tokenizer := CreateTokenizer()
	// token := tokenizer.tokenize(lexicons)
	// fmt.Printf("[%s]: %s", token.GetType(), token)
}
