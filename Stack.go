package main

import (
	"container/list"
	"fmt"
)

//LListLifo stack implementation on ll
func LListLifo() {
	//last in first out
	lifo := list.New()

	for i := 0; i < 10; i++ {
		lifo.PushBack(i) //push
	}
	for lifo.Back() != nil {
		fmt.Println(lifo.Back().Value) // get linked list tail value
		lifo.Remove(lifo.Back())       // remove linked list tail value, pop
	}
}

// Slice implementation

type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return ret
}

func main() {
	LListLifo()
	s := new(Stack)
        s.Push(9)
	fmt.Printf("%d\n", s.Pop()
}
