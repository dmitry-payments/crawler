package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type siteResult struct {
	url        string
	isRelative bool
}

type visitResult struct {
	sites []siteResult
}

var r = regexp.MustCompile(`<a\shref="([^"]+)"`)

func main() {
	visitResult, err := visit("https://go.dev/")

	if err != nil {
		fmt.Println("Error!!!", err)
		return
	}

	for _, sr := range visitResult.sites {
		if sr.isRelative {
			fmt.Println(sr.url)
		}
	}
}

func visit(url string) (*visitResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error!!!", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error!!!", err)
		return nil, err
	}
	//fmt.Printf("Количество символов:%d\n%s\n", len(body), body[0:1000])
	allSubMatches := r.FindAllSubmatch(body, -1)
	result := &visitResult{}
	for _, matches := range allSubMatches {
		for idx, subMatch := range matches {
			if idx%2 == 1 {
				result.sites = append(result.sites, siteResult{
					url:        string(subMatch),
					isRelative: strings.HasPrefix(string(subMatch), "/"),
				})
			}
		}
	}
	return result, nil
}

func findLinks(subMatches []uint8) {

}
