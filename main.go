package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
)

const word = 64

var d = map[string]int{
	"int":        word,
	"int8":       8,
	"int16":      16,
	"int32":      32,
	"int64":      64,
	"uint8":      8,
	"uint16":     16,
	"uint32":     32,
	"uint64":     64,
	"uint":       word,
	"string":     word * 2,
	"float32":    32,
	"float64":    64,
	"complex64":  64,
	"complex128": 128,
}

// Неправильно сортируются поля
func getTypeName(expr ast.Expr) int {
	switch t := expr.(type) {
	case *ast.Ident:
		return d[t.Name]
	case *ast.ArrayType:
		if t.Len == nil {
			return 3 * word
		}
		return getTypeName(t.Elt)
	case *ast.MapType:
		return word
	case *ast.StarExpr:
		return word
	case *ast.ChanType:
		return word
	case *ast.StructType:
		return -1
	case *ast.FuncType:
		return word
	case *ast.InterfaceType:
		return word * 2
	default:
		return -1
	}
}

func main() {
	src := `
		package main
type MyStruct struct {
    p    *int        
    f64  float64     
    c128 complex128   
    s    string      
    arr  [3]string      
    sl   []int       
    m    map[int]int   
    ch   chan int      
    i32  int32         
    f32  float32     
    i16  int16         
    u16  uint16       
    i8   int8         
    u8   uint8         
    b    bool      
}
	`

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse code: %v", err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok {
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				sort.Slice(structType.Fields.List, func(i, j int) bool {
					typeI := getTypeName(structType.Fields.List[i].Type)
					typeJ := getTypeName(structType.Fields.List[j].Type)
					return typeI > typeJ
				})
			}
		}
		return true
	})

	fmt.Println("\nModified Source Code:")
	err = format.Node(os.Stdout, fset, node)
	if err != nil {
		log.Fatalf("Failed to format modified code: %v", err)
	}
}
