package strike

/*
	Returns a map containing the current total torrent count on getstrike.net
 */
func Count() (result map[string]interface{}, err error) {
	return callApi(api[version]["Count"])
}