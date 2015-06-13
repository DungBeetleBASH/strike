package strike

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
    "fmt"
)

var version = "v2"

var api = map[string]map[string]string{
    "v2": map[string]string{
      "Search":"https://getstrike.net/api/v2/torrents/search/?phrase=%v&category=%v&subcategory=%v",
      "Info":"https://getstrike.net/api/v2/torrents/info/",
      "Download":"https://getstrike.net/api/v2/torrents/download/",
      "Count":"https://getstrike.net/api/v2/torrents/count/",
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

func callApi(url string) (result interface{}, err error) {
    var data interface{}
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

func getQuery(qType string, params ...string) (result string) {
	args := make([]interface{}, len(params))
	for i, v := range params {
	    args[i] = v
	}
	return fmt.Sprintf(api["v2"][qType], args...)
}

func SetVersion(v string) {
	version = v
}