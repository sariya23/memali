package memali

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"runtime"
	"sort"
)

func getTypeWeight(expr ast.Expr) int {
	wordSize := getWordSize()
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
		if t.Len == nil {
			return 3 * wordSize
		}
		return getTypeWeight(t.Elt)
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

func getWordSize() int {
	switch runtime.GOARCH {
	case "amd64", "arm64":
		return 64
	case "386", "arm":
		return 32
	default:
		return 64

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
	cq [3]complex128       
    m    map[int]int  
	ui uintptr 
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
					typeI := getTypeWeight(structType.Fields.List[i].Type)
					typeJ := getTypeWeight(structType.Fields.List[j].Type)
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
