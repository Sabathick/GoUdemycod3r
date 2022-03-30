package main

import "fmt"

func f1() {
	fmt.Println("F1")
} //função sem parâmetro e sem retorno

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
} //função com 2 parâmetros

func f3() string {
	return "F3"
} //função que declara o retorno "string", logo é obrigatório um retorno string

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {
	f1()
	f2("Parametro1", "Parametro2")

	r3, r4 := f3(), f4("Parametro1", "Parametro2")

	fmt.Println(r3, r4)

	r51, r52 := f5()
	fmt.Println("F5:", r51, r52)
}
