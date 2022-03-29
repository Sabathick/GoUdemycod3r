package main

import "fmt"

func main() {
	// homogênea == mesmo tipo de dado; tamanho fixo
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 5.6, 2.7, 8.8
	// se adicionar mais valores a uma array excedendo seu tamanho máximo, causará erro de compilação
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Println(media)
}
