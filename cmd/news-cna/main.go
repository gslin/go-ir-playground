package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.cna.com.tw/list/aall.aspx"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("res.StatusCode: %v", resp.StatusCode)
	}

	// #jsMainList li
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("#jsMainList li").Each(func (i int, li *goquery.Selection) {
		href, exists := li.Find("a").Attr("href")
		if !exists {
			return
		}

		title := li.Find("h2").Text()

		fmt.Printf("%v: %s\n", href, title)
	})
}
