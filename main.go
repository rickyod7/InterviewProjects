package main

import (
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/booklist", GetBookList)
	http.HandleFunc("/bookschedule", BookSchedule)
	http.HandleFunc("/getallschedule", GetAllSchedule)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
