package main

import (
	"net/http"

	"github.com/iamshreejit955/bookshop/controller"
)

func main() {
	http.HandleFunc("/books", controller.FetchAllBooks)
	http.ListenAndServe(":5000", nil)
}
