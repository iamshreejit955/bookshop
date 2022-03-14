package services

import (
	"encoding/json"
	"fmt"

	"github.com/iamshreejit955/bookshop/models"
)

func FetchAll() []byte {
	books := []models.Book{
		models.Book{Title: "Goosebumps"},
		models.Book{Title: "Harry potter"},
		models.Book{Title: "Shreejit the great"},
	}

	jsonBooks, err := json.Marshal(books)

	if err != nil {
		fmt.Println("Somehting failed")
	}

	return jsonBooks
}
