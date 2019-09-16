package structs

import (
	"fmt"
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