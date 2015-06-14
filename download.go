package strike

import (
    "fmt"
)

func Download(hash string) (result map[string]interface{}, err error) {
	fmt.Println(hash)
	query := fmt.Sprintf(api[version]["Download"], hash)
	fmt.Println(query)
	return callApi(query)
}