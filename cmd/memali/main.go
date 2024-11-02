package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"memali/internal/memali"
	"os"
)

func main() {
	filename := mustParseFilePath()
	wordSize := memali.GetWordSize()
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	ast.Inspect(node, func(n ast.Node) bool {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			return true
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				memali.SortStructFields(structType, wordSize)
			}
		}
		return true
	})

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		log.Fatalf("formatting node: %v", err)
	}

	if err := os.WriteFile(filename, buf.Bytes(), 0644); err != nil {
		log.Fatalf("writing file: %v", err)
	}

	fmt.Printf("Struct fields in %s have been sorted memory alignmenty.\n", filename)
}

func mustParseFilePath() string {
	var filename string
	flag.StringVar(&filename, "f", "", "path to file")
	flag.Parse()

	if filename == "" {
		log.Fatalln("path to file (--f) is required flag. Please specify it")
	}

	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("file does not exists: %v\n", err)
	}
	return filename
}
