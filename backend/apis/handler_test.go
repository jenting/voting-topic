package apis

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestGetTopTopic(t *testing.T) {
}

func TestGetTopicInvalidParameter(t *testing.T) {
	router := SetupRouter()

	// Perform a GET request with that handler.
	req, _ := http.NewRequest("GET", "/topic", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err := json.Unmarshal([]byte(resp.Body.String()), &respBody)
	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

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
	req, _ := http.NewRequest("GET", "/topic?name=mock", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err := json.Unmarshal([]byte(resp.Body.String()), &respBody)
	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "Topic name not exist"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestGetTopicOK(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{Name: "mock2", Upvote: 100, Downvote: 200}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a POST request with that handler.
	req, _ := http.NewRequest("POST", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Perform a GET request with that handler.
	req, _ = http.NewRequest("GET", "/topic?name=mock2", nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusOK, resp.Code)

	// Convert the JSON response to a map
	var respBody topicInfo
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.EqualValues(t, 100, respBody.Upvote)
	assert.EqualValues(t, 200, respBody.Downvote)
}

func TestCreateTopicNoName(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a POST request with that handler.
	req, _ := http.NewRequest("POST", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "Invalid JSON parameter"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestCreateTopicOK(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{Name: "1-1"}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a POST request with that handler.
	req, _ := http.NewRequest("POST", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "OK"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestCreateTopicAlreadyExist(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{Name: "1-2"}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a POST request with that handler.
	req, _ := http.NewRequest("POST", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Perform a POST request with that handler.
	req, _ = http.NewRequest("POST", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "Topic name already exist"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func TestCreateTopicOverLen(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{Name: randStringRunes(maxTopicNameLen + 1)}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a POST request with that handler.
	req, _ := http.NewRequest("POST", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "Invalid parameter"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestUpdateTopicNoName(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a PUT request with that handler.
	req, _ := http.NewRequest("PUT", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "Invalid JSON parameter"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestUpdateTopicNotExist(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{Name: "2-1"}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a PUT request with that handler.
	req, _ := http.NewRequest("PUT", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "Topic name is not exist"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestUpdateTopicOK(t *testing.T) {
	router := SetupRouter()

	reqBody := topicInfo{Name: "2-2"}
	b, err := json.Marshal(reqBody)
	assert.Equal(t, nil, err, "JSON marshal failed")

	// Perform a POST request with that handler.
	req, _ := http.NewRequest("POST", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Perform a PUT request with that handler.
	reqBody = topicInfo{Name: "2-2", Upvote: 1, Downvote: 2}
	b, err = json.Marshal(reqBody)
	req, _ = http.NewRequest("PUT", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]

	// Build our expected body
	expected := gin.H{"message": "OK"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}
