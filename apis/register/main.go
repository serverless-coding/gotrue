package main

import (
	"context"
	"fmt"
	"time"

	"github.com/netlify/gotrue/api"
	"github.com/netlify/gotrue/cmd"
	"github.com/netlify/gotrue/conf"
	"github.com/netlify/gotrue/storage"
	"github.com/sirupsen/logrus"
)

var configFile = ""

func main() {
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

	globalConfig.MultiInstanceMode = true
	_, api := api.NewRegisterAPI(context.Background(), globalConfig, db, cmd.Version)

	l := fmt.Sprintf("%v:%v", globalConfig.API.Host, globalConfig.API.Port)
	logrus.Infof("GoTrue API started on: %s", l)
	api.ListenAndServe(l)
}
