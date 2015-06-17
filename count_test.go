package strike

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
    ts := GetTestServer(`{"statuscode":200,"message":8000000}`)
    defer ts.Close()
    api[version]["Count"] = ts.URL
    testMap, err := Count()
    if assert.Nil(t, err) {
    	statusCode, ok := testMap["statuscode"]
    	if assert.Equal(t, true, ok) {
    		var OK_RESPONSE float64 = 200
    		assert.Equal(t, OK_RESPONSE, statusCode, "Status code OK")
    	}
    	message, ok := testMap["message"]
    	if assert.Equal(t, true, ok) {
    		var COUNT float64 = 8000000
    		assert.Equal(t, COUNT, message, "Count message OK")
    	}
    }
}