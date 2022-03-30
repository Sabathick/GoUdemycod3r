package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	// * é utilizado para referenciar
	//Ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n)
	fmt.Println(n)

	// & é utilizado para obter o endereço da variável
	inc2(&n)
	fmt.Println(n)
}
