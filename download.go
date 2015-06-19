package strike

import (
    "fmt"
)

/*
Download returns a map containing a .torrent file url. 
*/
func Download(hash string) (result TorrentDownload, err error) {
	query := fmt.Sprintf(api[version]["Download"], hash)
	err = callAPI(query, &result)
	return result, err
}