package inmet

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "http://sisdagro.inmet.gov.br/sisdagro/app/"
	userAgent      = "go-inmet"
)

var (
	errTraling = errors.New("BaseURL must have a trailing slash")
)

// Client manages communication
type Client struct {
	client    *http.Client // HTTP client used to communicate
	BaseURL   *url.URL
	UserAgent string

	commom service

	CAD     *CADService
	Soil    *SoilService
	Station *StationService
	BHC     *BHCService
	Culture *CultureService
}

type service struct {
	client *Client
}

// addQuery add paramenters to url
func addQuery(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}
	qs, err := query.Values(opt)

	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewClient returns new Inmet client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
	c.commom.client = c
	c.CAD = &CADService{client: c}
	c.Soil = &SoilService{client: c}
	c.Station = &StationService{client: c}
	c.BHC = &BHCService{client: c}
	c.Culture = &CultureService{client: c}
	return c
}

// NewRequest TODO
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, errTraling
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf *strings.Reader
	if body != nil {
		form := url.Values{}
		v := reflect.ValueOf(body).Elem()
		for i := 0; i < v.NumField(); i++ {
			form.Add(v.Type().Field(i).Tag.Get("json"), v.Field(i).String())
		}
		buf = strings.NewReader(form.Encode())
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response TODO
type Response struct {
	*http.Response
}

func newResponse(resp *http.Response) *Response {
	response := &Response{Response: resp}
	return response
}

// ResponseError TODO
type ResponseError struct {
	Response *http.Response
}

func (rr *ResponseError) Error() string {
	return fmt.Sprintf("%v %v: %v", rr.Response.Request.Method, rr.Response.Request.URL, rr.Response.StatusCode)
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	response := newResponse(resp)

	err = CheckResponse(resp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return response, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	// errorResponse := &ResponseError{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	fmt.Println(data)
	// if err == nil && data != nil {
	// 	json.Unmarshal(data, errorResponse)
	// }
	return nil
}
