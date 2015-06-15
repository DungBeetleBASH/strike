package strike

import (
    "fmt"
    "encoding/base64"
    "errors"
)

func Description(params ...interface{}) (result map[string]interface{}, err error) {
	l := len(params)
	if (l == 0) {
		return make(map[string]interface{}, 0), errors.New("Expecting at least one parameter")
	}
	hash, ok := params[0].(string)
	if (!ok) {
		return make(map[string]interface{}, 0), errors.New("Expecting 1st parameter to be of type string")
	}
	decode := false
	if (len(params) > 1) {
		decode, ok = params[1].(bool)
		if (!ok) {
			return make(map[string]interface{}, 0), errors.New("Expecting 2nd optional parameter to be of type bool")
		}
	}
	query := fmt.Sprintf(api[version]["Description"], hash)
	result, err = callApi(query)
	if (err != nil) {
		return result, err
	}
	if (decode) {
		byteArray, err := base64.StdEncoding.DecodeString(result["message"].(string))
		if (err != nil) {
			return result, err
		}
		result["message"] = string(byteArray[:])
		return result, err
	}
	return result, err
}