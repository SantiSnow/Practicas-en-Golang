package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

func main() {
	fmt.Println("El valor de PI según la API de Go es: ", math.Pi)

	file, err := os.Open("texto.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println(scanner.Bytes())
	}

	creacionArchivos()
}

func creacionArchivos() {
	saludo := []byte("Hola, esto fue creado con Go")

	err := ioutil.WriteFile("test.txt", saludo, 0644)

	if err != nil {
		panic(err)
	}

	//leer el archivo
	datos, err := ioutil.ReadFile("test.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("El archivo contiene: " + string(datos))
}
