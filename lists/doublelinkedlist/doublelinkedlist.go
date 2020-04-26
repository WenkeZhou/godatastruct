package doublelinkedlist

import (
	"fmt"
	"strings"
)

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	prev  *element
	next  *element
}

func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value, prev: list.last}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

func (list *List) Prepend(values ...interface{}) {
	for i := len(values) - 1; i >= 0; i-- {
		newElement := &element{value: values[i], next: list.first}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	if list.size-index > index {
		e := list.first
		for i := 0; i != index; i, e = i+1, e.next {
		}
		return e, true
	} else {
		e := list.last
		for i := list.size - 1; i != index; i, e = i+1, e.prev {
		}
		return e.value, true
	}
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}

	var e *element
	if list.size-index > index {
		e = list.first
		for i := 0; i != index; i, e = i+1, e.next {
		}
	} else {
		e = list.last
		for i := list.size - 1; i != index; i, e = i+1, e.prev {
		}
	}

	if e == list.first {
		list.first = e.next
	}
	if e == list.last {
		list.last = e.prev
	}
	if e.prev != nil {
		e.prev.next = e.next
	}
	if e.next != nil {
		e.next.prev = e.prev
	}
	e = nil
	list.size--
}

func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for i, e := 0, list.first; e != nil; i, e = i+1, e.next {
		values[i] = e.value
	}
	return values
}

func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for i, e := 0, list.first; i < list.size; i, e = i+1, e.next {
		if e.value == value {
			return i
		}
	}
	return -1
}

func (list *List) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *List) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
			return
		}
	}

	list.size += len(values)

	sentry := &element{next: list.first}
	var cursorQuick, cursorSlow = list.first, sentry // 快慢指针
	for i := 0; i != index; i++ {
		cursorSlow = cursorSlow.next
		cursorQuick = cursorQuick.next
	}
	for _, value := range values {
		newElement := &element{value: value}
		newElement.prev = cursorSlow
		newElement.next = cursorQuick
		cursorSlow.next = newElement
		cursorQuick.prev = newElement
		cursorSlow = newElement
	}

	list.first = sentry.next
	if list.first.prev != nil {
		list.first.prev = nil
	}
}

func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		return
	}
	if list.size-index > index {
		e := list.first
		for i := 0; i != index; i, e = i+1, e.next {
		}
		e.value = value
	} else {
		e := list.last
		for i := list.size - 1; i != index; i, e = i+1, e.prev {
		}
		e.value = value
	}
}

func (list *List) String() string {
	var values []string
	for e := list.first; e != nil; e = e.next {
		values = append(values, fmt.Sprintf("%v", e.value))
	}
	return strings.Join(values, ", ")
}

// Check that the index is within bounds of the list
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
