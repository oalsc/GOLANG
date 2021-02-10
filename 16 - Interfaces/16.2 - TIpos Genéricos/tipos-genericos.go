package main

import "fmt"

func generica(i interface{}) {
	fmt.Println(i)
}

func main() {
	generica("String")
	generica(123)
	generica(true)

	mapa := map[interface{}]interface{}{
		uint8(3):    "String",
		float64(12): false,
		false:       123,
	}

	fmt.Println(mapa)
}
