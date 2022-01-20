package gen

import (
	"bytes"
	"os"
	"time"
)

func GenerateSExpressionDefinitions(sexpressions []string, directory string, packageName string) {
	var buffer bytes.Buffer
	buffer.WriteString(generateHeader())

	// package name
	buffer.WriteString("package ")
	buffer.WriteString(packageName)
	buffer.WriteString("\n\n")

	// SExpression defintion
	buffer.WriteString("type SExpression interface {\n\tGetName() string\n}\n")

	// atom s-expression
	buffer.WriteString(generateAtomSExpression())

	// list s-expression
	buffer.WriteString(generateListSExpression())

	// basic punctuation
	for _, sexpression := range sexpressions {
		buffer.WriteString("\n")
		buffer.WriteString(generateSExpressionBasic(sexpression))
	}

	// write file
	os.WriteFile(directory+"/SExpressions.go", buffer.Bytes(), 0644)
}

func generateSExpressionBasic(name string) string {
	var buffer bytes.Buffer

	buffer.WriteString("type ")
	buffer.WriteString(name)
	buffer.WriteString(" struct{}\n\n")

	buffer.WriteString("func (a ")
	buffer.WriteString(name)
	buffer.WriteString(") GetName() string { \n\treturn \"")
	buffer.WriteString(name)
	buffer.WriteString("\"\n}\n")

	return buffer.String()
}

func generateHeader() string {
	var buffer bytes.Buffer
	buffer.WriteString("// ================================================================================================================== //\n")
	buffer.WriteString("// Auto generated file created at ")
	buffer.WriteString(time.Now().Format(time.RFC850))
	buffer.WriteString(". Do not edit this file, edit generation code.\n")
	buffer.WriteString("// ================================================================================================================== //\n\n")

	return buffer.String()
}

func generateAtomSExpression() string {
	return `
/*
The most elementary type of S-Expression is the atomic symbol.

Definition: An atomic symbol is a string of no more than thirty numerals and capital
letters; the first character must be a letter.

Examples:
	A
	APPLE
	PART2
	EXTRALONGSTRINGOFLETTERS
	A4B66XYZ2
*/
type Atom struct {
	name string
}

func (a Atom) GetName() string {
	return "Atom"
}
`
}

func generateListSExpression() string {
	return `
type List struct {
	sexps []SExpression
}

func (l List) GetName() string {
	return "List"
}
`
}
