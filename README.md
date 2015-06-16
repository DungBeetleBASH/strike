#strike

    strike is a Go wrapper for the getstrike.net torrent search API: https://getstrike.net/api/ 
    All of the API's methods are covered and all calls return map[string]interface{}, error

    Variadic functions have been used for some API calls to make the interface more flexible. 
    These are documented below.

[![GoDoc](https://godoc.org/github.com/Pappa/strike?status.svg)](https://godoc.org/github.com/Pappa/strike)

##API

_func Search(params ...string) (result map[string]interface{}, err error)_
	Accepts 1-3 parameters of type string
	param 1 - the search query 
	param 2 - an optional category
	param 3 - an optional subcategory

_func Description(params ...interface{}) (result map[string]interface{}, err error)_

_func Info(params ...interface{}) (result map[string]interface{}, err error)_
	Accepts either one []string parameter or any number of string parameters
	param 1 - string - an info hash 

_func Download(hash string) (result map[string]interface{}, err error)_
	hash - an info hash

_func Count() (result map[string]interface{}, err error)_
	Returns a map containing the current total torrent count on getstrike.net

_func Top(params ...string) (result map[string]interface{}, err error)_
	Accepts 1-2 parameters of type string 
	param 1 - the category 
	param 2 - an optional subcategory

_func SetVersion(v string) bool_
	Sets the getstrike.net API version to be used. Cuttently only accepts "v2"


