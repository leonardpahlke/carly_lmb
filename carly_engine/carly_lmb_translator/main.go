package main

import (
	lmb_translator "carly_lmb_translator/handler"

	pkg "github.com/leonardpahlke/carly_pkg"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(event pkg.CarlyEngineLmbTranslatorEvent) (pkg.CarlyEngineLmbTranslatorResponse, error) {
	return lmb_translator.Handler(event)
}

func main() {
	lambda.Start(Handler)
}
