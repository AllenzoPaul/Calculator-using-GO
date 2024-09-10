package main

import (
	"fmt"

	"example.com/arithmetic"
)

func main() {
	var a, b float64
	var s int8
	var err error
	var result float64

	fmt.Print("Enter the first number : ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number : ")
	fmt.Scan(&b)
	fmt.Print("\nSELECT THE OPERATOR\nAdd : 1\nSUB : 2\nMUL : 3\nDIV : 4\nMOD : 5")
	fmt.Scan(&s)

	if s == 1 {
		fmt.Println("Result : ", arithmetic.Add(a, b))
	} else if s == 2 {
		fmt.Println("Result : ", arithmetic.Sub(a, b))
	} else if s == 3 {
		fmt.Println("Result : ", arithmetic.Mul(a, b))
	} else if s == 4 {
		result, err = arithmetic.Div(a, b)
		if err != nil {
			fmt.Print(err.Error())
		}

		fmt.Println("Result : ", result)
	} else if s == 5 {
		fmt.Println("Result : ", arithmetic.Mod(a, b))
	}

}
