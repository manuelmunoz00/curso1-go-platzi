package main

import (
	/*
		"cursogo/name"
		numeros2 "cursogo/numeros"
	*/
	"cursogo/maps"
	"cursogo/structs"
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

func deferTest()  {
	fmt.Println("Función defer que se ejecuta al final")
}

func main() {
	defer deferTest()
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

	/*
	curso1 := structs.Curso{
		Temario: "Docker",
		Tiempo:  20,
	}
	curso1.Suscribirse("Manuel Muñoz")
	curso2 := structs.Curso{
		Temario: "Kubernetes",
		Tiempo:  30,
	}

	carrera1 := structs.Carrera{
		Nombre: "Devops",
		Cursos: [] structs.Curso{ curso1, curso2},
	}
	fmt.Println(carrera1)
	*/
	//structs.InterfaceTest()


	r1 := new(structs.Rectangulo)
	r1.Base = 4
	r1.Altura = 7.5
	t1 := new(structs.Trapecio)
	t1.BaseMayor = 5
	t1.BaseMenor = 2
	t1.Altura = 3
	fmt.Println(structs.ObtenerArea(r1))
	fmt.Println(structs.ObtenerArea(t1))
	fmt.Println("-------")
	fmt.Println(structs.ObtenerPerimetro(r1))
	fmt.Println(structs.ObtenerPerimetro(t1))
	fmt.Println("Circulo a continuación")
	c1 := new(structs.Circulo)
	c1.Radio = 20
	fmt.Println(structs.ObtenerArea(c1))
	fmt.Println(structs.ObtenerPerimetro(c1))
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
