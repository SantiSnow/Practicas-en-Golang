package main

import "fmt"

func main() {
	var resultado int
	resultado = sumar(15, 39)
	fmt.Println("El resultado de la suma es:", resultado)

	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println(" ")

	fmt.Println("Dados los números 1, 2 y 3")
	fmt.Println("El mínimo y el máximo son: ")

	min, max := minmax(2, 3, 4)

	fmt.Println(min)
	fmt.Println(max)
}

func sumar(num1 int, num2 int) int {
	var resultado int
	resultado = num1 + num2
	return resultado
}

//funcion para encontrar el minimo y el maximo de cualquier numero
func minmax(x, y, z int) (int, int) {
	var min int
	var max int

	if x > y && x > z {
		max = x
		if y < z {
			min = y
		} else {
			min = z
		}
	} else if y > x && y > z {
		max = y
		if x < z {
			min = x
		} else {
			min = z
		}
	} else {
		max = z
		if x < y {
			min = x
		} else {
			min = y
		}
	}
	return min, max
}
