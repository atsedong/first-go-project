package main

import (
	"html/template"
	"net/http"
)

type pageVariables struct {
	Special   string
	BestOfDay string
}

func main() {
	http.HandleFunc("/homepage", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("index.html"))

	// var homePageVariables pageVariables
	// homePageVariables.special = "Cherry blossom and strawberry cake"
	// homePageVariables.bestOfDay = "Lemon yogurt tart"

	homePageVariables := pageVariables{
		Special:   "Cherry blossom and strawberry cake",
		BestOfDay: "Lemon yogurt tart",
	}

	err := tmp.Execute(w, homePageVariables)
	if err != nil {
		println("error occurred: ", err)
	}
}
