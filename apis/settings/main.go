package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/apex/gateway"
	"github.com/netlify/gotrue/api"
	"github.com/netlify/gotrue/cmd"
	"github.com/netlify/gotrue/conf"
	"github.com/netlify/gotrue/storage"
	"github.com/sirupsen/logrus"
)

var (
	port       = flag.Int("port", -1, "specify a port")
	configFile = ""
)

func main() {
	flag.Parse()

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

	apis := api.NewEmptyApi(context.Background(), globalConfig, db, cmd.Version)
	http.HandleFunc("/api/settings", func(w http.ResponseWriter, r *http.Request) {
		err = apis.Settings(w, r)
		if err != nil {
			logrus.Error(err)
		}
	})
	listener := gateway.ListenAndServe
	portStr := "n/a"

	if *port != -1 {
		portStr = fmt.Sprintf(":%d", *port)
		listener = http.ListenAndServe
		http.Handle("/", http.FileServer(http.Dir("./public")))
	}

	log.Fatal(listener(portStr, nil))
}
