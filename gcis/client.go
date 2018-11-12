package gcis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	libraryVersion   = "0.0.1"
	defaultBaseURL   = "https://data.gcis.nat.gov.tw"
	defaultUserAgent = "go-gcis/" + libraryVersion
)

type Client struct {
	HTTPClient *http.Client

	BaseURL   *url.URL
	UserAgent string
}

// NewClient returns a new GCIS API client.
func NewClient() *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		BaseURL:   baseURL,
		UserAgent: defaultUserAgent,
	}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a GCIS API response.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	Response *http.Response
	Message  string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %s",
		e.Response.Request.Method,
		e.Response.Request.URL,
		e.Response.StatusCode,
		e.Message)
}

// CheckResponse checks the API response for errors.
func CheckResponse(r *http.Response) error {
	c := r.StatusCode
	if c == 200 {
		if r.ContentLength == 0 {
			return &ErrorResponse{
				Response: r,
				Message:  "unexpected empty body",
			}
		}
		if ct := r.Header.Get("Content-type"); !strings.HasPrefix(ct, "application/json") {
			err := &ErrorResponse{
				Response: r,
				Message:  "unexpected body",
			}

			data, _ := ioutil.ReadAll(r.Body)
			if len(data) > 0 {
				err.Message = string(data)
			}

			return err
		}
		return nil
	}
	// GCIS API always return status code 200
	return fmt.Errorf("unexpected status code: %d", c)
}
