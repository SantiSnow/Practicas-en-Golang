package main

import (
	"bufio"
	"fmt"
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
}
