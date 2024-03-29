package main

import "fmt"

// Abstract Interface
type iAlpha interface {
	work()
	common()
}

// Abstract Concrete Type
type alpha struct {
	name string
	work func()
}

func (a *alpha) common() {
	fmt.Println("common called")
	a.work()
}

// Implementing Type
type beta struct {
	alpha
}

func (b *beta) work() {
	fmt.Println("work called")
	fmt.Printf("name is %s\n", b.name)
}

func main() {
	a := alpha{
		name: "test",
	}
	b := &beta{
		alpha: a,
	}
	b.alpha.work = b.work
	b.common()
}

/* Output:
common called
work called
name is test
*/
