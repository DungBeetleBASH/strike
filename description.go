package strike

import (
    "fmt"
    "encoding/base64"
    "errors"
)

/*
Description returns a map containing a base64 encoded HTML description, 
or a map containing a plain text description if true is passed as the 
second parameter. 
*/
func Description(params ...interface{}) (result TorrentDescription, err error) {
	l := len(params)
	if (l == 0) {
		return result, errors.New("expecting at least one parameter")
	}
	hash, ok := params[0].(string)
	if (!ok) {
		return result, errors.New("expecting 1st parameter to be of type string")
	}
	decode := false
	if (len(params) > 1) {
		decode, ok = params[1].(bool)
		if (!ok) {
			return result, errors.New("expecting 2nd optional parameter to be of type bool")
		}
	}
	query := fmt.Sprintf(api[version]["Description"], hash)
	err = callAPI(query, &result)
	if (err != nil) {
		return result, err
	}
	if (decode) {
		byteArray, err := base64.StdEncoding.DecodeString(result.Description)
		if (err != nil) {
			return result, err
		}
		result.Description = string(byteArray[:])
		return result, err
	}
	return result, err
}