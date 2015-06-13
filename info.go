package strike

import (
    "strings"
    "fmt"
)

func Info(params []string) (result interface{}, err error) {
	query := fmt.Sprintf(api[version]["Info"], strings.Join(params, ","))
	return callApi(query)
}