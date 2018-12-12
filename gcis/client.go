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
	defaultBaseURL   = "https://data.gcis.nat.gov.tw/"
	defaultUserAgent = "go-gcis/" + libraryVersion
)

type service struct {
	client *Client
}

type Client struct {
	HTTPClient *http.Client

	BaseURL   *url.URL
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Services used for talking to different parts of the GCIS API.
	Company *CompanyService
}

// NewClient returns a new GCIS API client.
func NewClient() *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    baseURL,
		UserAgent:  defaultUserAgent,
	}

	c.common.client = c
	c.Company = (*CompanyService)(&c.common)

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)
	uStr := strings.Replace(u.String(), " ", "%20", -1)

	req, err := http.NewRequest(method, uStr, nil)
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
			var body io.Reader
			if resp.ContentLength == 0 {
				// Workaround for empty body
				body = strings.NewReader("[]")
			} else {
				body = resp.Body
			}

			err = json.NewDecoder(body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return response, err
}

func (c *Client) get(ctx context.Context, urlStr string, v interface{}) (*Response, error) {
	req, err := c.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, v)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ErrorResponse reports error caused by an API request.
type ErrorResponse struct {
	Message string
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

// CheckResponse checks the API response for errors.
func CheckResponse(r *http.Response) error {
	// GCIS API always return status code 200
	if code := r.StatusCode; code != 200 {
		return &ErrorResponse{
			Message: fmt.Sprintf("unexpected status code: %d", code),
		}
	}
	if ct := r.Header.Get("Content-type"); !strings.HasPrefix(ct, "application/json") {
		err := &ErrorResponse{
			Message: "unexpected body",
		}

		data, _ := ioutil.ReadAll(r.Body)
		if len(data) > 0 {
			err.Message = string(data)
		}

		return err
	}
	return nil
}

type SearchOptions struct {
	Skip int
	Top  int
}
