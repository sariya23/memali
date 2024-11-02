package memali

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	wordSize64 = 64
	wordSize32 = 32
)

func TestGetTypeWeight64x(t *testing.T) {
	cases := []struct {
		name          string
		_type         ast.Expr
		expectedWight int
	}{
		{"int", ast.NewIdent("int"), wordSize64},
		{"int64", ast.NewIdent("int64"), 64},
		{"int32", ast.NewIdent("int32"), 32},
		{"int16", ast.NewIdent("int16"), 16},
		{"int8", ast.NewIdent("int8"), 8},
		{"uint", ast.NewIdent("uint"), wordSize64},
		{"uint64", ast.NewIdent("uint64"), 64},
		{"uint32", ast.NewIdent("uint32"), 32},
		{"uint16", ast.NewIdent("uint16"), 16},
		{"uint8", ast.NewIdent("uint8"), 8},
		{"float64", ast.NewIdent("float64"), 64},
		{"float32", ast.NewIdent("float32"), 32},
		{"comlex64", ast.NewIdent("complex64"), 64},
		{"complex128", ast.NewIdent("complex128"), 128},
		{"string", ast.NewIdent("string"), wordSize64 * 2},
		{"uintptr", ast.NewIdent("uintptr"), wordSize64},
		{"slice", &ast.ArrayType{Lbrack: token.NoPos, Elt: &ast.Ident{Name: "int"}}, wordSize64 * 3},
		{"[3]arr", &ast.ArrayType{
			Lbrack: token.NoPos,
			Elt:    &ast.Ident{Name: "int"},
			Len: &ast.BasicLit{
				Kind:  token.INT,
				Value: "3",
			},
		}, wordSize64},
		{"map", &ast.MapType{
			Map:   token.NoPos,
			Key:   &ast.Ident{Name: "string"},
			Value: &ast.Ident{Name: "string"},
		}, wordSize64},
		{"star", &ast.StarExpr{Star: token.NoPos, X: &ast.Ident{Name: "int"}}, wordSize64},
		{"interface", &ast.InterfaceType{Methods: &ast.FieldList{}}, wordSize64 * 2},
		{"func", &ast.FuncType{TypeParams: &ast.FieldList{}, Params: &ast.FieldList{}, Results: &ast.FieldList{}}, wordSize64},
		{"chan", &ast.ChanType{Value: &ast.Ident{Name: "int"}}, wordSize64},
	}
	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := getTypeWeight(tc._type, wordSize64)
			require.Equal(t, tc.expectedWight, w)
		})
	}
}

func TestGetTypeWeight32x(t *testing.T) {
	cases := []struct {
		name          string
		_type         ast.Expr
		expectedWight int
	}{
		{"int", ast.NewIdent("int"), wordSize32},
		{"int64", ast.NewIdent("int64"), 64},
		{"int32", ast.NewIdent("int32"), 32},
		{"int16", ast.NewIdent("int16"), 16},
		{"int8", ast.NewIdent("int8"), 8},
		{"uint", ast.NewIdent("uint"), wordSize32},
		{"uint64", ast.NewIdent("uint64"), 64},
		{"uint32", ast.NewIdent("uint32"), 32},
		{"uint16", ast.NewIdent("uint16"), 16},
		{"uint8", ast.NewIdent("uint8"), 8},
		{"float64", ast.NewIdent("float64"), 64},
		{"float32", ast.NewIdent("float32"), 32},
		{"comlex64", ast.NewIdent("complex64"), 64},
		{"complex128", ast.NewIdent("complex128"), 128},
		{"string", ast.NewIdent("string"), wordSize32 * 2},
		{"uintptr", ast.NewIdent("uintptr"), wordSize32},
		{"slice", &ast.ArrayType{Lbrack: token.NoPos, Elt: &ast.Ident{Name: "int"}}, wordSize32 * 3},
		{"[3]arr", &ast.ArrayType{
			Lbrack: token.NoPos,
			Elt:    &ast.Ident{Name: "int"},
			Len: &ast.BasicLit{
				Kind:  token.INT,
				Value: "3",
			},
		}, wordSize32},
		{"map", &ast.MapType{
			Map:   token.NoPos,
			Key:   &ast.Ident{Name: "string"},
			Value: &ast.Ident{Name: "string"},
		}, wordSize32},
		{"star", &ast.StarExpr{Star: token.NoPos, X: &ast.Ident{Name: "int"}}, wordSize32},
		{"interface", &ast.InterfaceType{Methods: &ast.FieldList{}}, wordSize32 * 2},
		{"func", &ast.FuncType{TypeParams: &ast.FieldList{}, Params: &ast.FieldList{}, Results: &ast.FieldList{}}, wordSize32},
		{"chan", &ast.ChanType{Value: &ast.Ident{Name: "int"}}, wordSize32},
	}
	t.Parallel()
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := getTypeWeight(tc._type, wordSize32)
			require.Equal(t, tc.expectedWight, w)
		})
	}
}
