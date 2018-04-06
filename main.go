package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type pageVariables struct {
	Special   string
	BestOfDay string
}

func main() {
	http.HandleFunc("/", homePage)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))

	homePageVariables := pageVariables{
		Special:   "Cherry blossom and strawberry cake",
		BestOfDay: "Lemon yogurt tart",
	}

	err := tmp.Execute(w, homePageVariables)
	if err != nil {
		fmt.Println("error occurred: ", err)
	}
}
