package strike

import (
    "fmt"
)

func Description(hash string) (result interface{}, err error) {
	query := fmt.Sprintf(api[version]["Description"], hash)
	return callApi(query)
}