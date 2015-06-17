package strike

import (
	"fmt"
)

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