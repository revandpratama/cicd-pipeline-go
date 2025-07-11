package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHandlers(t *testing.T) {
	handler := setupRouter() // your actual handler setup

	server := httptest.NewServer(handler)
	defer server.Close()

	t.Run("should return 200 on GET /", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/")
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})
	t.Run("should return 200 on GET /ping", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/ping")
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		var responseBody struct {
			Ping string `json:"ping"`
		}

		// Decode the response directly into our struct.
		// The decoder handles whitespace automatically.
		err = json.NewDecoder(resp.Body).Decode(&responseBody)
		if err != nil {
			t.Fatalf("failed to decode json: %v", err)
		}

		// Now assert the values from the struct
		expected := "pong"
		if responseBody.Ping != expected {
			t.Errorf("expected ping value %q, got %q", expected, responseBody.Ping)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected 200, got %d", resp.StatusCode)
		}
	})

	t.Run("should return 404 on GET /nonexistent", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/nonexistent")
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected 404, got %d", resp.StatusCode)
		}
	})
}
