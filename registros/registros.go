/*
	parecidos a los objetos y a los arreglos,
	los registros o structs permiten almacenar
	en un mismo identificador, diferentes miembros
	que ppueden variar en su tipo

*/

package main

import "fmt"

type Persona struct {
	name      string
	last_name string
	age       int
	height    float32
	job       Trabajo
}

type Trabajo struct {
	title      string
	salary     float32
	work_hours int
}

func main() {
	var mi_persona Persona
	var mi_trabajo Trabajo

	mi_trabajo.title = "Desarrollador"
	mi_trabajo.salary = 65000.00
	mi_trabajo.work_hours = 8

	mi_persona.name = "Santiago"
	mi_persona.last_name = "Aguirre"
	mi_persona.age = 25
	mi_persona.height = 1.70
	mi_persona.job = mi_trabajo

	fmt.Println("Hola, me llamo", mi_persona.name)
	fmt.Println("Mi edad es", mi_persona.age)
	fmt.Println("Trabajo de", mi_persona.job.title)
	fmt.Println("En horas, trabajo por día unas", mi_persona.job.work_hours)
}
