package strike

import (
    "fmt"
    "encoding/base64"
)

func Description(params ...interface{}) (result map[string]interface{}, err error) {
	hash, ok := params[0].(string)
	if (!ok) {
		return result, err
	}
	decode := false
	if (len(params) > 1) {
		decode, ok = params[1].(bool)
		if (!ok) {
			return result, err
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