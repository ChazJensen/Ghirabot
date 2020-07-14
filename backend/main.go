package main

import (
	"fmt"
	"regexp"
//	"html/template"
	"io/ioutil"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}

func search(w http.ResponseWriter, r *http.Request) {
	// for use later: this is a []byte storing the card to look up
	cardName := r.URL.Path[len("/lookup/"):]


	resp, err := http.Get("http://api.scryfall.com/cards/named?fuzzy=" + cardName)
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	print(string(body))

	re := regexp.MustCompile(",")
	fmt.Fprintf(w, "%s", re.ReplaceAllString(string(body), ",\n"))
}

func serveBotPage(w http.ResponseWriter, r *http.Request) {

	mainPage, err := ioutil.ReadFile("../index.htm")
	if err != nil {
		print(err)
	}

	fmt.Fprintf(w, string(mainPage))/**/

//	template.ExecuteTemplate(w, "index.htm")

}

func main() {
	http.HandleFunc("/ping", ping)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))

	http.HandleFunc("/lookup/", search)
	http.HandleFunc("/play/", serveBotPage)
	http.ListenAndServe(":8080", nil)
}
