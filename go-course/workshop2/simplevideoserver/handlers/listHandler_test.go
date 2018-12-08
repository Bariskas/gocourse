package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func TestListHandler(t *testing.T) {
	w := httptest.NewRecorder()
	listHandler(w, nil)
	response := w.Result()
	if response.StatusCode != http.StatusOK {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", response.StatusCode, http.StatusOK)
	}

	jsonString, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	items := make([]VideoListItem, 10)
	if err = json.Unmarshal(jsonString, &items); err != nil {
		t.Errorf("Can't parse json response with error %v", err)
	}
}

