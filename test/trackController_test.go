package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	reader io.Reader
)

const trackJson = `
{
	"id":"561a6a040529060013000001",
	"number":1,
	"name":"Get Free",
	"artists":null,
	"duration":126,
	"bpm":0,
	"album":
		{"id":"",
		"name":"",
		"artists":null,
		"nb_track":0,
		"duration":0,
		"cover":{
			"id":"",
			"path":"",
			"url":""
		},
		"images":null,
		"year":0
	},
	"path":"",
	"url":""
}`

func TestCreateTrack(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(w, `{"success": true, "id": 561a6a040529060013000001}`)
	}))
	defer server.Close()

	url := fmt.Sprintf("%s/v1/track", server.URL)
	reader = strings.NewReader(trackJson)
	response, err := http.Post(url, "applicationJSON", nil)
	defer response.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("status code expected: %v but was %d", http.StatusOK, response.StatusCode)
	}
}

func TestGetTrack(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(w, trackJson)
	}))
	defer server.Close()

	response, err := http.Get(server.URL + "/v1/track/get/561a6a040529060013000001")
	defer response.Body.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveTrack(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(w, `{"success": true, "id": 561a6a040529060013000001}`)
	}))
	defer server.Close()

	request, err := http.NewRequest("DELETE", server.URL+"/v1/track/561a6a040529060013000001", nil)
	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()

	if err != nil {
		t.Error(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("status code expected: %v but was %d", http.StatusOK, response.StatusCode)
	}
}
