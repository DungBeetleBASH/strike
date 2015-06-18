package strike

import (
	"fmt"
)

/*
Search returns collection of torrents and can be passed 1-3 parameters.
query string
category string - optional
subcategory string - optional
*/
func Search(params ...string) (result map[string]interface{}, err error) {
	last := len(params) - 1
	args := make([]interface{}, 3)
  	for i := range args {
		if (i > last) {
			args[i] = ""
		} else {
	  		args[i] = params[i]
		}
  	}
  	query := fmt.Sprintf(api[version]["Search"], args...)
	return callApi(query)
}