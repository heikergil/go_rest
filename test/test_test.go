package main

import (
  "testing"
 // "github.com/stretchr/testify/assert"
  "net/http/httptest"
  "net/http"
  //"github.com/heikergil/go_rest/controllers"
  "github.com/gin-gonic/gin"
)


type header struct {
	Key   string
	Value string
}
// PerformRequest for testing gin router.
func PerformRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}



func TestRouterGroupGETNoRootExistsRouteOK(t *testing.T) {
  // SETUP
  r := gin.Default()
  r.GET("/localhost:8080/users", func(c *gin.Context) {})

  // RUN
  w := PerformRequest(r, "GET", "/")

  // TEST
  if w.Code != http.StatusNotFound {
      // If this fails, it's because httprouter needs to be updated to at least f78f58a0db
      t.Errorf("Status code should be %v, was %d. Location: %s", http.StatusNotFound, w.Code, w.HeaderMap.Get("Location"))
  }
}

// func TestHealthzHandler(t * testing.T) {
//   //arrange
//   req, _ := http.NewRequest("GET", "/localhost:8080", nil)
//   w := httptest.NewRecorder()
//   //add header to the request
//   r.Header.Add("User-Agent", "myClient")
//   //act
//   controllers.FindUser(r)
//   //assert
//   resp := w.Result()
//   //body, _ := ioutil.ReadAll(resp.Body)
//   assert.Equal(t, 200, resp.StatusCode, "expected 200 status code")
// }