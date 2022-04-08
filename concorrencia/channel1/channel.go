package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //enviando dados para um canal
	ch <- 2

	fmt.Println(<-ch)
}
