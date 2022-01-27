package GoExplore

import "bytes"

type SExpression struct {
	Value Lexicon
}

func (s SExpression) GetType() string {
	return "SExpression"
}

func (s SExpression) PrettyFormat() string {
	if s.Value.Type == Atom {
		return s.Value.Value
	}

	return prettyFormatList(s.Value.ListValues)
}

func prettyFormatList(lexs []Lexicon) string {
	var listBuffer bytes.Buffer
	listBuffer.WriteByte('(')
	for i, l := range lexs {
		if i != 0 {
			listBuffer.WriteByte(' ')
		}
		if l.Type == List {
			listBuffer.WriteString(prettyFormatList(l.ListValues))
		} else {
			listBuffer.WriteString(l.Value)
		}
	}
	listBuffer.WriteByte(')')

	return listBuffer.String()
}

func CreateSExpression(value Lexicon) SExpression {
	return SExpression{Value: value}
}
