package main

import (
	"fmt"
	"memali/memali"
	"strings"
)

func main() {
	code := `// Some struct
				type s struct {


				// Comment
				a bool


				// Comment
				
				// long

				q string
	}
			`

	e, err := memali.FindField(code)
	fmt.Println(strings.Join(e, "."))
	fmt.Println(len(e[2]), e[2])
	fmt.Println(err)
}
