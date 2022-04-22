package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hey there! Building simple microservice.")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error %v\n", err)
			return
		}
		log.Printf("Hello %s\n", data)
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Good bye! Tested the very simple microservice.")
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error %v\n", err)
			return
		}
		log.Printf("Good bye %s\n", data)
	})
	http.ListenAndServe(":8080", nil)
}
