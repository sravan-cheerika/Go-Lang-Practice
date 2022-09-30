package main

import (
	"net/http"
	"net/http/httptest"
	"sum_test/subDirectory"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	subDirectory.PlayerServer(response, request)
	t.Run("returns Pepper's score", func(t *testing.T) {
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestMethod2(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/test2", nil)
	response := httptest.NewRecorder()
	subDirectory.Method2(response, request)
	t.Run("returns test2", func(t *testing.T) {
		got := response.Body.String()
		want := "Hello Team"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		
	})
}