package main

import (
	"fmt"
	"go/ast"
)

type s struct {
	b complex128
	a int
}

type b struct {
	b complex128
	a int

	qw struct {
		a int
		c complex128
	}
}

func main() {
	// Создаем выражение типа <-chan int (только для получения)
	recvChanExpr := &ast.ChanType{
		Dir:   ast.RECV,                // Направление канала (только получение)
		Value: &ast.Ident{Name: "int"}, // Указываем, что это канал int
	}

	// Вывод результата для проверки
	fmt.Printf("%#v\n", recvChanExpr)
}
