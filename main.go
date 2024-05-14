package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func handleF(url, selector string) ([]string, error){
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	var texts []string

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		texts = append(texts, s.Text())
	})

	return texts, nil


}



func wikiHandler(w http.ResponseWriter, r *http.Request) {

	link := "https://animego.online/17-sharlotta-k1.html" 

	titles, err := handleF(link, "a")

	if err != nil {
		log.Println(err)
	}

	for _, title := range titles {
		fmt.Println(title)
	}


	user, _ := json.Marshal(titles)

	w.Write(user)



}

func main() {

	http.HandleFunc("/anime", wikiHandler)
	http.ListenAndServe(":8080", nil)

}