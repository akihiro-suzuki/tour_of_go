package main

import (
	"fmt"
	"tourofgo/tree"
)

type Stack struct {
	elements []*tree.Tree
}

func (s *Stack) Push(t *tree.Tree) {
	s.elements = append(s.elements, t)
}
func (s *Stack) Pop() *tree.Tree {
	s_len := len(s.elements)
	if s_len == 0 {
		return nil
	}
	last := s.elements[s_len-1]
	s.elements = s.elements[:s_len-1]
	return last
}

func (s *Stack) Empty() bool {
	return len(s.elements) == 0
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	stack := Stack{elements: make([]*tree.Tree, 0)}
	stack.Push(t)
	for {
		ele := stack.Pop()
		if ele == nil {
			break
		}
		ch <- ele.Value
		if ele.Left != nil {
			stack.Push(ele.Left)
		}
		if ele.Right != nil {
			stack.Push(ele.Right)
		}
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	// q: how to await in go?
	// a: use channel
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
