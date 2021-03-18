package main

import (
	spider_parser "carly_lmb_parser/handler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/leonardpahlke/carly_config/pkg"
)

func Handler(event pkg.SpiderParserEvent) (pkg.SpiderParserResponse, error) {
	return spider_parser.Handler(event)
}

func main() {
	lambda.Start(Handler)
}
