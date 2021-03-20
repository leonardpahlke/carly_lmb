# install all dependencies
go get -v all

# build code to target directory -> ./build
GOOS=linux go build -o ./build/handler

# zip code to target directory -> ./build
zip -jrm ./build/carly_lmb_translator.zip ./build/handler

