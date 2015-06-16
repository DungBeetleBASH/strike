PACKAGE DOCUMENTATION

package strike
    import "github.com/Pappa/strike"

    #strike

    strike is a Go wrapper for the getstrike.net torrent search API:
    https://getstrike.net/api/ All of the API's methods are covered and all
    calls return map[string]interface{}, error

    Variadic functions have been used for some API calls to make the
    interface more flexible. These are documented below.

    [![GoDoc](https://godoc.org/github.com/Pappa/strike?status.svg)](https://godoc.org/github.com/Pappa/strike)

FUNCTIONS

func Count() (result map[string]interface{}, err error)
    Returns a map containing the current total torrent count on
    getstrike.net

func Description(params ...interface{}) (result map[string]interface{}, err error)
    param 1 - string - an info hash param 2 - an optional bool - if true the
    results will be returned base64 decoded

func Download(hash string) (result map[string]interface{}, err error)
    hash - an info hash

func Info(params ...interface{}) (result map[string]interface{}, err error)
    Accepts either one []string parameter or any number of string parameters

func Search(params ...string) (result map[string]interface{}, err error)
    Accepts 1-3 parameters of type string param 1 - the search query param 2
    - an optional category param 3 - an optional subcategory

func SetVersion(v string) bool
    Sets the getstrike.net API version to be used. Cuttently only accepts
    "v2"

func Top(params ...string) (result map[string]interface{}, err error)
    Accepts 1-2 parameters of type string param 1 - the category param 2 -
    an optional subcategory


