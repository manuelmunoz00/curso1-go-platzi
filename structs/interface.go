package structs

// Interfaz de Curso y Carrera
type Platzi interface {
	Suscribirse(name string)
}

func InterfaceTest()  {
	curso3 := Curso{
		Temario: "Python",
		Tiempo:  40,
	}
	carrera2 := new(Carrera)
	carrera2.Nombre = "Dev"
	carrera2.Cursos = []Curso{curso3}
	callSuscribe(curso3)
	callSuscribe(carrera2)
}

func callSuscribe(p Platzi)  {
	p.Suscribirse("Manuel")
}

//Interfaz de Rectangulo y Trapecio
type FiguraGeometrica interface {
	Area() float64
	Perimetro() float64
}

func ObtenerArea(figura FiguraGeometrica) float64  {
	return figura.Area()
}

func ObtenerPerimetro(figura FiguraGeometrica) float64  {
	return figura.Perimetro()
}