package util

import (
	"fmt"
	"testing"
)

type Int int

func (this *Int) HashCode() int64 {
	p := GetPointer(this)
	fmt.Println("addr", p)
	return p
}

func TestAdd(t *testing.T) {
	s := NewHashSet()
	i := Int(6)
	s.Add(&i)
	s.Add(&i)
	// s.Add(Int(7))
	// s.Add(Int(54))
	// s.Add(Int(6))
	s.Each(func(e SetElement) {
		fmt.Println("e:", e)
	})
	fmt.Println("size", s.Size())
	// s.Remove(&Int(6))
	s.Each(func(e SetElement) {
		fmt.Println("i:", e)
	})
	a := []int{1, 2, 3}
	fmt.Println(a[1:len(a)])
}
