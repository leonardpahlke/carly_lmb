package main

import (
	spider_ml "carly_lmb_ml/handler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/leonardpahlke/carly_config/pkg"
)

func Handler(event pkg.SpiderMLEvent) (pkg.SpiderMLResponse, error) {
	return spider_ml.Handler(event)
}

func main() {
	lambda.Start(Handler)
}