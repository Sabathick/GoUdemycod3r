package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[123456] = "Maria"
	aprovados[4343] = "João"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	funcsESalarios := map[string]float64{
		"Jose":     1143.0,
		"Gabriela": 3400.0,
	}

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 12345.0,
			"Guga":           3344.0,
		},
		"J": {
			"John": 2324.0,
			"Jîm":  3066.0,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

}
