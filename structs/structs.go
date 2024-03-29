package structs

import (
	"fmt"
	"math"
)

type Curso struct {
	Temario string
	Tiempo int
}
//Método de la estructura Curso, solo se puede iniciar con una instancia de la estructura Curso y recibe un string
func (c Curso) Suscribirse(nombre string)  {
	fmt.Printf("La persona %s se ha registrado en el curso %s",nombre,c.Temario)
}

type Carrera struct {
	Nombre string
	Cursos []Curso
}

//Método de la estructura Carrara, solo se puede iniciar con una instancia de la estructura Carrera y recibe un string
func (c Carrera) Suscribirse(nombre string)  {
	fmt.Printf("La persona %s se ha registrado en la carrera %s",nombre,c.Nombre)
}

// Estructura Rectangulo
type Rectangulo struct {
	Base float64
	Altura float64
}

// Estructura Trapecio
type Trapecio struct {
	BaseMayor float64
	BaseMenor float64
	Altura float64
}

type Circulo struct {
	Radio float64
}

// Método de la estructura Rectangulo
func (figura Rectangulo) Area() float64  {
	return figura.Base * figura.Altura
}

// Método de la estructura Trapecio
func (figura Trapecio) Area() float64  {
	return (figura.BaseMayor + figura.BaseMenor) * figura.Altura / 2
}

func (figura Rectangulo) Perimetro() float64  {
	return 2 * (figura.Base + figura.Altura)
}

func (figura Trapecio) Perimetro() float64  {
	return figura.BaseMenor + figura.BaseMayor + figura.Altura
}

// Método de la estructura Circulo
func (figura Circulo) Area() float64  {
	return math.Pi * figura.Radio * figura.Radio
}

// Método de la estructura Circulo
func (figura Circulo) Perimetro() float64  {
	return 2 * math.Pi * figura.Radio
}