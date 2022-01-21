package interpreter

import (
	"fmt"
)

func evaluate(sexps []SExpression, level int) SExpression {
	cur := sexps[0]

	fmt.Printf("[%d][%s]: %s\n", level, cur.GetName(), cur)

	// evaluate list
	if cur.GetName() == "List" {
		list := cur.(List)
		if len(list.sexps) == 0 {
			return cur
		}

		if list.sexps[0].GetName() == "Atom" {
			atom := list.sexps[0].(Atom)
			if atom.name != "cons" && atom.name != "car" {
				fmt.Printf("[%d][NOT CONS OR CAR, RETURN %s]: %s\n", level, cur.GetName(), cur)
				return cur
			}
		}

		return List{sexps: []SExpression{evaluate(list.sexps, level+1)}}
	}

	// evaluate atom
	if cur.GetName() == "Atom" {
		atom := cur.(Atom)
		switch atom.name {
		case "cons":
			fmt.Printf("[%d] Eval arg1\n", level)
			arg1 := evaluate([]SExpression{sexps[1]}, level+1)
			fmt.Printf("[%d]Eval arg2\n", level)
			arg2 := evaluate([]SExpression{sexps[2]}, level+1)
			return List{sexps: []SExpression{arg1, arg2}}
		case "car":
			fmt.Printf("[%d] Eval arg1\n", level)
			arg1 := evaluate([]SExpression{sexps[1]}, level+1)

			if arg1.GetName() != "List" {
				panic("Expected a list after car")
			}
			list := arg1.(List)

			if len(list.sexps) == 0 {
				panic("Tried to car an empty list")
			}

			return list.sexps[0]
		default:
			fmt.Printf("[%d][NOT CONS OR CAR, RETURN %s]: %s\n", level, cur.GetName(), cur)
			return cur
		}
	}

	panic("Unexpected SExpression encountered!")
}
