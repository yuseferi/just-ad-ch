package main

import (
	"context"
	"github.com/yuseferi/just-ad-ch/app"
)

func main() {
	config := app.NewConfig()
	//init context
	ctx, ctxCancel := context.WithCancel(context.Background())

	// init application
	application, err := app.New(ctx, ctxCancel, config)
	if err != nil {
		panic(err)
	}
	defer application.Close()
	application.Start()
}
