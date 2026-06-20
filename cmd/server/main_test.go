package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePage(t *testing.T) {
	handler := newHandler()

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	body := response.Body.String()

	assertContains(t, body, "I am Szymon")
	assertContains(t, body, "Projects")
	assertContains(t, body, "This Website")
	assertContains(t, body, "June 2026")
	assertContains(t, body, "Posts")
	assertContains(t, body, "Sixteenth Attempt")
	assertContains(t, body, `href="/sixteenth-attempt"`)
}

func assertContains(t *testing.T, body string, want string) {
	t.Helper()

	if !strings.Contains(body, want) {
		t.Fatalf("expected body to contain %q\n\nbody:\n%s", want, body)
	}
}
