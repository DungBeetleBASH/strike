package strike

import (
    "fmt"
)

/*
Imdb returns a map containing a IMDB movie information.
*/
func Imdb(imdbID string) (result IMDb, err error) {
	query := fmt.Sprintf(api[version]["Imdb"], imdbID)
	err = callAPI(query, &result)
	return result, err
}