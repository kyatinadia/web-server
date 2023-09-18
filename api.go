package main

import (
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"strconv"
	"text/template"
)

type Book struct {
	ID     int
	Title  string
	Stock  int
	Author string
}

var books = []Book{
	{ID: 1, Title: "Gun, Germs and Steel", Stock: 30, Author: "Jared Diamond"},
	{ID: 2, Title: "Sapiens", Stock: 50, Author: "Yoah Harari"},
	{ID: 3, Title: "Madilog", Stock: 10, Author: "Tan Malaka"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/books", getBooks)

	http.HandleFunc("/create-books", createBook)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			// log.Println("Error parsing template:", err)
			// return
		}

		tpl.Execute(w, books)
		return
	}

	// if r.Method == "GET" {
	// json.NewEncoder(w).Encode(books)
	// return
	// }

	http.Error(w, "Invalid method", http.StatusBadRequest)

}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		title := r.FormValue("title")
		stock := r.FormValue("stock")
		author := r.FormValue("author")

		convertStock, err := strconv.Atoi(stock)
		if err != nil {
			http.Error(w, "Invalid stock", http.StatusBadRequest)
		}

		newBook := Book{
			ID:     len(books) + 1,
			Title:  title,
			Stock:  convertStock,
			Author: author,
		}

		books = append(books, newBook)

		json.NewEncoder(w).Encode(books)
		return

	}

	http.Error(w, "Invalid method", http.StatusBadRequest)

}
