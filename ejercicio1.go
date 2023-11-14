package main

import "fmt"

func main() {
	var num1 int
	// var num2 int
	fmt.Println("ingresa un numero1: ")
	fmt.Scan(&num1)
	// fmt.Println("ingresa un numero2: ")
	// fmt.Scan(&num2)

	if num1 > 0 {
		fmt.Println("es mayor a 0")
	} else if num1 < 0 {
		fmt.Println("es menor a 0")
		// 	fmt.Println("num1 es mayor a num2")
		// } else {
		// 	fmt.Println("num2 es mayor a num1")
	} else {
		fmt.Println("es igual a 0")
	}
}
