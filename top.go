package strike

import (
	"fmt"
)
/*
	Accepts 1-2 parameters of type string
	param 1 - the category
	param 2 - an optional subcategory
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