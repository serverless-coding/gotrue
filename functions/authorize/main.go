package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/netlify/gotrue/api"
	"github.com/netlify/gotrue/cmd"
	"github.com/netlify/gotrue/conf"
	"github.com/netlify/gotrue/storage"
	"github.com/sirupsen/logrus"
)

var configFile = ""

type hf func(w http.ResponseWriter, r *http.Request, config *conf.Configuration) error

func handlerFunc(f hf, config *conf.Configuration) http.HandlerFunc {
	return func(ww http.ResponseWriter, rr *http.Request) {
		err := f(ww, rr, config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	start()
}

func start() {
	globalConfig, err := conf.LoadGlobal(configFile)
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %+v", err)
	}
	if globalConfig.OperatorToken == "" {
		logrus.Fatal("Operator token secret is required")
	}

	var db *storage.Connection
	// try a couple times to connect to the database
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Duration((i-1)*100) * time.Millisecond)
		db, err = storage.Dial(globalConfig)
		if err == nil {
			break
		}
		logrus.WithError(err).WithField("attempt", i).Warn("Error connecting to database")
	}
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}
	defer db.Close()

	config, err := conf.LoadConfig(configFile)
	if err != nil {
		logrus.Fatalf("Error opening database: %+v", err)
	}

	globalConfig.MultiInstanceMode = false
	lambda.Start(func(ctx context.Context, req events.APIGatewayProxyRequest) (resp events.APIGatewayProxyResponse, err error) {
		router, apis := api.NewRegisterAPI(ctx, globalConfig, db, cmd.Version)
		fmt.Println("-----router:", router)
		apis.CtxConfig = config
		adapter := httpadapter.New(handlerFunc(apis.ExternalProviderRedirect, config))
		resp, err = adapter.Proxy(req)
		resp.StatusCode = 200
		return
	})
}

// npm install netlify-cli -g
