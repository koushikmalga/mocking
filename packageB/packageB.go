package packageB

import (
	"fmt"
	"mocking/packageC"
)

// B Interface having three functions
type BIfc interface {
	Sum(int, int) int
	UseC(int, int) int
	Prod(int, int) int
}

type B struct {
	c packageC.CIfc
}

func NewB(c packageC.CIfc) *B {
	return &B{c: c}
}

func (b *B) Sum(ele1, ele2 int) int {

	return ele1 + ele2
}

func (b *B) UseC(ele1, ele2 int) int {
	// This function calls its dependency packageC
	return b.c.Sub(ele1, ele2)
}

func (b *B) Prod(ele1, ele2 int) int {
	fmt.Println("hola")
	return ele1 * ele2
}
