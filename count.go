package strike

func Count() (result interface{}, err error) {
	return callApi(api[version]["Count"])
}