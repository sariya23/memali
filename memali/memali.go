package memali

import (
	"fmt"
	"go/format"
	"strings"
)

var builtInDataTypes = [...]string{
	"bool", "int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	"float32", "float64", "complex64", "complex128",
	"string", "byte", "rune",
}

const (
	typeKeyWord = "type"
)

func formatGoCode(code string) (string, error) {
	formattedCode, err := format.Source([]byte(code))
	if err != nil {
		return "", err
	}
	return string(formattedCode), nil
}

// Строка нам не подходит:
//
// - если есть "//", так как это комментарий строка;
//
// - если это определение структуры, то есть в строке есть type
//
// - если строка - }
func FindField(code string) ([]string, error) {
	var res []string
	formattedCode, err := formatGoCode(code)
	if err != nil {
		return nil, err
	}
	elements := strings.Split(formattedCode, "\n")
	for _, v := range elements {
		if strings.Contains(v, "//") {
			continue
		} else if strings.Contains(v, typeKeyWord) {
			continue
		} else if v == "}" || v == "" || v == " " || v == "\n" || v == "\t" {
			continue
		} else {
			fmt.Printf("add %v\n", v)
			res = append(res, strings.TrimSpace(v))
		}
	}
	return res, nil
}
