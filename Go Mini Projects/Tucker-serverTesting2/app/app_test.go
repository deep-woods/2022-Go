package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHanlder(t *testing.T) {
	assert := assert.New(t)

	// Create a mock response and request recorders.
	res := httptest.NewRecorder() // this doesn't use the real http protocol (does not use the network).
	req := httptest.NewRequest("GET", "/", nil)

	// IndexHandler(res, req)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

/*	assert.Equal(http.StatusOK, res.Code)
	if res.Code != http.StatusOK {
		t.Fatal("Request Failed - ", res.Code)
	}
*/

func TestBarHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	// Create a mock response and request recorders.
	res := httptest.NewRecorder() // this doesn't use the real http protocol (does not use the network).
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}

func TestBarHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	name := "Forest"

	// Create a mock response and request recorders.
	res := httptest.NewRecorder() // this doesn't use the real http protocol (does not use the network).
	req := httptest.NewRequest("GET", fmt.Sprintf("/bar?name=%v", name), nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal(fmt.Sprintf("Hello %v", name), string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"Coding",
							"last_name":"Forest",
							"email":"coding@forest.here"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("Coding", user.FirstName)
	assert.Equal("Forest", user.LastName)
}
