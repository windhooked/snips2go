package main

/*
You are given two linked-lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	l.PushBack(2)
	l.PushBack(4)
	l.PushBack(3)

	m := list.New()
	m.PushBack(5)
	m.PushBack(6)
	m.PushBack(4)

	f := m.Back()
	r := list.New()
	carry := 0

	// Iterate through list and print its contents.
	for e := l.Back(); e != nil; e = e.Prev() {
		z := (e.Value.(int) + f.Value.(int)) + carry
		f = f.Prev()
		if z < 10 {
			r.PushFront(z)
			carry = 0
		} else {
			r.PushFront(0)
			carry = 1
		}
	}
	for e := r.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%d\n", (e.Value.(int)))
	}

}
