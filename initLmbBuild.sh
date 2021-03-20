CARLY_ENGINE_PATH="carly_engine"
BUILD_FILE_NAME="build.sh"

# downloader
CARLY_LMB_DOWNLOADER_NAME="carly_lmb_downloader"
BUILD_CARLY_LMB_DOWNLOADER=1
# parser
CARLY_LMB_PARSER_NAME="carly_lmb_parser"
BUILD_CARLY_LMB_PARSER=1
# ml
CARLY_LMB_ML_NAME="carly_lmb_ml"
BUILD_CARLY_LMB_ML=1
# translator
CARLY_LMB_TRANSLATOR_NAME="carly_lmb_translator"
BUILD_CARLY_LMB_TRANSLATOR=1

echo ""
echo "Initiating Build"
echo ""

# carly lmb downloader
if [ $BUILD_CARLY_LMB_DOWNLOADER -eq 1 ]
then
    echo "Started build ${CARLY_LMB_DOWNLOADER_NAME}"
    cd ./$CARLY_ENGINE_PATH/$CARLY_LMB_DOWNLOADER_NAME
    ./$BUILD_FILE_NAME
    cd ..
    cd ..
else 
    echo "Skip ${CARLY_LMB_DOWNLOADER_NAME} build"
fi

# carly lmb parser
if [ $BUILD_CARLY_LMB_PARSER -eq 1 ]
then
    echo "Started build ${CARLY_LMB_PARSER_NAME}"
    cd ./$CARLY_ENGINE_PATH/$CARLY_LMB_PARSER_NAME
    ./$BUILD_FILE_NAME
    cd ..
    cd ..
else 
    echo "Skip ${CARLY_LMB_PARSER_NAME} build"
fi

# carly lmb translator
if [ $BUILD_CARLY_LMB_TRANSLATOR -eq 1 ]
then
    echo "Started build ${CARLY_LMB_TRANSLATOR_NAME}"
    cd ./$CARLY_ENGINE_PATH/$CARLY_LMB_TRANSLATOR_NAME
    ./$BUILD_FILE_NAME
    cd ..
    cd ..
else 
    echo "Skip ${CARLY_LMB_TRANSLATOR_NAME} build"
fi

# carly lmb ml
if [ $BUILD_CARLY_LMB_ML -eq 1 ]
then
    echo "Started build ${CARLY_LMB_ML_NAME}"
    cd ./$CARLY_ENGINE_PATH/$CARLY_LMB_ML_NAME
    ./$BUILD_FILE_NAME
    cd ..
    cd ..
else 
    echo "Skip ${CARLY_LMB_ML_NAME} build"
fi


echo ""
echo "Finished Build"
echo ""