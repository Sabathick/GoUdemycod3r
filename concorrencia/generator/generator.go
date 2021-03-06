package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//padrões de concorrência em go Google I/O 2012
func titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.youtube.com/", "https://github.com/")
	t2 := titulo("https://www.amazon.com.br/", "https://www.google.com/")
	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
