package curl

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

const (
	testTromsoURL    string = "https://forerett-adapter.atlas.vegvesen.no/provetimer?v=2&arbeidsflytId=689843742&klasse=B&trafikkstasjonId=561"

	honefossURL        = "https://forerett-adapter.atlas.vegvesen.no/provetimer?v=2&arbeidsflytId=689843742&klasse=B&trafikkstasjonId=141"
	drobakURL       = "https://forerett-adapter.atlas.vegvesen.no/provetimer?v=2&arbeidsflytId=689843742&klasse=B&trafikkstasjonId=051"
	kongsvingerURL          = "https://forerett-adapter.atlas.vegvesen.no/provetimer?v=2&arbeidsflytId=689843742&klasse=B&trafikkstasjonId=111"

	authentication                = "C"
)

// Anyvacancies should return any dates
func Anyvacancies() []string {
	return find()
}

// Anyvacancies should return if specific dates is true
func Specificvacancies() []string {
//	var strSlice = []string{"India", "Canada", "Japan"}
	var dates = []string{"2021-11-10", "2021-11-11", "2021-11-12", "2021-11-13", "2021-11-14", "2021-11-15", "2021-11-16", "2021-11-17", "2021-11-18", "2021-11-19", "2021-11-20", "2021-11-21", "2021-11-22", "2021-11-23", "2021-11-24", "2021-11-25", "2021-11-26", "2021-11-27", "2021-11-28", "2021-11-29", "2021-11-30", "2021-12-01", "2021-12-02", "2021-12-03", "2021-12-04", "2021-12-05", "2021-12-06", "2021-12-07", "2021-12-08", "2021-12-09", "2021-12-10", "2021-12-11", "2021-12-12", "2021-12-13", "2021-12-14", "2021-12-15"}
	return findDates(dates)
}

func find() []string {
	// curl := exec.Command("curl", testTromsoURL, "-H",authentication)
	// fmt.Println(curl)
	items := []string{}

	honeyfossOut, err := exec.Command("curl", honefossURL, "-H", authentication).Output()
	if (err!=nil){
		log.Fatalf("error executing command: %v", err)
	}
	honeyfossResult := string(honeyfossOut)
	if (strings.Contains(honeyfossResult, "Unauthorized")){
		log.Fatalf("Renew your authentication")
	} else {
		if (honeyfossResult!="[]"){
			items = append(items, "honeyfoss")
		}
	}

	drobakOut, _ := exec.Command("curl", drobakURL, "-H", authentication).Output()
	drobakResult := string(drobakOut)
	if (drobakResult!="[]"){
		items = append(items, "drobak")
	}

	kongsvingerOut, _ := exec.Command("curl", kongsvingerURL, "-H", authentication).Output()
	kongsvingerResult := string(kongsvingerOut)
	if (kongsvingerResult!="[]"){
		items = append(items, "kongsvinger")
	}

	// testTromsoOut, _ := exec.Command("curl", testTromsoURL, "-H", authentication).Output()
	// testTromsoResult := string(testTromsoOut)
	// if (testTromsoResult!="[]"){
	// 	items = append(items, "testTromso")
	// }
	fmt.Println(items)
	return items
}

func findDates(dates []string) []string {
	// curl := exec.Command("curl", testTromsoURL, "-H",authentication)
	// fmt.Println(curl)
	items := []string{}

	honeyfossOut, err := exec.Command("curl", honefossURL, "-H", authentication).Output()
	if (err!=nil){
		log.Fatalf("error executing command: %v", err)
	}
	honeyfossResult := string(honeyfossOut)
	if (strings.Contains(honeyfossResult, "Unauthorized")){
		log.Fatalf("Renew your authentication")
	} else {
		if (honeyfossResult!="[]"){
			exist := checkItem(dates, honeyfossResult)
			if (exist){
				items = append(items, "honeyfoss")
			}
	
		}
	}

	drobakOut, _ := exec.Command("curl", drobakURL, "-H", authentication).Output()
	drobakResult := string(drobakOut)
	if (drobakResult!="[]"){
		exist := checkItem(dates, drobakResult)
			if (exist){
				items = append(items, "drobak")
			}
	}

	kongsvingerOut, _ := exec.Command("curl", kongsvingerURL, "-H", authentication).Output()
	kongsvingerResult := string(kongsvingerOut)
	if (kongsvingerResult!="[]"){
		exist := checkItem(dates, kongsvingerResult)
		if (exist){
			items = append(items, "kongsvinger")
		}
	}

	// testTromsoOut, _ := exec.Command("curl", testTromsoURL, "-H", authentication).Output()
	// testTromsoResult := string(testTromsoOut)
	// if (testTromsoResult!="[]"){
	// 	exist := checkItem(dates, testTromsoResult)
	// 		if (exist){
	// 			items = append(items, "testTromso")
	// 		}
	// }

	fmt.Println(items)
	return items
}


func checkItem(dates []string, result string) bool{
	exist := false
	for _, val := range dates{
		if (strings.Contains(result, val)){
			exist = true
			break
		}
	}
	return exist
}