package gopocket

import (
	"io"
	"bytes"
	"net/http"
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func Test_BuildPocketRequest(t *testing.T) {
	requestBody := map[string]interface{} {
		"a": "b",
	}

	urlToCall := "http://THE_URI/the/url/to/request"
	req, _ := buildPocketRequest(urlToCall, requestBody)

	assert.Equal(t, urlToCall, req.URL.String())
	assert.Equal(t, "POST", req.Method)
	
	var reqBody map[string]interface{}
	json.NewDecoder(req.Body).Decode(&reqBody)
	assert.Equal(t, requestBody, reqBody)

	assert.Equal(t, "application/json; charset=UTF-8", req.Header.Get("Content-Type"))
	assert.Equal(t, "application/json", req.Header.Get("X-Accept"))
}

func Test_ReadPocketResponseOk(t *testing.T) {
	data := struct {
		Data string `json:"data"`
	}{}

	response := http.Response {
		StatusCode: 200,
		Body: nopCloser{bytes.NewBufferString("{ \"data\": \"abc\" }")},
	}

	err := readPocketResponse(&response, 200, &data)
	if err != nil {
		t.Errorf("Error raised during response check %s", err)
	}
	
	assert.Equal(t, "abc", data.Data)
}