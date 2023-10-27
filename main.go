package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// type visitResult struct {
// 	sites map[string]siteResult
// }

type SiteResult map[string]bool

var r = regexp.MustCompile(`<a\shref="([^"]+)"`)

func main() {

	result := SiteResult{}

	err := visit("https://go.dev", "/", &result)

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

func visit(host string, url string, result *SiteResult) error {
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
			// sr := siteResult{
			// 	url:        string(subMatch),
			// 	isRelative: strings.HasPrefix(string(subMatch), "/"),
			// }
			rel := strings.HasPrefix(tempUrl, "/")
			_, ok := (*result)[tempUrl]
			if !ok && rel {
				(*result)[tempUrl] = rel
				fmt.Println("...Scan", tempUrl)
				visit(host, tempUrl, result)
				fmt.Println("...End Scan")
			} else {
				fmt.Println("Skip", tempUrl)
			}

		}
	}
	return nil
}

func findLinks(subMatches []uint8) {

}

//сделать внутри визит вызов самой себя, при этом изменить визит таким образом что бы рещультат был общим у всей цепочки визитов
//что бы визит не вызывал саму себя для тех сайтов для которых результат уже есть, так как результаты мы будем записывать в мапу
//использовать обращения к мапеи
