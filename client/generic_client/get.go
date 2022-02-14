package generic_client

import (
	"context"
	"github.com/yuseferi/just-ad-ch/endpoint"
	"github.com/yuseferi/just-ad-ch/log"
	"go.uber.org/zap"
	"net/http"
)

func (e *Endpoint) GetContentFromUrl(ctx context.Context, urlString string) (body *string, err error) {
	log.Get(ctx).Info("get content from url", zap.String("object", "Url"), zap.String("url", urlString))

	data, err := e.Request(ctx, http.MethodGet, urlString, nil)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, endpoint.NotFound
	}
	body = new(string)
	temp := string(data)
	body = &temp
	return body, err
}
