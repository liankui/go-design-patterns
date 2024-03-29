package main

import (
	"container/list"
	"fmt"
)

// https://golangbyexample.com/stack-in-golang/

type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value string) {
	c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		element := c.stack.Front()
		c.stack.Remove(element)
	}
	return fmt.Errorf("pop error: Stack is empty")
}

// Front 相当于查看栈顶元素（Peek）
func (c *customStack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peek error: Stack Datatype is incorrent")
	}
	return "", fmt.Errorf("peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}

func main() {
	cs := &customStack{stack: list.New()}
	fmt.Printf("push: %s\n", "A")
	cs.Push("A")
	fmt.Printf("push: %s\n", "B")
	cs.Push("B")
	fmt.Printf("size: %d\n", cs.Size())
	for cs.Size() > 0 {
		front, _ := cs.Front()
		fmt.Printf("front: %s and pop: %s\n", front, front)
		cs.Pop()
		fmt.Printf("size: %d\n", cs.Size())
	}
}

/* Output：
push: A
push: B
size: 2
front: B and pop: B
size: 1
front: A and pop: A
size: 0

*/
