package main

// Разработать программу, которая перемножает, делит, складывает,
//вычитает две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

func t22() {
	a := big.NewInt(11743121857)
	b := big.NewInt(1073721824)
	var tmp big.Int
	fmt.Println("a + b = ", tmp.Add(a, b))
	fmt.Println("a - b = ", tmp.Add(a, tmp.Neg(b)))
	fmt.Println("b - a = ", tmp.Add(tmp.Neg(a), b))
	fmt.Println("a * b = ", tmp.Mul(a, b))
	fmt.Println("a / b = ", tmp.Quo(a, b))
	fmt.Println("b / a = ", tmp.Quo(b, a))
}
