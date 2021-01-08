package main

import "fmt"

//Operadores en go

func main() {

	fmt.Println("Operadores en Go")

	/*
		+   -   *   /    %   --    ++
	*/

	var num1 = 28
	var num2 = 14

	//operaciones matematicas
	var suma = num1 + num2
	var resta = num1 - num2
	var multi = num1 * num2
	var divi = num1 / num2

	fmt.Println("La suma dió", suma)
	fmt.Println("La resta dió", resta)
	fmt.Println("La multiplicación dio", multi)
	fmt.Println("La división dio", divi)

	//operaciones de comparacion

	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)

	var num3 = 50
	var num4 = 50
	var num5 = 60

	//operadores de conjuncion
	/*
		&& y
		|| o
		! negacion
	*/

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Operadores lógicos y de conjunción")

	fmt.Println(num3 == num4 && num4 < num5)  //true
	fmt.Println(num4 != num5 || num4 <= num3) //true
	fmt.Println(num3 == num4 && num4 > num5)  //false

}
