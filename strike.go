/*
  #strike

  strike is a Go wrapper for the getstrike.net torrent search API: https://getstrike.net/api/
  All of the API's methods are covered and all calls return map[string]interface{}, error

  Variadic functions have been used for some API calls to make the interface more flexible.
  These are documented below.

  [![GoDoc](https://godoc.org/github.com/Pappa/strike?status.svg)](https://godoc.org/github.com/Pappa/strike)
*/
package strike

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
)

var version = "v2"

var api = map[string]map[string]string{
    "v2": map[string]string{
      "Search":"https://getstrike.net/api/v2/torrents/search/?phrase=%s&category=%s&subcategory=%s",
      "Info":"https://getstrike.net/api/v2/torrents/info/?hashes=%s",
      "Download":"https://getstrike.net/api/v2/torrents/download/?hash=%s",
      "Count":"https://getstrike.net/api/v2/torrents/count/",
      "Description":"https://getstrike.net/api/v2/torrents/descriptions/?hash=%s",
      "Top":"https://getstrike.net/api/v2/torrents/top/?category=%s&subcategory=%s",
    },
  }

func getResponse(url string) (resp *http.Response, err error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", url, nil)
    if (err != nil) {
    	return nil, err
    }
    return client.Do(req)
}

func callApi(url string) (result map[string]interface{}, err error) {
    var data map[string]interface{}
	response, err := getResponse(url)
    if (err != nil) {
    	return data, err
    }
    defer response.Body.Close()
    body, err := ioutil.ReadAll(response.Body)
    err = json.Unmarshal(body, &data)
    if (err != nil) {
    	return data, err
    }
    return data, nil
}

/*
  Sets the getstrike.net API version to be used.
  Cuttently only accepts "v2"
 */
func SetVersion(v string) (bool) {
  _, ok := api[v]
  if (ok) {
    version = v
    return true
  } else {
    return false
  }
}