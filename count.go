package strike

func Count() (result map[string]interface{}, err error) {
	return callApi(api[version]["Count"])
}