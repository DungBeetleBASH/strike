package strike

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
    ts := GetTestServer(`{"statuscode":200,"message":8000000}`)
    defer ts.Close()
    api[version]["Count"] = ts.URL
    count, err := Count()
    if assert.Nil(t, err) {
    	var expecting int64 = 8000000
    	assert.Equal(t, expecting, count, "Count OK")
    }
}