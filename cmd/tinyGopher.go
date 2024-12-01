package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	urlDatabase := make(map[string]string)
	urlId := "/" + uuid.New().String()

	urlDatabase[urlId] = "https://gobyexample.com/maps"
	fmt.Println(urlDatabase)

	http.HandleFunc(urlId, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			http.Redirect(w, r, urlDatabase[urlId], http.StatusMovedPermanently)
		}
	})
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
