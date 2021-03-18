# Carly Lambda Translator

This repository contains lambda code for the lambda function carly-lmb-translator, which is described in the repository carly_aws.
As part of the carly-engine processing AWS StepFunction, this lambda function processes article data.
Specifically, it uses AWS Translate to tanslate german articles to english and stores the translated article information in s3.

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

1. check if document is german
2. create article json document to store it in the s3-bucket
3. store document in s3 bucket
4. if article is in german translate it to english and store it
