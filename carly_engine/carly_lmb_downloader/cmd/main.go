package main

import (
	lmb_downloader "carly_lmb_downloader/handler"
	"fmt"
	"os"
	"strconv"

	pkg "github.com/leonardpahlke/carly_pkg"
	log "github.com/sirupsen/logrus"
)

const NewspaperNameTAZ = "taz"
const NewspaperNameZeitOnline = "zeitonline"
const NewspaperNameFrankfurterRundschau = "frankfurterrundschau"

const name = "LocalTest" + pkg.EnvSpiderName

func main() {
	_ = os.Setenv("AWS_REGION", pkg.AWSDeployRegion)
	_ = os.Setenv(pkg.EnvSpiderName, "SpiderDownloader")
	_ = os.Setenv(pkg.EnvArticleBucket, "carly-dev-bucket-article-dom-store")
	_ = os.Setenv(pkg.EnvLogLevel, strconv.Itoa(int(log.InfoLevel)))

	err := os.Mkdir("tmp/", 0755)
	if err == nil {
		pkg.LogInfo(name, "folder created")
	}

	ok := reqSpiderDownloader(pkg.CarlyEngineLmbDownloaderEvent{
		ArticleReference: fmt.Sprintf("%s-5751873", NewspaperNameTAZ),
		ArticleUrl:       "https://taz.de/Neue-All-Parteien-Koalition-in-Italien/!5751873/",
		Newspaper:        NewspaperNameTAZ,
	}, NewspaperNameTAZ)
	if !ok {
		return
	}

	ok = reqSpiderDownloader(pkg.CarlyEngineLmbDownloaderEvent{
		ArticleReference: fmt.Sprintf("%s-202102-5751873", NewspaperNameZeitOnline),
		ArticleUrl:       "https://www.zeit.de/wirtschaft/2021-02/umweltverschmutzung-shell-nigeria-bauern-klage-oel-leck",
		Newspaper:        NewspaperNameZeitOnline,
	}, NewspaperNameTAZ)
	if !ok {
		return
	}

	ok = reqSpiderDownloader(pkg.CarlyEngineLmbDownloaderEvent{
		ArticleReference: fmt.Sprintf("%s-202102-90200529", NewspaperNameFrankfurterRundschau),
		ArticleUrl:       "https://www.fr.de/politik/corona-grenze-kontrolle-seehofer-oesterreich-tirol-tschechien-bayern-sachsen-seehofer-soeder-merkel-eu-zr-90200529.html",
		Newspaper:        NewspaperNameFrankfurterRundschau,
	}, NewspaperNameTAZ)
	if !ok {
		return
	}
}

func reqSpiderDownloader(event pkg.CarlyEngineLmbDownloaderEvent, newspaper string) bool {
	resp, err := lmb_downloader.Handler(event)
	if err != nil {
		pkg.LogError(name, fmt.Sprintf("%s-downloader processing error", newspaper), err)
		return false
	}
	pkg.LogInfo(name, fmt.Sprintf("%s-downloader resp: %s", newspaper, resp))
	return true
}
