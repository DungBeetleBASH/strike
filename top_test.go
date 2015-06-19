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
    result, err := Top("Movies")
    if assert.Nil(t, err) {
        assert.Equal(t, 2, len(result.Torrents), "Torrents OK")
        assert.Equal(t, "89599BF4DC369A3A8ECA26411C5CCF922D78B486", result.Torrents[0].Hash, "Torrent hash OK")
    }
}

func TestTopTwoCategories(t *testing.T) {
    ts := GetTestServer(`{"results":2,"torrents":[{"torrent_hash":"89599BF4DC369A3A8ECA26411C5CCF922D78B486"},{"torrent_hash":"BE4AA07856542F2DCD4E12362EF9F0A1AC88C358"}]}`)
    defer ts.Close()
    api[version]["Top"] = strings.Join([]string{ts.URL, "?category=%s&subcategory=%s"}, "")
    result, err := Top("Movies", "Asian")
    if assert.Nil(t, err) {
    	assert.Equal(t, 2, len(result.Torrents), "Torrents OK")
    	assert.Equal(t, "89599BF4DC369A3A8ECA26411C5CCF922D78B486", result.Torrents[0].Hash, "Torrent hash OK")
    }
}