package singlycirclelinkedlist

import "fmt"

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

// 新建一个 List 实例
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			newElement.next = list.first
			list.first = newElement
			list.last = newElement
		} else {
			newElement.next = list.first
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
		list.first = newElement
		if list.size == 0 {
			list.last = newElement
		}
		list.last.next = newElement
		list.size++
	}
}

func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	temp := list.first
	for i := 0; i != index; i, temp = i+1, temp.next {
	}
	return temp.value, true
}

func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}
	var beforeElement *element
	currentElement := list.first
	for i := 0; i != index; i, currentElement = i+1, currentElement.next {
		beforeElement = currentElement
	}

	if currentElement == list.first {
		list.first = currentElement.next
		list.last.next = currentElement.next
	}

	if currentElement == list.last {
		list.last = beforeElement
		list.last.next = list.first
	}

	if beforeElement != nil {
		beforeElement.next = currentElement.next
	}
	currentElement = nil
	list.size--
}

func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	e := list.first
	if e.value == value {
		return 0
	}
	i := 1
	for {
		e = e.next
		if e != list.first {
			if e.value == value {
				return i
			}
			i++
		} else {
			break
		}
	}
	return -1
}

func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	if list.size == 0 {
		return values
	}
	e := list.first
	values[0] = e.value
	fmt.Println("this:", values)
	i := 1
	for {
		fmt.Println("for this:", values)
		e = e.next
		if e != list.first {
			values[i] = e.value
			fmt.Println("for this if:", values)
			i++
		} else {
			break
		}
	}
	return values
}

func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	list.size += len(values)

	var beforeElement *element
	foundElement := list.first
	for i := 0; i != index; i, foundElement = i+1, foundElement.next {
		beforeElement = foundElement
	}

	if foundElement == list.first {
		oldElement := list.first
		for i, value := range values {
			newElement := &element{value: value}
			if i == 0 {
				list.first = newElement
				list.last.next = newElement
			} else {
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		beforeElement.next = oldElement
	} else {
		oldElement := beforeElement.next
		for _, value := range values {
			newElement := &element{value: value}
			beforeElement.next = newElement
			beforeElement = newElement
		}
		beforeElement.next = oldElement
	}
}

func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}

	foundElement := list.first
	for i := 0; i != index; {
		i, foundElement = i+1, foundElement.next
	}
	foundElement.value = value
}

func (list *List) Empty() bool {
	return list.size == 0
}

func (list *List) Size() int {
	return list.size
}

func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}
