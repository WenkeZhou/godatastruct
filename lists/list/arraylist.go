package arraylist

import (
	"fmt"
)

// 数组表(顺序表)的最大长度
const MaxSize = 50

type List struct {
	elements []interface{}
	size     int
	maxsize  int
}

func New(maxsize int, values ...interface{}) *List {
	list := &List{}
	list.elements = make([]interface{}, maxsize, maxsize)
	list.maxsize = maxsize
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	if list.size+len(values) > list.maxsize {
		panic(fmt.Sprintf("out of range, maxsize is %d", list.maxsize))
	}
	for _, v := range values {
		list.elements[list.size] = v
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

func (list *List) Remove(index int) {
	if !list.withRange(index) {
		return
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
}

func (list *List) Insert(index int, values ...interface{}) {
	l := len(values)

	if !list.withRange(index) {
		panic("out of range")
	}

	if list.size+l > list.maxsize {
		panic("out of range, list is full")
	}

	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if element == value {
			return index
		}
	}
	return -1
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) withRange(index int) bool {
	return index >= 0 && index < list.size
}
