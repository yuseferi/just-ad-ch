package thttp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TestHttpHandler struct {
	responseMap      map[string]Response
	RecordedRequests map[string][]RecordedRequest
}

type Response struct {
	StatusCode int
	Body       interface{}
}

type RecordedRequest struct {
	QueryParams url.Values
	Headers     http.Header
	Body        []byte
}

func New(responses map[string]Response) *TestHttpHandler {
	return &TestHttpHandler{
		responseMap:      responses,
		RecordedRequests: map[string][]RecordedRequest{},
	}
}

func (th *TestHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestStr := r.Method + " " + r.URL.Path
	if err := th.recordRequest(requestStr, r); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		th.serveResponse(requestStr, w)
	}
}

func (th *TestHttpHandler) recordRequest(requestStr string, r *http.Request) error {
	query := r.URL.Query()
	header := r.Header
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	rr := RecordedRequest{
		QueryParams: query,
		Headers:     header,
		Body:        body,
	}
	if _, found := th.RecordedRequests[requestStr]; !found {
		th.RecordedRequests[requestStr] = []RecordedRequest{rr}
	}
	return nil
}

func (th *TestHttpHandler) serveResponse(requestStr string, w http.ResponseWriter) {
	if response, found := th.responseMap[requestStr]; found {
		if response.Body != nil {
			body, err := json.Marshal(response.Body)
			if err == nil {
				w.WriteHeader(response.StatusCode)
				_, _ = w.Write(body)
			} else {
				fmt.Println("Response marshalling ", err)
				w.WriteHeader(http.StatusNotImplemented)
			}
		} else {
			w.WriteHeader(response.StatusCode)
		}
	} else {
		fmt.Println("No response for ", requestStr)
		w.WriteHeader(http.StatusNotImplemented)
	}
}
