package main

import (
   "io/ioutil"
   "net/http"
   "net/http/httptest"
   "testing"
)

func TestServeHTTP(t *testing.T) {
	server := httptest.NewServer(RunServer())
        defer server.Close()


	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
        if resp.StatusCode != 200 {
                t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
        }
	expected := `{"message":"hello world!!"}`
	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(actual) {
		t.Errorf("Expected the message '%s' but got '%s'\n", expected,actual)
	}
}
