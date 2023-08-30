package main

import (
	"fmt"
	list "thinknetica_go/hw-04/1-list"
)

func main() {
	// демонстрация работы функций реализованного списка
	var l *list.List
	l = list.New()

	var elem list.Elem
	elem.Val = 1
	l.Push(elem)

	elem.Val = 2
	l.Push(elem)

	elem.Val = 3
	l.Push(elem)

	fmt.Println(l.String())
	fmt.Println(l.Reverse().String())

	l.Pop()

	fmt.Println(l.String())
	fmt.Println(l.Reverse().String())

}
