package main

import "fmt"

type embedded struct {
	i int
}

func (x embedded) doSameting() {
	fmt.Println("test.doSometing()")
}

type test struct {
	embedded
}

func main() {
	var x test

	x.doSameting()

	fmt.Println(x.i)
}
