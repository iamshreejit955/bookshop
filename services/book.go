package services

import (
	"encoding/json"
	"fmt"

	"github.com/iamshreejit955/bookshop/models"
)

func FetchAll() []byte {
	books := []models.Book{
		{Title: "Goosebumps"},
		{Title: "Harry potter"},
		{Title: "Shreejit the great"},
	}

	jsonBooks, err := json.Marshal(books)

	if err != nil {
		fmt.Println("Somehting failed")
	}

	return jsonBooks
}
