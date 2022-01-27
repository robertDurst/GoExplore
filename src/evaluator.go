package GoExplore

import "fmt"

type Evaluator struct {
	functionMap map[string]func([]Token) (Token, error)
	labelMap    map[string]Function
}

func CreateEvaluator() Evaluator {
	fmap := make(map[string]func([]Token) (Token, error))
	fmap["cons"] = cons
	fmap["car"] = car
	fmap["cdr"] = cdr
	fmap["eq"] = eq
	fmap["atom"] = atom

	lmap := make(map[string]Function)

	return Evaluator{functionMap: fmap, labelMap: lmap}
}

func (e Evaluator) Eval(tk Token) Token {
	switch tk.GetType() {
	case "SExpression":
		return e.evalSExpression(tk.(SExpression))
	case "Function":
		return e.evalFunction(tk.(Function))
	case "FunctionLabel":
		return e.evalFunctionLabel(tk.(FunctionLabel))
	case "ConditionalStatement":
		return e.evalConditionalStatement(tk.(ConditionalStatement))
	default:
		return nil
	}
}

func (e Evaluator) evalSExpression(sexp SExpression) Token {
	return sexp
}

func (e Evaluator) evalFunction(fn Function) Token {
	args := make([]Token, 0)
	for _, arg := range fn.Args {
		evaldArg := e.Eval(arg)
		if evaldArg.GetType() == "SExpression" || evaldArg.GetType() == "Ident" {
			args = append(args, evaldArg)
		} else {
			panic("Function arguments must be either a SExpression or a Variable")
		}
	}

	name := fn.Name
	if f, ok := e.functionMap[name]; ok {
		val, _ := f(args)
		return val
	}

	if f, ok := e.labelMap[name]; ok {
		fmt.Printf("We want to execute %s() which has %d args and we've supplied %d args", f.Name, len(f.Args), len(fn.Args))
		return nil
	}

	return nil
}

// returns whether or not was able to successfully "create" function
func (e Evaluator) evalFunctionLabel(fl FunctionLabel) Token {
	name := fl.Name.Value

	// already exists
	if _, ok := e.functionMap[name]; ok {
		return boolToSExpression(false)
	}

	if _, ok := e.labelMap[name]; ok {
		return boolToSExpression(false)
	}

	if fl.Fn.GetType() != "Function" {
		panic("expected function label to be labeling a function")
	}

	e.labelMap[name] = fl.Fn.(Function)

	return boolToSExpression(true)
}

func (e Evaluator) evalConditionalStatement(cs ConditionalStatement) Token {
	for _, cp := range cs.ConditionalPairs {
		isTruthy := e.Eval(cp.Predicate)

		if isTruthy.GetType() == "SExpression" && isTruthy.(SExpression).Value.Value == "T" {
			return e.Eval(cp.Result)
		}
	}

	return nil
}
