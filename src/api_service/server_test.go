package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldSucceed(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClinet := testServer.Client()

	response, err := testClinet.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello world!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestTHelloWorldFailing(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClinet := testServer.Client()

	body := strings.NewReader("some body")

	response, err := testClinet.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 405, response.StatusCode)
}

func TestHealthEndpointSucceed(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClinet := testServer.Client()

	response, err := testClinet.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "OK!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestHealthEndpointFailing(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClinet := testServer.Client()

	body := strings.NewReader("some body")

	response, err := testClinet.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 405, response.StatusCode)
}
