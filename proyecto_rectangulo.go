package main

import "fmt"

func main() {
	var longitud int
	var anchura int

	fmt.Println("Ingrese la longitud")
	fmt.Scan(&longitud)
	fmt.Println("Ingrese la anchura")
	fmt.Scan(&anchura)

	var area int = longitud * anchura
	var perimetro int = 2 * (longitud + anchura)

	fmt.Printf("El area del recangulo es %d \n", area)
	fmt.Printf("El perimetro del recangulo es %d", perimetro)
}