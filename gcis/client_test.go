package gcis

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the Apple Music client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a gcis.Client that is configured to talk to that test server.
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// GCIS client configured to use test server
	client = NewClient()
	u, _ := url.Parse(server.URL)
	client.BaseURL = u
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient()

	if got, want := c.BaseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, defaultUserAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient()

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	req, _ := c.NewRequest("GET", inURL, nil)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, "ok")
	})

	req, _ := client.NewRequest("GET", "/", nil)
	client.Do(context.Background(), req, nil)
}

func TestCheckResponse(t *testing.T) {
	tests := []struct {
		statusCode int
		body       string
		want       string
	}{
		{
			http.StatusOK,
			"",
			"unexpected body",
		},
		{
			http.StatusOK,
			"$format參數有誤，請查明後繼續。",
			"$format參數有誤，請查明後繼續。",
		},
		{
			http.StatusNotFound,
			"",
			"unexpected status code: 404",
		},
		{
			http.StatusHTTPVersionNotSupported,
			"",
			"unexpected status code: 505",
		},
	}

	for i, test := range tests {
		res := &http.Response{
			Request:       &http.Request{},
			StatusCode:    test.statusCode,
			Body:          ioutil.NopCloser(strings.NewReader(test.body)),
			ContentLength: int64(len(test.body)),
		}
		err := CheckResponse(res)
		if err == nil {
			t.Errorf("(%v) Expected error response", i)
		}
		if got := err.Error(); test.want != got {
			t.Errorf("(%v) Expected: %v, got: %v", i, test.want, got)
		}
	}
}
