package main

import (
	"fmt"
	"net/http"
	"time"
)

func ThePrinter() {
	for {
		fmt.Println("It did a thing")
		time.Sleep(time.Duration(time.Second))
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey there\n"))
}

func main() {
	http.HandleFunc("/", Handler)
	go http.ListenAndServe(":8080", nil)
	ThePrinter()
}
