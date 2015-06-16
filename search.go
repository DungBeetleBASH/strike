package strike

import (
	"fmt"
)
/*
	Accepts 1-3 parameters of type string
	param 1 - the search query
	param 2 - an optional category
	param 3 - an optional subcategory
 */
func Search(params ...string) (result map[string]interface{}, err error) {
	last := len(params) - 1
	args := make([]interface{}, 3)
  	for i, _ := range args {
		if (i > last) {
			args[i] = ""
		} else {
	  		args[i] = params[i]
		}
  	}
  	query := fmt.Sprintf(api[version]["Search"], args...)
	return callApi(query)
}