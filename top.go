package strike

import (
	"fmt"
)

/*
Search returns the top 100 torrents in a given category and (optional) subcategory.
category string
subcategory string - optional
*/
func Top(params ...string) (result map[string]interface{}, err error) {
	last := len(params) - 1
	args := make([]interface{}, 2)
  	for i, _ := range args {
		if (i > last) {
			args[i] = ""
		} else {
	  		args[i] = params[i]
		}
  	}
  	query := fmt.Sprintf(api[version]["Top"], args...)
	return callApi(query)
}