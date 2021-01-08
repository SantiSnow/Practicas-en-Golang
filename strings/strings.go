package main

import "fmt"

//strings en go

func main() {

	var mi_texto string = "Cadena de texto en Go"
	var mi_segundo_texto string = "Esta es una cadena de texto o String"

	fmt.Println(mi_texto)
	fmt.Println(mi_segundo_texto)

	fmt.Printf("Tipo de dato de la variable mi_texto: %T", mi_texto)

	/*
		Printf permite imprimir texto con cierto formato.
		%T especifica que se debe imprimir el tipo
		de dato de la variable desps de la coma
	*/

	//constantes
	fmt.Println(" ")
	fmt.Println("Constantes en Go")

	const mi_constante int = 53

	fmt.Println("Mi constante vale:", mi_constante)

	//o tambien

	const (
		x = "Estas"
		y = "Son"
		z = "Constantes"
	)

	fmt.Println(x, y, z)

}
