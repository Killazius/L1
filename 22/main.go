package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2 << 20)
	b := big.NewInt(2 << 30)
	fmt.Printf("a = %d\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("multiply %d\n", new(big.Int).Mul(a, b))
	fmt.Printf("add %d\n", new(big.Int).Add(a, b))
	fmt.Printf("subtract %d\n", new(big.Int).Sub(b, a))
	fmt.Printf("divide %d\n", new(big.Int).Div(b, a))
}
