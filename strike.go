package strike

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
)

var version = "v2"

var api = map[string]map[string]string{
    "v2": map[string]string{
      "Search":"https://getstrike.net/api/v2/torrents/search/?phrase=%v&category=%v&subcategory=%v",
      "Info":"https://getstrike.net/api/v2/torrents/info/?hashes=%v",
      "Download":"https://getstrike.net/api/v2/torrents/download/",
      "Count":"https://getstrike.net/api/v2/torrents/count/",
      "Description":"https://getstrike.net/api/v2/torrents/descriptions/?hash=%v",
      "Top":"https://getstrike.net/api/v2/torrents/top/",
      "Imdb":"https://getstrike.net/api/v2/media/imdb/",
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

func SetVersion(v string) {
	version = v
}