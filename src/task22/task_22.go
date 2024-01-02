package main

import (
	"fmt"
	"math/big"
)

type Data struct {
	a *big.Int
	b *big.Int
}

func (d Data) Sum() *big.Int {
	return new(big.Int).Add(d.a, d.b)
}

func (d Data) Sub() *big.Int {
	return new(big.Int).Sub(d.a, d.b)
}

func (d Data) Mul() *big.Int {
	return new(big.Int).Mul(d.a, d.b)
}

func (d Data) Div() *big.Int {
	if d.b.Cmp(big.NewInt(0)) == 0 {
		fmt.Println("Error: division by zero")
		return nil
	}
	return new(big.Int).Div(d.a, d.b)
}

func main() {
	data := Data{
		a: big.NewInt(10000000),
		b: big.NewInt(20000001),
	}

	min := new(big.Int).Exp(big.NewInt(2), big.NewInt(20), nil)
	if data.a.Cmp(min) <= 0 || data.b.Cmp(min) <= 0 {
		fmt.Println("Error: a and b must be greater than 2^20")
		return
	}

	fmt.Println("Sum:", data.Sum())
	fmt.Println("Difference:", data.Sub())
	fmt.Println("Product:", data.Mul())
	if data.Div() != nil {
		fmt.Println("Quotient:", data.Div())
	}
}
