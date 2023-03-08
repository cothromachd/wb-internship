package main

import (
	"fmt"
	"math/big"
)

func main() {	
	fmt.Println("=============BIG NUMBERS MATH OPERATIONS=============")
	var a int
	fmt.Printf("number is 2^a\ninput \"a\": ")
	fmt.Scan(&a)
	//bigInt := big.NewInt(1<<13 + 1<<54 + 1<<60)
	//bigInt_ := big.NewInt(1 << 18 + 1 << 2)
	bigInt := big.NewInt(2)
	for i:=0; i<a-1; i++ {
		bigInt = bigInt.Mul(bigInt, big.NewInt(2))
	}
	fmt.Println("first number:", bigInt)
	fmt.Println()

	var b int
	fmt.Printf("number is 2^a\ninput \"a\": ")
	fmt.Scan(&b)

	bigInt_ := big.NewInt(2)
	for i:=0; i<b-1; i++ {
		bigInt_ = bigInt_.Mul(bigInt_, big.NewInt(2))
	}
	fmt.Println("second number:", bigInt_)
	fmt.Println()
	fmt.Println("sum:", new(big.Int).Add(bigInt, bigInt_))
	fmt.Println("sub:", new(big.Int).Sub(bigInt, bigInt_))
	fmt.Println("multiplication:", new(big.Int).Mul(bigInt, bigInt_))
	fmt.Println("divide:", new(big.Float).Quo(new(big.Float).SetInt(bigInt), new(big.Float).SetInt(bigInt_)))
	fmt.Println("=====================================================")
}