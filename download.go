package strike

import (
    "fmt"
)

/*
Download returns a map containing a .torrent file url. 
*/
func Download(hash string) (result map[string]interface{}, err error) {
	query := fmt.Sprintf(api[version]["Download"], hash)
	return callAPI(query)
}