package utils

import (
	"encoding/json"
	"github.com/mostafijurj/go-book-store/pkg/models"
	"io/ioutil"
	"net/http"
)

func ParseRequestBody(r *http.Request) (book models.Book, err error) {
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &book)
	if err != nil {
		return models.Book{}, err
	}
	return book, nil

}
