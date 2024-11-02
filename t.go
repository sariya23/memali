package main

import (
	"fmt"
	"go/ast"
)

func main() {
	// Создаем выражение типа <-chan int (только для получения)
	recvChanExpr := &ast.ChanType{
		Dir:   ast.RECV,                // Направление канала (только получение)
		Value: &ast.Ident{Name: "int"}, // Указываем, что это канал int
	}

	// Вывод результата для проверки
	fmt.Printf("%#v\n", recvChanExpr)
}
