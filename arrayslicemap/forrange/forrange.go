package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4}

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
		// _ é uma forma de não precisar declarar uma variável, logo, sendo obrigado a utilizá-la
	}

}
