package lmb_parser

import (
	"errors"

	pkg "github.com/leonardpahlke/carly_pkg"
	html "golang.org/x/net/html"
)

// the scraper get the input from the S3
/*
	1. Lambda function receives event and starts lambda handler with article dom information
	2. swtich article newspaper cases to apply parsing information
	3. parse dom with parsing config
*/

/*
	1. Lambda function receives event and starts lambda handler with article dom information
*/
func Handler(request pkg.CarlyEngineLmbParserEvent) (pkg.CarlyEngineLmbParserResponse, error) {
	spiderName, _ := pkg.CheckEnvNotEmpty(pkg.EnvSpiderName)
	pkg.SetLogLevel()

	articleDomElement := ""
	var articleDomElementValues []html.Attribute
	var textTagsToParse []string
	var whitelistAttributes []html.Attribute

	/*
		2. swtich article newspaper cases to apply parsing information
	*/

	switch request.Newspaper {

	// Taz parsing configuration
	case pkg.NewspaperNameTAZ:
		articleDomElement = "article"
		textTagsToParse = []string{
			"a", "p", "em", "string", "blockquote", "q", "cite", "h1", "h2", "h3", "h4", "h5", "h6", "span", "strong",
		}
		whitelistAttributes = []html.Attribute{
			{
				Key: "class",
				Val: "hide",
			},
			{
				Key: "class",
				Val: "credit",
			},
			{
				Key: "class",
				Val: "caption",
			},
		}
		articleDomElementValues = []html.Attribute{
			{
				Key: "class",
				Val: "sectbody",
			},
		}

	// ZeitOnline parsing configuration
	case pkg.NewspaperNameZeitOnline:
		articleDomElement = "article"
		textTagsToParse = []string{
			"a", "p", "em", "string", "blockquote", "q", "cite", "h1", "h2", "h3", "h4", "h5", "h6", "span", "strong",
		}
		whitelistAttributes = []html.Attribute{}
		articleDomElementValues = []html.Attribute{
			{
				Key: "class",
				Val: "article article--padded article--article",
			},
		}

	// FrankfurterRundschau parsing configuration
	case pkg.NewspaperNameFrankfurterRundschau:
		articleDomElement = "div"
		textTagsToParse = []string{
			"a", "p", "em", "string", "blockquote", "q", "cite", "h1", "h2", "h3", "h4", "h5", "h6", "span", "strong",
		}
		whitelistAttributes = []html.Attribute{
			{
				Key: "class",
				Val: "id-AuthorList id-Article-content-item ",
			},
		}
		articleDomElementValues = []html.Attribute{{
			Key: "class",
			Val: "id-Article-body lp_article_content",
		}}
	default:
		pkg.LogError(
			spiderName,
			"The requested newspaper reference could not get resolved",
			errors.New("request.Newspaper unrecognized"))
	}

	// decrease DOM to article (remove everything else besides the article dom elements)
	elementDomStr, err := parseDomElement(request.ArticleDom, articleDomElement, articleDomElementValues)
	if err != nil {
		pkg.LogError(spiderName, "pkg.ParseDomElement error", err)
	}

	/*
		3. parse dom with parsing config
	*/
	articleParseResponse, err := parseArticle(articleParseRequest{
		ArticleHtmlDom:           elementDomStr,
		TextTagsToParse:          textTagsToParse,
		WhitelistAttributeValues: whitelistAttributes,
	})
	if err != nil {
		pkg.LogError(spiderName, "pkg.ParseArticle error", err)
	}

	updatedArticleText := bundleSentences(articleParseResponse.ArticleText)

	return pkg.CarlyEngineLmbParserResponse{
		ArticleReference: request.ArticleReference,
		S3ArticleDomLink: request.S3ArticleDomLink,
		Newspaper:        request.Newspaper,
		ArticleText:      updatedArticleText,
	}, nil
}
