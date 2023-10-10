package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

var r = regexp.MustCompile(`<a\shref="([^"]+)"`)

func main() {
	fmt.Println("Hello World!!!")
	visit("https://go.dev/")
}

func visit(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error!!!", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error!!!", err)
		return
	}
	//fmt.Printf("Количество символов:%d\n%s\n", len(body), body[0:1000])
	allSubMatches := r.FindAllSubmatch(body, -1)
	for _, matches := range allSubMatches {
		for idx, subMatches := range matches {
			if idx%2 == 1 {
				fmt.Println(string(subMatches))
			}
		}
	}
}

//написать функцию, которая отделит внешнюю от внутренней и в конце пройтись по всем страницам этого сайта, собрать список этих страниц
