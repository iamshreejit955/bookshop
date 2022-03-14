package controller

import (
	"net/http"

	"github.com/iamshreejit955/bookshop/services"
)

func FetchAllBooks(w http.ResponseWriter, response *http.Request) {
	x := services.FetchAll()
	w.Write(x)
}
