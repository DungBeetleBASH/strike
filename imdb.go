package strike

import (
    "fmt"
)

func Imdb(imdbId string) (result map[string]interface{}, err error) {
	query := fmt.Sprintf(api[version]["Imdb"], imdbId)
	return callApi(query)
}