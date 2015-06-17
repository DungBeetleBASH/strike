package strike

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  "fmt"
)

func GetTestServer(responseText string) (*httptest.Server) {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintln(w, responseText)
  }))
}

func TestSetVersion(t *testing.T) {
  assert.Equal(t, SetVersion("v2"), true, "v2 should be valid")
  assert.Equal(t, SetVersion("v3"), false, "v3 should not be valid")
}