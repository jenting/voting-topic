package apis

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.Set("logtostderr", "true")
}

func TestGetTopTopic(t *testing.T) {
}

func TestGetTopicInvalidUUID(t *testing.T) {
	router := SetupRouter()

	// Perform a GET request with that handler.
	req, _ := http.NewRequest("GET", fmt.Sprintf("/topic?uuid=%v", "testuid"), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err := json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]
	expected := gin.H{"message": "Invalid input uuid"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestGetTopicNotExist(t *testing.T) {
	router := SetupRouter()

	// Perform a POST request with that handler.
	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err, "Generate uuid failed")

	// Perform a GET request with that handler.
	req, _ := http.NewRequest("GET", fmt.Sprintf("/topic?uuid=%v", uid), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 400
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Convert the JSON response to a map
	var respBody map[string]string
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Grab the value & whether or not it exists
	actual, exist := respBody["message"]
	expected := gin.H{"message": "Topic not exist"}

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

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, resp.Code)

	// Response body
	var respPostBody topicInfo
	err = json.Unmarshal([]byte(resp.Body.String()), &respPostBody)

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.EqualValues(t, "mock2", respPostBody.Name)
	assert.EqualValues(t, 100, respPostBody.Upvote)
	assert.EqualValues(t, 200, respPostBody.Downvote)

	// Perform a GET request with that handler.
	req, _ = http.NewRequest("GET", fmt.Sprintf("/topic?uuid=%v", respPostBody.UID), nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, resp.Code)

	// Convert the JSON response to a map
	var respGetBody topicInfo
	err = json.Unmarshal([]byte(resp.Body.String()), &respGetBody)

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.EqualValues(t, "mock2", respGetBody.Name)
	assert.EqualValues(t, 100, respGetBody.Upvote)
	assert.EqualValues(t, 200, respGetBody.Downvote)
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
	var respBody topicInfo
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.EqualValues(t, "1-1", respBody.Name)
	assert.EqualValues(t, 0, respBody.Upvote)
	assert.EqualValues(t, 0, respBody.Downvote)
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
	expected := gin.H{"message": "Topic name over length"}

	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exist)
	assert.Equal(t, expected["message"], actual)
}

func TestUpdateTopicNotExist(t *testing.T) {
	router := SetupRouter()

	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err)

	reqBody := topicInfo{UID: uid, Name: "2-1", Upvote: 1, Downvote: 2}
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
	expected := gin.H{"message": "UUID not exist"}

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

	// Convert the JSON response
	var respBody topicInfo
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)
	assert.Nil(t, err)
	assert.NotEqual(t, uuid.Nil, respBody.UID)

	// Perform a PUT request with that handler.
	reqBody = topicInfo{UID: respBody.UID, Name: "2-2", Upvote: 1, Downvote: 2}
	b, err = json.Marshal(reqBody)
	req, _ = http.NewRequest("PUT", "/topic", bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// Assert we encoded correctly, the request gives a 200
	assert.Equal(t, http.StatusOK, resp.Code)

	// Convert the JSON response
	err = json.Unmarshal([]byte(resp.Body.String()), &respBody)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, respBody.Upvote)
	assert.EqualValues(t, 2, respBody.Downvote)
}
