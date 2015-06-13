package strike

func Search(params ...string) (result interface{}, err error) {
    args := []string{"","",""}
    copy(args, params)
	return callApi(getQuery("Search", args...))
}