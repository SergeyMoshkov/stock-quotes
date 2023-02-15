package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

//  downloadStockWebpageBody загружает страницу по указанному URL
func downloadStockWebpageBody(symbol string) io.ReadCloser {
	res, err := http.Get("https://finance.yahoo.com/quote/" + symbol)
	if err != nil {
		log.Fatalln(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code %d: %s", res.StatusCode, res.Status)
	}

	return res.Body
}

const (
	symbol = "TSLA"
)

// main точка входа в программу
func main() {
	body := downloadStockWebpageBody(symbol)
	defer body.Close()

	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatalln(err)
	}

	getTextByField := func(field string) string {
		s := "[data-field=" + field + "][data-symbol=" + symbol + "]"

		return doc.Find(s).Text()
	}

	marketPrice := getTextByField("regularMarketPrice")
	marketChange := getTextByField("regularMarketChange")
	marketChangePercent := getTextByField("regularMarketChangePercent")

	if marketPrice == "" || marketChange == "" || marketChangePercent == "" {
		log.Fatalln("cannot access market price")
	}

	fmt.Printf(
		"%s\n%s\n\t %s\n\t%s\n",
		symbol,
		marketPrice,
		marketChange,
		marketChangePercent,
	)
}