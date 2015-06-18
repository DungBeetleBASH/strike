package strike

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "strings"
)

func TestTopNoCategories(t *testing.T) {
    ts := GetTestServer(`{"statuscode":404,"message":"Please enter a category"}`)
    defer ts.Close()
    api[version]["Top"] = ts.URL
    _, err := Top()
    assert.NotNil(t, err)
}

func TestTopOneCategory(t *testing.T) {
    ts := GetTestServer(`{"results":2,"torrents":[{"torrent_hash":"89599BF4DC369A3A8ECA26411C5CCF922D78B486"},{"torrent_hash":"BE4AA07856542F2DCD4E12362EF9F0A1AC88C358"}]}`)
    defer ts.Close()
    api[version]["Top"] = strings.Join([]string{ts.URL, "?category=%s&subcategory=%s"}, "")
    testMap, err := Top("Movies")
    if assert.Nil(t, err) {
    	statusCode, ok := testMap["results"]
    	if assert.Equal(t, true, ok) {
    		var OK_RESPONSE float64 = 2
    		assert.Equal(t, OK_RESPONSE, statusCode, "Status code OK")
    	}
    	torrents, ok := testMap["torrents"]
    	if assert.Equal(t, true, ok) {
    		assert.Equal(t, 2, len(torrents.([]interface{})), "Torrents OK")
    		assert.Equal(t, "89599BF4DC369A3A8ECA26411C5CCF922D78B486", torrents.([]interface{})[0].(map[string]interface{})["torrent_hash"], "Torrent hash OK")
    	}
    }
}

func TestTopTwoCategories(t *testing.T) {
    ts := GetTestServer(`{"results":2,"torrents":[{"torrent_hash":"89599BF4DC369A3A8ECA26411C5CCF922D78B486"},{"torrent_hash":"BE4AA07856542F2DCD4E12362EF9F0A1AC88C358"}]}`)
    defer ts.Close()
    api[version]["Top"] = strings.Join([]string{ts.URL, "?category=%s&subcategory=%s"}, "")
    testMap, err := Top("Movies", "Asian")
    if assert.Nil(t, err) {
    	statusCode, ok := testMap["results"]
    	if assert.Equal(t, true, ok) {
    		var OK_RESPONSE float64 = 2
    		assert.Equal(t, OK_RESPONSE, statusCode, "Status code OK")
    	}
    	torrents, ok := testMap["torrents"]
    	if assert.Equal(t, true, ok) {
    		assert.Equal(t, 2, len(torrents.([]interface{})), "Torrents OK")
    		assert.Equal(t, "89599BF4DC369A3A8ECA26411C5CCF922D78B486", torrents.([]interface{})[0].(map[string]interface{})["torrent_hash"], "Torrent hash OK")
    	}
    }
}