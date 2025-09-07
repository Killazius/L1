package main

import "fmt"

type Target interface {
	Operation()
}

type Adaptee struct {
}

func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("Adaptee.AdaptedOperation()")
}

type Adapter struct {
	*Adaptee
}

func (adapter *Adapter) Operation() {
	adapter.AdaptedOperation()
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func main() {
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}
