package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/name?name=tucker", nil) // a request created.

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// barHandler(res, req) // instead of calling barHandler, we need to use mux.

	assert.Equal(http.StatusOK, res.Code)
	// res.Body() // res.Body() is a buffer class, hence, not possible to use it as is.
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello tucker!", string(data))
	/*	if res.Code != http.StatusOK {
		t.Fatal("Routing Failed!", res.Code)
	} */
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	//req := httptest.NewRequest("GET", "/foo", nil)
	req := httptest.NewRequest("POST", "/foo", strings.NewReader(`{"first_name":"Coding", "last_name":"Forest", "email":"coding@forest.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)
	//assert.Equal(http.StatusBadRequest, res.Code)
	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("Coding", user.FirstName)
	assert.Equal("Forest", user.LastName)
	assert.Equal("coding@forest.com", user.Email)
}

/*
3 - Testing, Tucker Programming, https://www.youtube.com/watch?v=1BeLiLwCS1c&list=PLy-g2fnSzUTDALoERcKDniql16SAaQYHF&index=3
4 - FILE upload, Tucker Programming, https://www.youtube.com/watch?v=bRqhlxWspIk&list=PLy-g2fnSzUTDALoERcKDniql16SAaQYHF&index=4
*/
