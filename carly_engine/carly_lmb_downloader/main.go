package main

import (
	lmb_downloader "carly_lmb_downloader/handler"

	"github.com/aws/aws-lambda-go/lambda"
	pkg "github.com/leonardpahlke/carly_pkg"
)

func Handler(event pkg.CarlyEngineLmbDownloaderEvent) (pkg.CarlyEngineLmbDownloaderResponse, error) {
	return lmb_downloader.Handler(event)
}

func main() {
	lambda.Start(Handler)
}
