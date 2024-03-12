package main

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	arrayStack := new(ArrayStack)
	arrayStack.Push("1.cat")
	arrayStack.Push("2.dog")
	arrayStack.Push("3.hen")
	fmt.Println("size:", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size:", arrayStack.Size())
	arrayStack.Push("4.drag")
	fmt.Println("pop:", arrayStack.Pop())
}
