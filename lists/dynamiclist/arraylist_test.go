package arraylist

import (
	"fmt"
	"testing"
)

// go test -v -run="Insert"

func TestNew(t *testing.T) {
	list1 := New(10)
	if actualValue := list1.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list2 := New(10, "b", 2)

	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	if actualValue, ok := list2.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v except %v", actualValue, "b")
	}

	if actualValue, ok := list2.Get(1); actualValue != 2 || !ok {
		t.Errorf("Got %v except %v", actualValue, 2)
	}

	if actualValue, ok := list2.Get(2); actualValue != nil || ok {
		t.Errorf("Got %v except %v", actualValue, nil)
	}
}

func TestList_Add(t *testing.T) {
	list1 := New(5)
	list1.Add("a")
	list1.Add("b", "c")
	if actualValue := list1.Empty(); actualValue != false {
		t.Errorf("Got %v except %v", actualValue, false)
	}

	if actualValue := list1.Size(); actualValue != 3 {
		t.Errorf("Got %v except %v", actualValue, 3)
	}

	if actualValue, ok := list1.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}
	list1.Add("d")
	list1.Add("e")

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic recover! p: %v\n", p)
			//debug.PrintStack()
		}
	}()
	list1.Add("f")
}

func TestList_IndexOf(t *testing.T) {
	list := New(10)
	list.Add("a")
	list.Add("b", "c", "d")

	if actualValue := list.IndexOf("a"); actualValue != 0 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	if actualValue := list.IndexOf("f"); actualValue != -1 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}
}

func TestList_Remove(t *testing.T) {
	list := New(10, "a")
	list.Add("b")
	list.Add("c", "d")

	if actualValue := list.IndexOf("a"); actualValue != 0 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	list.Remove(0)

	if actualValue := list.IndexOf("a"); actualValue != -1 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	if actualValue := list.IndexOf("b"); actualValue != 0 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	if actualValue := list.IndexOf("c"); actualValue != 1 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	if actualValue := list.IndexOf("d"); actualValue != 2 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v except %v", actualValue, 3)
	}

	list.Remove(2)
	list.Remove(1)
	list.Remove(0)
	list.Remove(0)

	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	if actualValue := list.Size(); actualValue != 0 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}
}

func TestList_Insert(t *testing.T) {
	list := New(10)
	list.Add("a", "b", "c", "d")
	list.Insert(2, "A", "B")

	fmt.Println(list.elements)

	if actualValue := list.Size(); actualValue != 6 {
		t.Errorf("Got %v except %v", actualValue, 6)
	}

	if actualValue, ok := list.Get(2); actualValue != "A" || !ok {
		t.Errorf("Got %v except %v", actualValue, "A")
	}

	if actualValue, ok := list.Get(3); actualValue != "B" || !ok {
		t.Errorf("Got %v except %v", actualValue, "B")
	}

	if actualValue, ok := list.Get(4); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}
}

func doubleValue(x int) int {
	return x * x
}

//func TestList_ForAll(t *testing.T) {
//	list := New(10)
//	list.Add(1, 2, 3, 4)
//
//	list.ForAll(doubleValue)
//
//	for index, value := range list.elements {
//		fmt.Printf("%d: %v\n", index, value)
//	}
//
//}

func benchmarkGet(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Get(n)
		}
	}
}

func benchmarkAdd(b *testing.B, list *List, size int) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < size; n++ {
			list.Add(n)
		}
		list.size = 0
		list.elements = make([]interface{}, size, size)
	}
}

func BenchmarkArrayListGet100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New(100)
	for n := 1; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New(10000)
	for n := 1; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListGet100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	list := New(100000)
	for n := 1; n < size; n++ {
		list.Add(n)
	}
	b.StartTimer()
	benchmarkGet(b, list, size)
}

func BenchmarkArrayListAdd100(b *testing.B) {
	b.StopTimer()
	size := 100
	list := New(size)
	b.StartTimer()
	benchmarkAdd(b, list, size)
}

func BenchmarkArrayListAdd10000(b *testing.B) {
	b.StopTimer()
	size := 10000
	list := New(size)
	b.StartTimer()
	benchmarkAdd(b, list, size)
}
