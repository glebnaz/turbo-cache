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

//ToFront create new element in root
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

//ToBack create new element on head
func (l *List) ToBack(val interface{}) *Element {
	newElem := &Element{next: nil, prev: l.head, value: val}

	l.head.next = newElem
	l.head = newElem
	l.len++

	return newElem
}

//PushToFront replace this element to front
func (l *List) PushToFront(el *Element) (*Element, error) {
	if el == nil {
		return nil, errors.New("element is nil")
	}
	if l.len == 1 {
		return nil, errors.New("len of list == 1, root == head")
	}
	if l.root == el {
		return el, nil
	}

	if el.next != nil {
		el.next.prev = el.prev
	}
	if el.prev != nil {
		el.prev.next = el.next
	}

	l.root.prev = el
	el.next = l.root
	el.prev = nil
	l.root = el
	if l.head == el {
		l.head = el.next
	}

	return el, nil
}

//PushToBack replace this element to back
func (l *List) PushToBack(el *Element) (*Element, error) {
	if el == nil {
		return nil, errors.New("element is nil")
	}
	if l.len == 1 {
		return nil, errors.New("len of list == 1, root == head")
	}
	if l.head == el {
		return el, nil
	}

	if el.next != nil {
		el.next.prev = el.prev
	}
	if el.prev != nil {
		el.prev.next = el.next
	}

	l.head.next = el
	el.prev = l.head
	el.next = nil
	l.head = el
	if l.root == el {
		l.root = el.prev
	}

	return el, nil
}

func (l *List) Head() *Element {
	return l.head
}

func (l *List) Root() *Element {
	return l.root
}
