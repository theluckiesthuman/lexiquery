package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/theluckiesthuman/lexiquery/internal/models"
)

type DictionaryQuery interface {
	Query(word string) (string, error)
}

type dictionaryQuery struct {
	url    string
	apiKey string
}

func NewDictionaryQuery(url, apiKey string) DictionaryQuery {
	return dictionaryQuery{url: url, apiKey: apiKey}
}

func (d dictionaryQuery) Query(word string) (string, error) {
	url := fmt.Sprintf("%s/%s?key=%s", d.url, word, d.apiKey)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	var definition models.WordDefinition
	if err := json.NewDecoder(response.Body).Decode(&definition); err != nil {
		return "", err
	}

	if len(definition) == 0 {
		return "", fmt.Errorf("definition not found for %s", word)
	}
	//`ek-sər-sīz` (noun): the act of bringing into play or realizing in action
	formattedDefinition := definition[0].Hwi.Prs[0].Mw + " (" + definition[0].Fl + "): " + definition[0].Shortdef[0]
	return formattedDefinition, nil
}
