package main

import (
	/*
		"cursogo/name"
		numeros2 "cursogo/numeros"
	*/
	"cursogo/maps"
	"fmt"
	"strings"
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

type PlatziCourse struct {
	Name string
	Slug string
	Skills [] string
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

func main() {
	/*
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
	fmt.Println(telefono)

	name.GetNombre()

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

	a, b, c := name.GetVariables()
	fmt.Println(a,b,c)

	//Calculadora
	var num1, num2 int
	fmt.Println("Ingrese primer número: ")
	fmt.Scanf("%d",&num1)
	fmt.Println("Ingrese segundo número: ")
	fmt.Scanf("%d",&num2)
	fmt.Println(numeros2.Sumar(num1,num2))
	fmt.Println(numeros2.Restar(num1,num2))
	fmt.Println(numeros2.Multiplicar(num1,num2))
	fmt.Println(string("Hola" [0]))

	name.GetArray()
	name.GetSlice()
	numeros2.GetParoImpar()
	fortest()
	string2()
	*/
	fmt.Println(maps.GetMap())
	fmt.Println(maps.GetMapaEdades())
	fmt.Println(maps.GetEdad("Diego"))

	pc := PlatziCourse{
		Name:   "Go",
		Slug:   "go",
		Skills: []string{"1","2"},
	}
	fmt.Println(pc)
	pc1 := new(PlatziCourse)
	pc1.Name = "Node"
	pc1.Slug = "node"
	pc1.Skills = []string{"backend","js"}
	fmt.Println(pc1)
}



func fortest()  {
	for i := 0; i<100; i++  {
		fmt.Println("valor de i: ",i)
	}

	c := 100
	for c > 0  {
		c -= 10
		fmt.Println("valor de c",c)
	}

	s := 1000
	for{
		s -= 1
		if s==0 {
			fmt.Println("Termina el for infinito")
			break
		}
	}
}
// métodos de string
func string2()  {
	var hello string = "hola"
	fmt.Println(strings.ToUpper(hello))
	var reemplazar string = "hello 1, hello 2, hello 3"
	fmt.Println(strings.Replace(reemplazar,"hello","hola",-1))
	fmt.Println(strings.Split(reemplazar," "))
}
