package integration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yuseferi/just-ad-ch/test/thttp"
	"net/http"
	"testing"
)

func TestApplication_GetUrls(t *testing.T) {
	dataProviders := map[string]struct {
		urls           []string
		parallelNo     int
		expectedResult map[string]string
	}{
		"get data with parallel running one": {
			parallelNo:     1,
			urls:           []string{"https://web.archive.org/web/20131001152630/http://www.spiegel.de/", "https://www.adjust.com/"},
			expectedResult: map[string]string{"https://web.archive.org/web/20131001152630/http://www.spiegel.de/": "55c74e0b88c007f848e9827f49171c65", "https://www.adjust.com/": "9975a7270a4fa7c549003db335bfbc74"},
		},
	}

	for testCase, data := range dataProviders {

		t.Run(testCase, func(t *testing.T) {
			mocks := make(map[string]thttp.Response)

			for key, value := range urlMocks(data.urls) {
				mocks[key] = value
			}

			testApp, err := newApp(&testConfig{})
			if err != nil {
				t.Error(fmt.Sprintf("%v", err))
				return
			}
			result := testApp.Run(data.urls, data.parallelNo)
			for url, value := range result {
				assert.Equal(t, data.expectedResult[url], value)
			}

		})

	}
}

// I tried to mock the websites with this function but du to lake of the time I couldn't finish that.
// so as improvement for this algorithm, we can extend this function to mock the urls instead of getting data from the main url.
func urlMocks(urls []string) map[string]thttp.Response {
	urlMocks := map[string]thttp.Response{}
	for _, urlItem := range urls {
		urlMocks[urlItem] = thttp.Response{
			StatusCode: http.StatusOK,
			Body:       "<html><body>just test</body></html",
		}
	}
	return urlMocks
}
