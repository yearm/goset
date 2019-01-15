package structure

import (
	"fmt"
	"testing"
)

func TestNewDList(t *testing.T) {
	dlist := NewDList()
	dlist.Append("yearm")
	dlist.Append("male")
	dlist.Append(21)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
}

func TestDList_Append(t *testing.T) {
	dlist := NewDList()
	dlist.Append("yearm")
	dlist.Append("male")
	dlist.Append(21)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
}

func TestDList_Insert(t *testing.T) {
	dlist := NewDList()
	dlist.Append("yearm")
	dlist.Append("male")
	dlist.Insert(1, "gopher")
	dlist.Append(21)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
}

func TestDList_Get(t *testing.T) {
	dlist := NewDList()
	dlist.Append("yearm")
	dlist.Append("male")
	dlist.Append(21)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
}

func TestDList_Remove(t *testing.T) {
	dlist := NewDList()
	dlist.Append("yearm")
	dlist.Append("male")
	dlist.Append(21)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
	dlist.Remove(2)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
}

func TestDList_RemoveAll(t *testing.T) {
	dlist := NewDList()
	dlist.Append("yearm")
	dlist.Append("male")
	dlist.Append(21)
	dlist.RemoveAll()
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
}

func TestDList_Update(t *testing.T) {
	dlist := NewDList()
	dlist.Append("yearm")
	dlist.Append("male")
	dlist.Append(21)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
	dlist.Update(2, 22)
	for i := 0; int64(i) < dlist.Size; i++ {
		fmt.Println(dlist.Get(int64(i)))
	}
}
