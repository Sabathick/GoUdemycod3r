package main

import "fmt"

func main() {
	fmt.Print("Mesma Linha")
	fmt.Println("Nova Linha")

	x := 3.141516

	fmt.Println("O valor de x é:", fmt.Sprint(x))
	fmt.Printf("O valor de x é %.2f", x)
}
