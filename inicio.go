package main

import (
	"fmt"
)

const hello  string = "Hola %s"

type Cliente struct {
	rut string
	nombres string
	apellidos string
	edad int
}

type Producto struct {
	codigo string
	nombre string
	valor int
}

func getNombre()  {
	fmt.Println("funcion getNombre")
}

func getNombreyApellido() (string,string) {
	return "nombre", "apellido"
}

func getNombreyEdad() (dato1 string, dato2 int)  {
	dato1 = "Manuel"
	dato2 = 34
	return
}

func getPedidos(pedido int, moneda string) (string,int,string) {
	precio := func() int{
		return pedido * 7
	}
	return "El precio del pedido: ",precio(), moneda
}

func pedidos(caracteristicas ...string)  {
	for _, caracteristica := range caracteristicas  {
		fmt.Println(caracteristica)
	}
}

func getVariables() (int, int, int)  {
	return 1,2,3
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

	var telefono = Producto{"123","Samsung", 50000}
	//fmt.Println(telefono.nombre)
	fmt.Print(telefono)

	getNombre()

	fmt.Println(getNombreyApellido())
	fmt.Println(getNombreyEdad())
	fmt.Println(getPedidos(22,"$"))
	pedidos("pantalon","zapatillas","gorra")

	//array
	var peliculas = [3]string{
		"Batman",
		"Uno",
		"Dos"}
	fmt.Println(peliculas)

	//slices (arrays dinámicos)
	var numeros = [] string{
		"uno",
		"dos",
		"tres"}

	fmt.Println(numeros)
	numeros = append(numeros, "cuatro")
	fmt.Println(numeros [0:2])

	a, b, c := getVariables()
	fmt.Println(a,b,c)

	//Calculadora
	var num1, num2 int
	fmt.Println("Ingrese primer número: ")
	fmt.Scanf("%d",&num1)
	fmt.Println("Ingrese segundo número: ")
	fmt.Scanf("%d",&num2)
	fmt.Println(sumar(num1,num2))
	fmt.Println(restar(num1,num2))
	fmt.Println(multiplicar(num1,num2))

	fmt.Println(string("Hola" [0]))
}

func sumar(a int, b int) int  {
	return a+b
}

func restar(a int, b int) int {
	return a-b
}

func multiplicar(a int, b int)  int{
	return a*b
}
