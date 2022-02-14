package endpoint

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/yuseferi/just-ad-ch/log"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (e *Endpoint) Request(ctx context.Context, method string, url string, body interface{}) (result []byte, err error) {
	var (
		resp *http.Response
		req  *http.Request
	)

	defer func(start time.Time) {
		logger := log.Get(ctx).With(
			zap.String("url", url),
			zap.String("method", method),
			zap.Any("body", body),
			zap.ByteString("result", result),
			zap.Duration("duration", time.Since(start)),
			zap.Error(err))
		if logger == nil {
			return
		}
		if resp != nil {
			logger = logger.With(
				zap.String("status", resp.Status),
				zap.Int64("content_length", resp.ContentLength))
			if resp.Header != nil {
				logger = logger.With(zap.Any("headers", resp.Header))
			}
		}
		if err != nil {
			if err == NotFound {
				logger.Info("got Response")
			} else {
				logger.Warn("got Response")
			}
		} else {
			logger.Debug("got Response")
		}
	}(time.Now())

	req, err = e.httpRequest(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err = e.client().Do(req)
	if err != nil {
		return
	}
	defer func() {
		if resp.Body != nil {
			if cErr := resp.Body.Close(); cErr != nil {
				log.Get(ctx).Warn("Error on close: HTTP response body", zap.Error(cErr))
			}
		}
	}()

	result, err = ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusNotFound {
		log.Get(ctx).Debug("got StatusNotFound",
			zap.String("method", method),
			zap.String("url", url))
		err = NotFound
	} else if resp.StatusCode >= http.StatusMultipleChoices {
		err = fmt.Errorf("wrong response status: %s", resp.Status)
	}
	return
}

func (e *Endpoint) httpRequest(ctx context.Context, method string, urlString string, body interface{}) (*http.Request, error) {
	target, err := url.Parse(urlString)

	// add scheme if it does not have
	if target.Scheme == "" {
		urlString = "https://" + urlString
	}

	target, err = url.Parse(urlString)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(nil)
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, target.String(), buffer)
	if err != nil {
		return nil, err
	}

	return req.WithContext(ctx), nil
}
