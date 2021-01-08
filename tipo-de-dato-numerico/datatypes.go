package main

import "fmt"

//data types en Go

func main() {

	fmt.Println("Tipos de datos en Go")
	fmt.Println("Datos numericos enteros:")

	/*uint8, 8-bits, (0 a 255)
	uint16, 16-bits (0 a 65535)
	uint32, 32-bits (0 a 4294967295)
	uint64, 64-bits (0 a 18446744073709551615)
	int8, 8-bits (-128 a 127)
	int16, 16-bits (-32768 a 32767)
	int32, 32-bits (-2147483648 a 2147483647)
	int64, 64-bits (-9223372036854775808 a 9223372036854775807)
	*/

	//definicion de variables

	var numero_uno int
	var numero_dos int32
	var numero_tres uint8
	var numero_cuatro uint32

	//definicion y asignacion

	numero_uno = 32
	numero_dos = 22
	numero_tres = 3
	numero_cuatro = 3421

	//o tambien pueden definirse y asignarse

	var numero_cinco int64 = 123123
	var numero_seis int32 = -12312
	var numero_siete uint32 = 7896

	//o incluso sin su tipo de dato

	var z = 3
	var x = 7
	var y = -14

	fmt.Println(numero_uno)
	fmt.Println(numero_dos)
	fmt.Println(numero_tres)
	fmt.Println(numero_cuatro)
	fmt.Println(numero_cinco)
	fmt.Println(numero_seis)
	fmt.Println(numero_siete)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println("Datos numericos decimales y cientificos:")

	/*
		byte          8-bits (0 a 255)
		rune          32-bits (-2147483648 a 2147483647)
		uint          32 o 64 bits
		int           32 o 64 bits
		uintptr       Entero guarda los bits no interpretados de un puntero
	*/

	var numero_decimal_uno float32 = 55.5
	var numero_decimal_dos float64 = 432.342
	var numero_decimal_tres complex64 = 3.322
	var numero_decimal_cuatro complex128 = 3.2

	fmt.Println(numero_decimal_uno)
	fmt.Println(numero_decimal_dos)
	fmt.Println(numero_decimal_tres)
	fmt.Println(numero_decimal_cuatro)

	//otros
	fmt.Println("Otros numeros: ")

	var numero_distinto byte = 2 //igual a uint8

	fmt.Println(numero_distinto)

}
