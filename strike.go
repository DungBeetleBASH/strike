/*
Package strike is a Go wrapper for the getstrike.net torrent search API: https://getstrike.net/api/

Variadic functions have been used for some API calls to make the interface more flexible.
These are documented below.
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
      "Imdb":"https://getstrike.net/api/v2/media/imdb/?imdbid=%s",
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

func callAPI(url string, result interface{}) (err error) {
  response, err := getResponse(url)
  if (err != nil) {
    return err
  }
  defer response.Body.Close()
  body, err := ioutil.ReadAll(response.Body)
  if (err != nil) {
    return err
  }
  err = json.Unmarshal(body, &result)
  if (err != nil) {
    return err
  }
  return nil
}

/*
SetVersion is for future use. It currently accepts only "v2", which is the default value.
*/
func SetVersion(v string) (bool) {
  _, ok := api[v]
  if (ok) {
    version = v
    return true
  }
  return false
}