package memali

import (
	"go/ast"
	"runtime"
	"sort"
)

func getTypeWeight(expr ast.Expr, wordSize int) int {
	var typeAndItWeight = map[string]int{
		"int":        wordSize,
		"int8":       8,
		"int16":      16,
		"int32":      32,
		"int64":      64,
		"uint8":      8,
		"uint16":     16,
		"uint32":     32,
		"uint64":     64,
		"uint":       wordSize,
		"string":     wordSize * 2,
		"float32":    32,
		"float64":    64,
		"complex64":  64,
		"complex128": 128,
		"uintptr":    wordSize,
	}

	switch t := expr.(type) {
	case *ast.Ident:
		return typeAndItWeight[t.Name]
	case *ast.ArrayType:
		// slice case
		if t.Len == nil {
			return 3 * wordSize
		}
		return getTypeWeight(t.Elt, wordSize)
	case *ast.MapType:
		return wordSize
	case *ast.StarExpr:
		return wordSize
	case *ast.ChanType:
		return wordSize
	case *ast.StructType:
		return -1
	case *ast.FuncType:
		return wordSize
	case *ast.InterfaceType:
		return wordSize * 2
	default:
		return -1
	}
}

func GetWordSize() int {
	switch runtime.GOARCH {
	case "amd64", "arm64":
		return 64
	case "386", "arm":
		return 32
	default:
		return 64

	}
}

func SortStructFields(structType *ast.StructType, wordSize int) {
	sort.Slice(structType.Fields.List, func(i, j int) bool {
		wI := getTypeWeight(structType.Fields.List[i].Type, wordSize)
		wJ := getTypeWeight(structType.Fields.List[j].Type, wordSize)
		if wI == wJ {
			return
		}
		return wI > wJ
	})
}
