package generic_client

import "github.com/yuseferi/just-ad-ch/endpoint"

type Endpoint struct {
	*endpoint.Endpoint
}

func New() *Endpoint {
	return &Endpoint{
		endpoint.New(),
	}
}
