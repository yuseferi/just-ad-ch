package integration

import (
	"context"
	"github.com/yuseferi/just-ad-ch/app"
	"os"
	"testing"
)

var (
	config *app.Config
)

func TestMain(m *testing.M) {
	config = app.NewConfig()
	os.Exit(m.Run())
}

func newApp(testConfig *testConfig) (*app.Application, error) {
	//init context
	ctx, ctxCancel := context.WithCancel(context.Background())

	return app.New(ctx, ctxCancel, config)
}
