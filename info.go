package strike

import (
    "strings"
    "fmt"
    "errors"
)
/**
 * @method Info
 * Accepts either one []string parameter or any number of string parameters
 */
func Info(params ...interface{}) (result map[string]interface{}, err error) {
	var args []string
	l := len(params)
	switch l {
	case 0:
		return make(map[string]interface{}, 0), errors.New("Expecting at least one parameter")
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
		  	return make(map[string]interface{}, 0), errors.New("Expecting a single parameter to be of type string or []string")
		}
	default:
		args = make([]string, l)
	  	for i, v := range params {
	  		if (fmt.Sprintf("%T", v) == "string") {
	  			args[i] = v.(string)
  			} else {
  				return make(map[string]interface{}, 0), errors.New("Expecting multiple parameters to be of type string")
  			}
			
	  	}
	}
	if (len(args) == 0) {
		return make(map[string]interface{}, 0), errors.New("Unexpected error")
	}
	query := fmt.Sprintf(api[version]["Info"], strings.Join(args, ","))
	return callApi(query)
}