package list

import (
	"errors"
)

type Element struct {
	next, prev *Element

	value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

func (e *Element) Value() interface{} {
	return e.value
}

func (e *Element) SetValue(val interface{}) {
	e.value = val
}

type List struct {
	len int64

	root, head *Element
}

func NewList(initVal interface{}) *List {
	el := &Element{value: initVal}
	l := List{root: el, head: el, len: 1}
	return &l
}

func (l *List) ToFront(val interface{}) (*Element, error) {
	if l.len == 0 || l.root == nil {
		return nil, errors.New("list is empty(len == 0)")
	}

	newElem := &Element{next: l.root, prev: nil, value: val}
	l.root.prev = newElem
	l.root = newElem
	l.len++

	return newElem, nil
}

func (l *List) ToBack(val interface{}) *Element {
	newElem := &Element{next: nil, prev: l.head, value: val}

	l.head.next = newElem
	l.head = newElem
	l.len++

	return newElem
}

func (l *List) Head() *Element {
	return l.head
}

func (l *List) Root() *Element {
	return l.root
}
