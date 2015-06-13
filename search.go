package strike

import (
    "fmt"
)

func Search(params ...string) (result interface{}, err error) {
    args := []string{"","",""}
    copy(args, params)
  	fArgs := make([]interface{}, len(args))
  	for i, v := range args {
  	    fArgs[i] = v
  	}
  	query := fmt.Sprintf(api[version]["Search"], fArgs...)
	return callApi(query)
}