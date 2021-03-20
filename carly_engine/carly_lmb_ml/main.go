package main

import (
	spider_ml "carly_lmb_ml/handler"

	"github.com/aws/aws-lambda-go/lambda"
	pkg "github.com/leonardpahlke/carly_pkg"
)

func Handler(event pkg.CarlyEngineLmbMLEvent) (pkg.CarlyEngineLmbMLResponse, error) {
	return spider_ml.Handler(event)
}

func main() {
	lambda.Start(Handler)
}
