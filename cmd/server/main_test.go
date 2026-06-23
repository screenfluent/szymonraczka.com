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

	assertContains(t, body, "Szymon Rączka")
	assertContains(t, body, "Projects")
	assertContains(t, body, "This Website")
	assertContains(t, body, "June 2026")
	assertContains(t, body, "Posts")
	assertContains(t, body, "Sixteenth Attempt")
	assertContains(t, body, `href="/sixteenth-attempt"`)
	assertContains(t, body, `href="/now"`)
}

func TestSixteenthAttemptPost(t *testing.T) {
	handler := newHandler()

	request := httptest.NewRequest(http.MethodGet, "/sixteenth-attempt", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	body := response.Body.String()

	assertContains(t, body, "Sixteenth Attempt")
}

func TestNowPage(t *testing.T) {
	handler := newHandler()

	request := httptest.NewRequest(http.MethodGet, "/now", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	body := response.Body.String()

	assertContains(t, body, "now")
	assertContains(t, body, "Current focus goes here.")
}

func assertContains(t *testing.T, body string, want string) {
	t.Helper()

	if !strings.Contains(body, want) {
		t.Fatalf("expected body to contain %q\n\nbody:\n%s", want, body)
	}
}

func TestPublicPagesUseSharedShell(t *testing.T) {
	handler := newHandler()

	paths := []string{"/", "/sixteenth-attempt", "/now"}

	for _, path := range paths {
		t.Run(path, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("expected status 200, got %d", response.Code)
			}

			body := response.Body.String()

			assertContains(t, body, `<meta name="viewport" content="width=device-width, initial-scale=1">`)
			assertContains(t, body, `<header>`)
			assertContains(t, body, `href="/"`)
			assertContains(t, body, `<main>`)
			assertContains(t, body, `<link rel="stylesheet" href="/static/site.css">`)
		})
	}
}

func TestStaticStylesheet(t *testing.T) {
	handler := newHandler()

	request := httptest.NewRequest(http.MethodGet, "/static/site.css", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	body := response.Body.String()

	assertContains(t, body, "body")
}
