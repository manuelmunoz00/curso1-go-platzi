package name

import "fmt"

// arrays(de tamaño fijo)
func GetArray() {
	var arr1[2] string
	arr1[0] = "uno"
	arr1[1] = "dos"
	fmt.Println(arr1)
}

//array (no se declara tamaño)
func GetSlice()  {
	var slice1[] string
	slice1 = append(slice1, "1")
	slice1 = append(slice1, "2")
	fmt.Println(slice1)
}
