package interpreter

func tokenize(line string) []SExpression {
	tokens := lex(line)

	sexps := make([]SExpression, 0)
	inList := 0
	listQueue := make([]List, 0)
	for _, token := range tokens {

		switch token.GetName() {
		case "Atom":
			if inList > 0 {
				listQueue[inList-1].sexps = append(listQueue[inList-1].sexps, token)
			} else {
				sexps = append(sexps, token)
			}
		case "LParen":
			listQueue = append(listQueue, List{sexps: make([]SExpression, 0)})
			inList++
		case "RParen":
			if inList == 0 {
				panic("tried to close a list when not open")
			} else {
				inList--
				if inList == 0 {
					sexps = append(sexps, listQueue[len(listQueue)-1])
				} else {
					listQueue[len(listQueue)-2].sexps = append(listQueue[len(listQueue)-2].sexps, listQueue[len(listQueue)-1])
				}
				listQueue = listQueue[:len(listQueue)-1]
			}
		}
	}

	if inList != 0 {
		panic("expected a closing paren")
	}

	return sexps
}
