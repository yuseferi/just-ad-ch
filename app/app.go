package app

import (
	"context"
	"flag"
	"fmt"
	"github.com/yuseferi/just-ad-ch/client/generic_client"
	"github.com/yuseferi/just-ad-ch/service"
	"go.uber.org/zap"
	"io"
	"math"
	"sync"
	"time"
)

type Application struct {
	Config             *Config
	Ctx                context.Context
	Error              chan error
	Logger             *zap.Logger
	WaitGroup          *sync.WaitGroup
	ctxCancel          context.CancelFunc
	GenericClient      *generic_client.Endpoint
	genericHttpService *service.GenericHttpService
}

func New(ctx context.Context, ctxCancel context.CancelFunc, config *Config) (app *Application, err error) {
	logger, err := NewLogger(config.Level)
	if err != nil {
		return nil, err
	}

	app = &Application{
		Config:    config,
		Logger:    logger,
		Error:     make(chan error, math.MaxUint8),
		WaitGroup: new(sync.WaitGroup),
	}

	app.Ctx = ctx
	app.ctxCancel = ctxCancel

	defer func() {
		if err != nil {
			app.Close()
		}
	}()

	if err = app.setClients(); err != nil {
		logger.Panic("cannot setup the client", zap.Error(err))
		return nil, err
	}

	app.setupServices()

	return app, nil
}

// Endpoints
func (app *Application) setClients() (err error) {
	app.Logger.Debug("Client setup")
	app.GenericClient = generic_client.New()

	return
}

func (app *Application) Close() {
	app.Logger.Debug("Application stops")

}

func (app *Application) Start() {
	defer app.Stop()
	// list of requested teams
	parallelNo := flag.Int("parallel", 10, "concurrent request")
	flag.Parse()
	urls := flag.Args()
	//fmt.Println(*parallelNo)
	//fmt.Println("tail:", urls)
	app.Run(urls, int(*parallelNo))
}

func (app *Application) Stop() {
	app.Logger.Info("service stopping...")
	app.ctxCancel()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)

	go func() {
		defer cancel()
		app.WaitGroup.Wait()
	}()

	<-ctx.Done()

	if ctx.Err() != context.Canceled {
		app.Logger.Panic("service stopped with timeout")
	} else {
		app.Logger.Info("service stopped with success")
	}
}

func (app *Application) Closer(closer io.Closer, scope string) {
	if closer != nil {
		if err := closer.Close(); err != nil {
			app.Logger.Warn("closer error", zap.String("scope", scope), zap.Error(err))
		}
	}
}
func (app *Application) setupServices() {
	app.genericHttpService = service.NewGenericHttpService(app.GenericClient)
}

func (app *Application) Run(urls []string, parallelNo int) (output map[string]string) {

	output = app.genericHttpService.GetUrls(app.Ctx, urls, parallelNo)

	for url, body := range output {
		fmt.Printf("%s => %s\n", url, body)
	}
	return
}
