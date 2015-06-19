package strike

/*
Count returns the number torrents indexed by https://strikesearch.net
*/
func Count() (count uint64, err error) {
	torrentCount := TorrentCount{}
	err = callAPI(api[version]["Count"], &torrentCount)
	if (err != nil) {
		return 0, err
	}
	return torrentCount.Count, nil
}