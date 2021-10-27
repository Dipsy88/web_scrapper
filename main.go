package main

import (
	"golang-web-scrapper/curl"
	"golang-web-scrapper/email"
	"golang-web-scrapper/website"
	"strconv"
	"strings"
	"time"
)

func main() {
	
	for range time.Tick(time.Minute*10){
		items := curl.Anyvacancies()
		if (len(items)>0){
			msg := "Available booking: " + strings.Join(items,", ")
			email.Send(msg, msg)
		}
	}

	for range time.Tick(time.Minute*10){
		//		items := curl.Anyvacancies()
				items := curl.Specificvacancies()
				if (len(items)>0){
					msg := "Available booking specific dates: " + strings.Join(items,", ")
					email.Send(msg, msg)
				}
			}

}

func SendWebsite(){
	itemsMap := website.ScrapAllItems()
	var str strings.Builder
	for key, value := range itemsMap {
		str.WriteString(key + " found " + strconv.Itoa(len(value)) + " items " + strings.Join(value, ""))
		//fmt.Print(key + " found " + strconv.Itoa(len(value)) + " items " + strings.Join(value, ""))
	}
	//fmt.Println(str.String())
	email.Send(str.String(), "items")
}
