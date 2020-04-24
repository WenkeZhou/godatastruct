package singlylinkedlist

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	list := &List{}
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list.Add("a")
	list.Add("b", "c")

	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v except %v", actualValue, false)
	}

	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v except %v", actualValue, 3)
	}

	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v except %v", actualValue, "a")
	}

	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v except %v", actualValue, "b")
	}

	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}
}

func TestList_Append(t *testing.T) {
	list := &List{}
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list.Append("a")
	list.Add("b", "c")
	list.Append("d")
	fmt.Println(list.Values())

	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v except %v", actualValue, false)
	}

	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v except %v", actualValue, 4)
	}

	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v except %v", actualValue, "a")
	}

	if actualValue, ok := list.Get(1); actualValue != "b" || !ok {
		t.Errorf("Got %v except %v", actualValue, "b")
	}

	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}

	if actualValue, ok := list.Get(3); actualValue != "d" || !ok {
		t.Errorf("Got %v except %v", actualValue, "d")
	}
}

func TestList_IndexOf(t *testing.T) {
	list := &List{}
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list.Append("a")
	list.Add("b", "c")
	list.Append("d")

	if actualValue := list.IndexOf("a"); actualValue != 0 {
		t.Errorf("Got %v except %v", actualValue, 0)
	}

	if actualValue := list.IndexOf("d"); actualValue != 3 {
		t.Errorf("Got %v except %v", actualValue, 3)
	}

	if actualValue := list.IndexOf("f"); actualValue != -1 {
		t.Errorf("Got %v except %v", actualValue, -1)
	}
}

func TestList_Insert(t *testing.T) {
	list := &List{}
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list.Append("a")
	list.Add("b", "c")
	list.Append("d")

	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}

	if actualValue := list.IndexOf("d"); actualValue != 3 {
		t.Errorf("Got %v except %v", actualValue, 3)
	}

	list.Insert(2, "e", "f")

	if actualValue, ok := list.Get(2); actualValue != "e" || !ok {
		t.Errorf("Got %v except %v", actualValue, "e")
	}

	if actualValue, ok := list.Get(3); actualValue != "f" || !ok {
		t.Errorf("Got %v except %v", actualValue, "f")
	}

	if actualValue := list.IndexOf("c"); actualValue != 4 {
		t.Errorf("Got %v except %v", actualValue, 4)
	}

	if actualValue := list.IndexOf("d"); actualValue != 5 {
		t.Errorf("Got %v except %v", actualValue, 5)
	}
}

func TestList_Prepend(t *testing.T) {
	list := &List{}
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list.Append("a")
	list.Add("b", "c")
	list.Append("d")

	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v except %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}

	list.Prepend("e", "f")

	if actualValue, ok := list.Get(0); actualValue != "e" || !ok {
		t.Errorf("Got %v except %v", actualValue, "e")
	}

	if actualValue, ok := list.Get(1); actualValue != "f" || !ok {
		t.Errorf("Got %v except %v", actualValue, "f")
	}

	if actualValue, ok := list.Get(2); actualValue != "a" || !ok {
		t.Errorf("Got %v except %v", actualValue, "a")
	}

	if actualValue, ok := list.Get(3); actualValue != "b" || !ok {
		t.Errorf("Got %v except %v", actualValue, "b")
	}
}

func TestList_Remove(t *testing.T) {
	list := &List{}
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list.Append("a")
	list.Add("b", "c")
	list.Append("d")

	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v except %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}

	if actualValue := list.Size(); actualValue != 4 {
		t.Errorf("Got %v except %v", actualValue, 4)
	}

	list.Remove(0)

	if actualValue := list.Size(); actualValue != 3 {
		t.Errorf("Got %v except %v", actualValue, 3)
	}

	if actualValue, ok := list.Get(0); actualValue != "b" || !ok {
		t.Errorf("Got %v except %v", actualValue, "b")
	}

	list.Remove(2)

	if actualValue := list.Size(); actualValue != 2 {
		t.Errorf("Got %v except %v", actualValue, 2)
	}

	if actualValue := list.IndexOf("d"); actualValue != -1 {
		t.Errorf("Got %v except %v", actualValue, -1)
	}
}

func TestList_Set(t *testing.T) {
	list := &List{}
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v except %v", actualValue, true)
	}

	list.Append("a")
	list.Add("b", "c")
	list.Append("d")

	if actualValue, ok := list.Get(0); actualValue != "a" || !ok {
		t.Errorf("Got %v except %v", actualValue, "a")
	}
	if actualValue, ok := list.Get(2); actualValue != "c" || !ok {
		t.Errorf("Got %v except %v", actualValue, "c")
	}

	list.Set(0, "A")
	list.Set(2, "C")

	if actualValue, ok := list.Get(0); actualValue != "A" || !ok {
		t.Errorf("Got %v except %v", actualValue, "A")
	}

	if actualValue, ok := list.Get(2); actualValue != "C" || !ok {
		t.Errorf("Got %v except %v", actualValue, "C")
	}
}
