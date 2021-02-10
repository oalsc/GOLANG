package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posicao 1", "Posicao 2", "Posicao 3", "Posicao 4", "Posicao 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 3, 4, 5, 6, 7, 3}
	fmt.Println(slice)

	slice = append(slice, 39)
	fmt.Println(slice)

	slice2 := array2[0:3]
	fmt.Println(slice2)

	array2[1] = "Posicao alterada"
	fmt.Println(slice2)

	// Arrays internos
	fmt.Println("----------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 12)
	slice3 = append(slice3, 16)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Length
	fmt.Println(cap(slice3)) // Capacity

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 12.25)
	fmt.Println(len(slice3)) // Length
	fmt.Println(cap(slice3)) // Capacity
}
