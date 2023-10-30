package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type SiteResult map[string]bool

var r = regexp.MustCompile(`<a\shref="([^"]+)"`)

func main() {

	result := SiteResult{}

	err := visit("https://go.dev", "/", result)

	if err != nil {
		fmt.Println("Error!!!", err)
		return
	}

	for url, rel := range result {
		if rel {
			fmt.Println(url)
		}
	}
}

func visit(host string, url string, result SiteResult) error {
	resp, err := http.Get(host + url)

	if err != nil {
		fmt.Println("Error!!!", err)
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error!!!", err)
		return err
	}

	allSubMatches := r.FindAllSubmatch(body, -1)

	for _, matches := range allSubMatches {
		for _, subMatch := range matches {
			tempUrl := string(subMatch)
			if strings.HasPrefix(tempUrl, "<") {
				continue
			}
			rel := strings.HasPrefix(tempUrl, "/")
			_, ok := result[tempUrl]
			if !ok && rel {
				result[tempUrl] = rel
				fmt.Println("Visited", tempUrl, len(result))
				visit(host, tempUrl, result)
			} else {
			}

		}
	}
	return nil
}
