package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
	fmt.Println(w, r) //debug request info
}

func main() {
	fmt.Println("Listening on port 8008")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/")))) // send assets as files
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8008", nil)
}