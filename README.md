#strike

    strike is a Go wrapper for the getstrike.net torrent search API: (https://getstrike.net/api/) 
    All of the API's methods are covered and all calls return map[string]interface{}, error

    Variadic functions have been used for some API calls to make the interface more flexible. 
    These are documented below.

[![GoDoc](https://godoc.org/github.com/Pappa/strike?status.svg)](https://godoc.org/github.com/Pappa/strike)

##API

**func Search(params ...string) (result map[string]interface{}, err error)**
    
    Accepts 1-3 parameters of type string
    
    params[0] string - the search query 
    
    params[1] string - optional - category
    
    params[2] string - optional - subcategory

**func Description(params ...interface{}) (result map[string]interface{}, err error)**

    params[0] string - info hash

    params[1] bool - optional - if true the results will be returned base64 decoded

**func Info(params ...interface{}) (result map[string]interface{}, err error)**
    
    Accepts either one []string parameter or any number of string parameters
    
    params[] string - info hash 

**func Download(hash string) (result map[string]interface{}, err error)**
    
    hash string - info hash 

**func Count() (result map[string]interface{}, err error)**
    
    Returns a map containing the current total torrent count on https://getstrike.net/

**func Top(params ...string) (result map[string]interface{}, err error)**
    
    Accepts 1-2 parameters of type string 
    
    params[0] string - category
    
    params[1] string - optional - subcategory

**func SetVersion(v string) bool**
    
    Sets the getstrike.net API version to be used. Cuttently only accepts "v2"


