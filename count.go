package strike

/*
Count returns the number torrents indexed by https://strikesearch.net
*/
func Count() (count int64, err error) {
	result, err := callApi(api[version]["Count"])
	if (err != nil) {
		return 0, err
	}
	count = int64(result["message"].(float64))
	return count, nil
}