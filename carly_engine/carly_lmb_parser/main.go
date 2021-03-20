package main

import (
	spider_parser "carly_lmb_parser/handler"

	"github.com/aws/aws-lambda-go/lambda"
	pkg "github.com/leonardpahlke/carly_pkg"
)

func Handler(event pkg.CarlyEngineLmbParserEvent) (pkg.CarlyEngineLmbParserResponse, error) {
	return spider_parser.Handler(event)
}

func main() {
	lambda.Start(Handler)
}
