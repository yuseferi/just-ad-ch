package endpoint

import (
	"fmt"
	"net/http"
)

var (
	NotFound = fmt.Errorf("%d response code has been found", http.StatusNotFound)
)

type Endpoint struct {
	client func() *http.Client
}

func New() *Endpoint {
	return &Endpoint{
		client: Client(),
	}
}
func Client() func() *http.Client {
	return func() *http.Client {
		return &http.Client{}
	}
}
