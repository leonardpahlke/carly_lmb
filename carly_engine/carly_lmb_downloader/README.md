# Carly Lambda Downloader

This repository contains lambda code for the lambda function carly-lmb-downloader, which is described in the repository carly_aws.
As part of the carly-engine processing AWS StepFunction, this lambda function processes article data.
Specifically, it downloads a website dom and stores the .html file in s3.

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

1. set up lambda clients
2. create file to store article html
3. retrieve article html from article website
4. store html to local file
5. upload file to s3
