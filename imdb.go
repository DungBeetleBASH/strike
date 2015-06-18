package strike

import (
    "fmt"
)

/*
Imdb returns a map containing a IMDB movie information.
*/
func Imdb(imdbID string) (result map[string]interface{}, err error) {
	query := fmt.Sprintf(api[version]["Imdb"], imdbID)
	return callApi(query)
}