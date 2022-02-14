package service

import (
	"context"
	"github.com/yuseferi/just-ad-ch/client/generic_client"
	"github.com/yuseferi/just-ad-ch/helpers"
	"github.com/yuseferi/just-ad-ch/log"
	"go.uber.org/zap"
	"sync"
)

type GenericHttpService struct {
	genericClient *generic_client.Endpoint
}

func NewGenericHttpService(genericClient *generic_client.Endpoint) *GenericHttpService {
	return &GenericHttpService{
		genericClient: genericClient,
	}
}

// GetUrls get urls contents
func (s *GenericHttpService) GetUrls(ctx context.Context, urls []string, parallelNo int) (result map[string]string) {
	var wg sync.WaitGroup
	ch := make(chan string, parallelNo)
	result = make(map[string]string, len(urls))

	for _, urlString := range urls {

		ch <- urlString
		wg.Add(1)
		urlString := urlString
		go func() {
			defer wg.Done()
			res, err := s.worker(ctx, ch)
			if err != nil {
				log.Get(ctx).Error("error on get url content", zap.String("url", urlString))
				wg.Done()
				return
			}
			result[urlString] = helpers.MD5(*res)

		}()

	}
	wg.Wait()
	return result
}

func (s *GenericHttpService) worker(ctx context.Context, ch <-chan string) (data *string, err error) {
	url := <-ch
	data, err = s.genericClient.GetContentFromUrl(ctx, url)
	return
}
