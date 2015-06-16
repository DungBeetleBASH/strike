package strike

import (
    "fmt"
)

func Download(hash string) (result map[string]interface{}, err error) {
	query := fmt.Sprintf(api[version]["Download"], hash)
	return callApi(query)
}