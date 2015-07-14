package strike

import (
	"fmt"
)

/*
Search returns a TorrentResults object and can be passed 1-3 parameters.
query string
category string - optional
subcategory string - optional
*/
func Search(params ...string) (results TorrentResults, err error) {
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
	err = callAPI(query, &results)
    return results, err
}