package memali

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetTypeWeight(t *testing.T) {
	cases := []struct {
		name          string
		_type         ast.Expr
		expectedWight int
	}{
		{"int", ast.NewIdent("int"), 64},
		{"int64", ast.NewIdent("int64"), 64},
		{"int32", ast.NewIdent("int32"), 32},
		{"int16", ast.NewIdent("int16"), 16},
		{"int8", ast.NewIdent("int8"), 8},
		{"uint", ast.NewIdent("uint"), 64},
		{"uint64", ast.NewIdent("uint64"), 64},
		{"uint32", ast.NewIdent("uint32"), 32},
		{"uint16", ast.NewIdent("uint16"), 16},
		{"uint8", ast.NewIdent("uint8"), 8},
		{"float64", ast.NewIdent("float64"), 64},
		{"float32", ast.NewIdent("float32"), 32},
		{"comlex64", ast.NewIdent("complex64"), 64},
		{"complex128", ast.NewIdent("complex128"), 128},
		{"string", ast.NewIdent("string"), 128},
		{"uintptr", ast.NewIdent("uintptr"), 64},
		{"slice", &ast.ArrayType{Lbrack: token.NoPos, Elt: &ast.Ident{Name: "int"}}, 192},
		{"[3]arr", &ast.ArrayType{
			Lbrack: token.NoPos,
			Elt:    &ast.Ident{Name: "int"},
			Len: &ast.BasicLit{
				Kind:  token.INT,
				Value: "3",
			},
		}, 64},
		{"map", &ast.MapType{
			Map:   token.NoPos,
			Key:   &ast.Ident{Name: "string"},
			Value: &ast.Ident{Name: "string"},
		}, 64},
		{"star", &ast.StarExpr{Star: token.NoPos, X: &ast.Ident{Name: "int"}}, 64},
		{"interface", &ast.InterfaceType{Methods: &ast.FieldList{}}, 128},
		{"func", &ast.FuncType{TypeParams: &ast.FieldList{}, Params: &ast.FieldList{}, Results: &ast.FieldList{}}, 64},
		{"chan", &ast.ChanType{Value: &ast.Ident{Name: "int"}}, 64},
	}
	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := getTypeWeight(tc._type)
			require.Equal(t, tc.expectedWight, w)
		})
	}
}
