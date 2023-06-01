package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestM(t *testing.T) {

	os.Setenv("_LAMBDA_SERVER_PORT", "5024")
	os.Setenv("AWS_LAMBDA_RUNTIME_API", "go")

	go start()

	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second * 20)
	}
}
