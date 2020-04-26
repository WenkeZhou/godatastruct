package doublelinkedlist

import (
	"testing"
)

func TestList_New(t *testing.T) {
	list1 := New()
	list1.Append(1)
	if actualValue := list1.String(); actualValue != "1" {
		t.Errorf("Got %v except %v", actualValue, "1")
	}
	list2 := New(1, 2, 3)
	if actualValue := list2.String(); actualValue != "1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3")
	}
}

func TestList_Add(t *testing.T) {
	list1 := New()
	list1.Add(1)
	if actualValue := list1.String(); actualValue != "1" {
		t.Errorf("Got %v except %v", actualValue, "1")
	}
	list2 := New(1, 2)
	list2.Add(3, 4)
	if actualValue := list2.String(); actualValue != "1, 2, 3, 4" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3, 4")
	}
}

func TestList_Append(t *testing.T) {
	list1 := New()
	list1.Add(1, 2, 3)
	if actualValue := list1.String(); actualValue != "1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3")
	}
	list1.Append(4, 5, 6)
	if actualValue := list1.String(); actualValue != "1, 2, 3, 4, 5, 6" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3, 4, 5, 6")
	}
	list2 := New()
	list2.Append(1, 2, 3)
	if actualValue := list2.String(); actualValue != "1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3")
	}
}

func TestList_Prepend(t *testing.T) {
	list1 := New()
	list1.Add(1, 2, 3)
	if actualValue := list1.String(); actualValue != "1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3")
	}
	list1.Prepend(4, 5, 6)
	if actualValue := list1.String(); actualValue != "4, 5, 6, 1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "4, 5, 6, 1, 2, 3")
	}
	list2 := New()
	list2.Prepend(1, 2, 3)
	if actualValue := list2.String(); actualValue != "1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3")
	}
}

func TestList_Remove(t *testing.T) {
	list1 := New()
	list1.Add(1, 2, 3)
	if actualValue := list1.String(); actualValue != "1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3")
	}
	list1.Remove(0)
	if actualValue := list1.String(); actualValue != "2, 3" {
		t.Errorf("Got %v except %v", actualValue, "2, 3")
	}

	list1.Remove(1)
	if actualValue := list1.String(); actualValue != "2" {
		t.Errorf("Got %v except %v", actualValue, "2")
	}
}

func TestList_Insert(t *testing.T) {
	list1 := New()
	list1.Insert(0, 1, 2, 3)
	if actualValue := list1.String(); actualValue != "1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "1, 2, 3")
	}

	list1.Insert(0, 4, 5)
	if actualValue := list1.String(); actualValue != "4, 5, 1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "4, 5, 1, 2, 3")
	}

	list1.Insert(1, 6, 7)
	if actualValue := list1.String(); actualValue != "4, 6, 7, 5, 1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "4, 6, 7, 5, 1, 2, 3")
	}

	list1.Insert(7, 8, 9)
	if actualValue := list1.String(); actualValue != "4, 6, 7, 5, 1, 2, 3, 8, 9" {
		t.Errorf("Got %v except %v", actualValue, "4, 6, 7, 5, 1, 2, 3, 8, 9")
	}

	list2 := New()
	list2.Add(1, 2, 3)
	list2.Insert(0, 4, 5)
	if actualValue := list2.String(); actualValue != "4, 5, 1, 2, 3" {
		t.Errorf("Got %v except %v", actualValue, "4, 5, 1, 2, 3")
	}
}

func TestList_IndexOf(t *testing.T) {
	list1 := New(1, 2, 3)
	if actualValue := list1.IndexOf(2); actualValue != 1 {
		t.Errorf("Got %v except %v", actualValue, 1)
	}

}
