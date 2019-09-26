package numeros

import (
	"errors"
	"fmt"
)

// Sumar
func Sumar(a interface{}, b interface{}) (int,error)  {
	switch a.(type) {
	case string:
		return 0,errors.New("El valor de a es un string")
	}

	switch b.(type) {
	case string:
		return 0, errors.New("El valor de b es un string")
	}
	return a.(int) + b.(int),nil
}
// Restar
func Restar(a int, b int) int {
	return a-b
}
// Multiplicar
func Multiplicar(a int, b int)  int{
	return a*b
}

func GetParoImpar()   {
	var number = 0
	fmt.Println("ingresa un numero del 1 al 10: ")
	fmt.Scanf("%d",&number)
	if number % 2 == 0 {
		fmt.Println("es par")
	}else {
		fmt.Println("es impar")
	}

	//declaración y verificación en if
	if number2 := 3; number2 == 3 {
		fmt.Println("Entro al condicional")
	}
}
