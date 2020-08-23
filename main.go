package main

import (
	"golang-web-scrapper/email"
	"golang-web-scrapper/website"
	"strconv"
	"strings"
)

func main() {
	itemsMap := website.ScrapAllItems()
	var str strings.Builder
	for key, value := range itemsMap {
		str.WriteString(key + " found " + strconv.Itoa(len(value)) + " items " + strings.Join(value, ""))
		//fmt.Print(key + " found " + strconv.Itoa(len(value)) + " items " + strings.Join(value, ""))
	}
	//fmt.Println(str.String())
	email.Send(str.String())
}
