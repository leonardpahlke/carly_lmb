package main

import (
	spider_translator "carly_lmb_translator/handler"

	"github.com/leonardpahlke/carly_config/pkg"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(event pkg.SpiderTranslatorEvent) (pkg.SpiderTranslatorResponse, error) {
	return spider_translator.Handler(event)
}

func main() {
	lambda.Start(Handler)
}
