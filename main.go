package main

import (
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
)

func validateGoCode(code string) error {
	fset := token.NewFileSet()
	_, err := parser.ParseFile(fset, "", code, parser.AllErrors)
	return err
}

func formatGoCode(code string) (string, error) {
	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		return "", err
	}
	return string(formattedCode), nil
}

func main() {
	code := `
	package main
	type s  {
		a A
		b string
	}`
	q, ere := formatGoCode(code)
	fmt.Println(q, ere)
	// err := validateGoCode(code)
	// if err != nil {
	// 	fmt.Println("Code is invalid:", err)
	// } else {
	// 	fmt.Println("Code is valid")
	// }
}
