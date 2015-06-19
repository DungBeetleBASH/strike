package strike

import (
    "fmt"
)

/*
Imdb returns IMDB movie information.
*/
func Imdb(imdbID string) (result IMDb, err error) {
	query := fmt.Sprintf(api[version]["Imdb"], imdbID)
	err = callAPI(query, &result)
	return result, err
}