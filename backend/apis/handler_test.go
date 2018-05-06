package apis

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestGetTopTopic(t *testing.T) {
}

func TestGetTopicInvalidParameter(t *testing.T) {
	router := SetupRouter()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/topic")

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Convert the JSON response to a map
	var res map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &res)
	// Grab the value & whether or not it exists
	actual, exist := res["message"]

	// Build our expected body
	expected := gin.H{"message": "Invalid parameter"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestGetTopicNotExist(t *testing.T) {
	router := SetupRouter()

	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/topic?name=mock")

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Convert the JSON response to a map
	var res map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &res)
	// Grab the value & whether or not it exists
	actual, exist := res["message"]

	// Build our expected body
	expected := gin.H{"message": "Topic name not exist"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestCreateTopic(t *testing.T) {
}

func TestUpdateTopic(t *testing.T) {
}
