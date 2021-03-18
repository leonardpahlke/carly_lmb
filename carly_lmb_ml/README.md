# Carly Lambda ML

This repository contains lambda code for the lambda function carly-lmb-ml, which is described in the repository carly_aws.
As part of the carly-engine processing AWS StepFunction, this lambda function processes article data.
Specifically, it uses AWS Comprehend NLP to analyse an article text.

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

1. Lambda function receives event and starts lambda handler with article information
2. create comprehend client and analyze article-text
    1. entities
    2. key phrases
    3. sentiment by sentence
3. create JSON document with analytic results
4. store JSON document in s3 bucket
