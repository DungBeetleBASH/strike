package strike

import (
    "strings"
    "fmt"
    "errors"
)

/*
Info returns a map containing torrent information
*/
func Info(params ...interface{}) (results Results, err error) {
	var args []string
	l := len(params)
	switch l {
	case 0:
		return results, errors.New("expecting at least one parameter")
	case 1:
		switch params[0].(type) {
		case string:
			args = []string{ params[0].(string)}
		case []string:
			args = make([]string, len(params[0].([]string)))
		  	for i, v := range params[0].([]string) {
				args[i] = v
		  	}
		default:
		  	return results, errors.New("expecting a single parameter to be of type string or []string")
		}
	default:
		args = make([]string, l)
	  	for i, v := range params {
	  		if (fmt.Sprintf("%T", v) == "string") {
	  			args[i] = v.(string)
  			} else {
  				return results, errors.New("expecting multiple parameters to be of type string")
  			}
			
	  	}
	}
	if (len(args) == 0) {
		return results, errors.New("unexpected error")
	}
	query := fmt.Sprintf(api[version]["Info"], strings.Join(args, ","))
	err = callAPI(query, &results)
	return results, err
}