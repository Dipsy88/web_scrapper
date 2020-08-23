package website

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	womanShoeURL    string = "https://www.xxl.no/sko/damesko/c/140400?sort=campaign&pages=50"
	womanClothesURL        = "https://www.xxl.no/klar/klar-dame/c/120200?sort=campaign&pages=50"
	manShoeURL      string = "https://www.xxl.no/sko/herresko/c/140200?sort=campaign&pages=50"
	manClothesURL          = "https://www.xxl.no/klar/klar-herre/c/120400?sort=campaign?sort=campaign&pages=50"
	bagsURL                = "https://www.xxl.no/jakt-og-friluft/ryggsekk-bag-og-tilbehor/c/201400?sort=campaign&pages=100"
)

// ScrapAllItems should return a list of items
func ScrapAllItems() map[string][]string {
	itemsMap := make(map[string][]string, 5)
	addItems(itemsMap, "woman shoes", getItems(womanShoeURL))
	addItems(itemsMap, "man shoes", getItems(manShoeURL))
	addItems(itemsMap, "woman clothes", getItems(womanClothesURL))
	addItems(itemsMap, "man clothes", getItems(manClothesURL))
	addItems(itemsMap, "bags", getItems(bagsURL))

	return itemsMap
}

func getItems(url string) []string {
	items, err := getURLContents(url)
	if err != nil {
		log.Println(err)
	}
	return items
}

func addItems(itemsMap map[string][]string, key string, values []string) {
	if len(values) > 0 {
		itemsMap[key] = values
	} else {
		fmt.Printf("There are no items for %s \n", key)
	}

}

func getURLContents(url string) ([]string, error) {
	items := []string{}
	// Get the HTML
	resp, err := http.Get(url)
	if err != nil {
		return items, err
	}
	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return items, err
	}
	// Save each .post-title as a list
	doc.Find(".js-tracking").Children().Each(func(i int, s *goquery.Selection) {
		class, _ := s.Attr("class")

		if class == "product-list__info-wrapper" {
			var text string = s.Text()
			if strings.Contains(text, "Pris for innloggede abonnenter") || strings.Contains(text, "XXL Førpris") {
				regex, err := regexp.Compile("\n\n")
				if err != nil {
					return
				}
				s := regex.ReplaceAllString(text, "")
				items = append(items, s)
			}
		}
	})
	return items, nil
}
