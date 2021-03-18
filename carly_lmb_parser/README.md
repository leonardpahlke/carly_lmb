# Carly Lambda Parser

This repository contains lambda code for the lambda function carly-lmb-parser, which is described in the repository carly_aws.
As part of the carly-engine processing AWS StepFunction, this lambda function processes article data.
Specifically, it extracts the article text from an article website dom.

## Structure

```sh
── carly_lmb_ml
├── cmd
│   └── main.go     // use > "go run cmd/main.go" to run lambda locally
├── handler
│   └── handler.go  // lambda internal logic
└── main.go         // lambda entry point
```

## Process

1. Lambda function receives event and starts lambda handler with article dom information
2. swtich article newspaper cases to apply parsing information
3. parse dom with parsing config
