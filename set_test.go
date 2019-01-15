package structure

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New("yearm", "male", 21)
	s.List(func(items interface{}) {
		fmt.Println(items)
	})
}

func TestSet_Add(t *testing.T) {
	s := New()
	s.Add("yearm")
	s.Add("male")
	s.Add(21)
	s.List(func(items interface{}) {
		fmt.Println(items)
	})
}

func TestSet_Remove(t *testing.T) {
	s := New()
	s.Add("yearm")
	s.Add("male")
	s.Add(21)
	s.Remove(21)
	s.List(func(items interface{}) {
		fmt.Println(items)
	})
}

func TestSet_Size(t *testing.T) {
	s := New("yearm", "male", 21)
	fmt.Println(s.Size())
}

func TestSet_Contains(t *testing.T) {
	s := New("yearm", "male", 21)
	fmt.Println(s.Contains(21))
}

func TestSet_Clear(t *testing.T) {
	s := New("yearm", "male", 21)
	s.Clear()
	s.List(func(items interface{}) {
		fmt.Println(items)
	})
}

func TestSet_List(t *testing.T) {
	s := New("yearm", "male", 21)
	s.List(func(items interface{}) {
		fmt.Println(items)
	})
}
