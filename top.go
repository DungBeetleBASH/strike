package strike

import (
	"fmt"
)

func Top(params ...string) (result interface{}, err error) {
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