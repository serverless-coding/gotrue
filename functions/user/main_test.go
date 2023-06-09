package main

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

func TestC(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Head(`http://cps-dev.oss-cn-shenzhen.aliyuncs.com/2023-06-07/4247528111733500492/%E6%B5%81%E7%A8%8B%E5%8A%9F%E8%83%BD.mp4?OSSAccessKeyId=LTAI5tK6oBosamYznXrbiA5H&Expires=1686221103&Signature=kJDz0BTSUbnebpX8N%2F2XZHFTVMs%3D`)
	if err == nil {
		fmt.Println(resp.RawBody(), resp.Header().Get("Content-Type"), err)
	}
}
