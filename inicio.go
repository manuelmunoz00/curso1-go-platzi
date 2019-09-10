package main

import (
	"fmt"
)

type Cliente struct {
	rut string
	nombres string
	apellidos string
	edad int
}

func main() {
	var persona1 = Cliente {
		rut: "16006363-7",
		nombres: "Manuel Alejandro",
		apellidos:"Muñoz Ayala",
		edad:34}
	var suma int
	suma = 8 + 9
	fmt.Println(suma)
	//time.Sleep(time.Second * 5)
	fmt.Println("Hello world")
	var nombre string = "Manuel"
	fmt.Println(nombre)
	pais := "Chile"
	fmt.Println(pais)
	fmt.Println(persona1)

	
}
