package main

import (
	spider_downloader "carly_lmb_downloader/handler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/leonardpahlke/carly_config/pkg"
)

func Handler(event pkg.SpiderDownloaderEvent) (pkg.SpiderDownloaderResponse, error) {
	return spider_downloader.Handler(event)
}

func main() {
	lambda.Start(Handler)
}
