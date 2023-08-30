// Реализуация двусвязного списка вместе с базовыми операциями.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New создаёт список и возвращает указатель на него.
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	el := l.root.next
	new_el := l.root.next.next
	new_el.prev = l.root
	l.root.next = new_el
	el.next = nil
	el.prev = nil
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	var new_list List
	new_list.root = &Elem{}
	new_list.root.prev = new_list.root
	new_list.root.next = new_list.root

	el := l.root.next
	for el != l.root {
		new_list.Push(*el)
		el = el.next
	}
	return &new_list
}
