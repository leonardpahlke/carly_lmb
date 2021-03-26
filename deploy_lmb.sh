#!/bin/sh

CARLY_ENGINE_PATH="carly_engine"
S3_LAMBDA_CODE_BUCKET="s3://carly-dev-article-bucket"

# downloader
CARLY_LMB_DOWNLOADER_NAME="carly_lmb_downloader"
DEPLOY_CARLY_LMB_DOWNLOADER=$1
# parser
CARLY_LMB_PARSER_NAME="carly_lmb_parser"
DEPLOY_CARLY_LMB_PARSER=$2
# ml
CARLY_LMB_ML_NAME="carly_lmb_ml"
DEPLOY_CARLY_LMB_ML=$3
# translator
CARLY_LMB_TRANSLATOR_NAME="carly_lmb_translator"
DEPLOY_CARLY_LMB_TRANSLATOR=$4

echo ""
echo "Initiating Deployment"
echo ""

# carly lmb downloader
if [ $DEPLOY_CARLY_LMB_DOWNLOADER -eq 1 ]
then
    echo "... STARTED deployment ${CARLY_LMB_DOWNLOADER_NAME}"
    aws s3 cp "./${CARLY_ENGINE_PATH}/${CARLY_LMB_DOWNLOADER_NAME}/build/${CARLY_LMB_DOWNLOADER_NAME}.zip" "${S3_LAMBDA_CODE_BUCKET}/${CARLY_LMB_DOWNLOADER_NAME}/latest.zip"
else 
    echo "!!! SKIP ${CARLY_LMB_DOWNLOADER_NAME} deployment"
fi

# carly lmb parser
if [ $DEPLOY_CARLY_LMB_PARSER -eq 1 ]
then
    echo "... STARTED deployment ${CARLY_LMB_PARSER_NAME}"
    aws s3 cp "./${CARLY_ENGINE_PATH}/${CARLY_LMB_PARSER_NAME}/build/${CARLY_LMB_PARSER_NAME}.zip" "${S3_LAMBDA_CODE_BUCKET}/${CARLY_LMB_PARSER_NAME}/latest.zip"
else 
    echo "!!! SKIP ${CARLY_LMB_PARSER_NAME} deployment"
fi

# carly lmb translator
if [ $DEPLOY_CARLY_LMB_TRANSLATOR -eq 1 ]
then
    echo "... STARTED deployment ${CARLY_LMB_TRANSLATOR_NAME}"
    aws s3 cp "./${CARLY_ENGINE_PATH}/${CARLY_LMB_TRANSLATOR_NAME}/build/${CARLY_LMB_TRANSLATOR_NAME}.zip" "${S3_LAMBDA_CODE_BUCKET}/${CARLY_LMB_TRANSLATOR_NAME}/latest.zip"
else 
    echo "!!! SKIP ${CARLY_LMB_TRANSLATOR_NAME} deployment"
fi

# carly lmb ml
if [ $DEPLOY_CARLY_LMB_ML -eq 1 ]
then
    echo "... STARTED deployment ${CARLY_LMB_ML_NAME}"
    aws s3 cp "./${CARLY_ENGINE_PATH}/${CARLY_LMB_ML_NAME}/build/${CARLY_LMB_ML_NAME}.zip" "${S3_LAMBDA_CODE_BUCKET}/${CARLY_LMB_ML_NAME}/latest.zip"
else 
    echo "!!! SKIP ${CARLY_LMB_ML_NAME} deployment"
fi


echo ""
echo "Finished Deployment"
echo ""
 