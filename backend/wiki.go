package main

import (
	"fmt"
	"regexp"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title	string
	Body	[]byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}
/*
func hello(w http.ResponseWriter, r *http.Request) {
	ioutil.WriteString(w, "Hello, World!\n");
}
/**/

func viewPage(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage(r.URL.Path[len("/view/"):])

	if err != nil {
		return
	}


	fmt.Fprintf(w, "<h1>Ghirabot</h1><hr><h3>%s</h3><div>%s</div>", p.Title, p.Body)
}

func main() {
	p := Page{Title: "testPage", Body: []byte("Hello, World!")}
	p.save()
	// http.HandleFunc("/", hello);
	http.HandleFunc("/view/", viewPage)
	http.HandleFunc("/lookup/", search)
	http.HandleFunc("/play/", serveBotPage)
	http.ListenAndServe(":8080", nil)
}
