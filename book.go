package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type BookList struct {
	Books []Book `json:"works"`
}

type Book struct {
	Title         string   `json:"title"`
	Author        []Author `json:"authors"`
	EditionNumber int32    `json:"edition_count"`
}

type Author struct {
	Name string `json:"name"`
}

func GetBookList(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://openlibrary.org/subjects/love.json?published_in=1500-1600")
	if err != nil {
		log.Fatalln(err)
	}

	bookList := BookList{}
	err = json.NewDecoder(resp.Body).Decode(&bookList)
	if err != nil {
		fmt.Println(err.Error())
	}

	result, err := json.Marshal(bookList.Books)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
