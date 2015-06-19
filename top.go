package strike

import (
	"fmt"
	"errors"
)

/*
Top returns the top 100 torrents in a given category and (optional) subcategory.
category string
subcategory string - optional
*/
func Top(params ...string) (results Results, err error) {
	if (len(params) == 0) {
		return results, errors.New("expecting at least one parameter")
	}
	last := len(params) - 1
	args := make([]interface{}, 2)
  	for i := range args {
		if (i > last) {
			args[i] = ""
		} else {
	  		args[i] = params[i]
		}
  	}
  	query := fmt.Sprintf(api[version]["Top"], args...)
	err = callAPI(query, &results)
    return results, err
}