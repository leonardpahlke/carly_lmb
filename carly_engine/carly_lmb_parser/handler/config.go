package lmb_parser

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	pkg "github.com/leonardpahlke/carly_pkg"
	"github.com/prometheus/common/log"
	html "golang.org/x/net/html"
)

// handler is a simple function that takes a string and does a ToUpper.
func parseArticle(request articleParseRequest) (articleParseResponse, error) {
	methodName := "pkg.ParserArticle"

	tag := ""
	enter := false

	article := ""

	tokenizer := html.NewTokenizer(strings.NewReader(request.ArticleHtmlDom))
	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		switch tokenType {
		case html.ErrorToken:
			pkg.LogError(methodName, "Token error", err)

		case html.StartTagToken, html.SelfClosingTagToken:
			enter = false
			tagAttribute := token.Attr
			tag = token.Data
			for _, textTag := range request.TextTagsToParse {
				if tag == textTag && !checkElementAttribute(request.WhitelistAttributeValues, tagAttribute) {
					enter = true
					break
				}
			}

		case html.TextToken:
			if enter {
				data := strings.TrimSpace(token.Data)
				if len(data) > 0 {
					article += fmt.Sprintf("%s\n", data)
				}
			}
		}
	}

	return articleParseResponse{
		ArticleText: article,
	}, nil
}

func parseDomElement(htmlDom string, element string, requiredAttribute []html.Attribute) (string, error) {
	doc, _ := html.Parse(strings.NewReader(htmlDom))
	bn, err := getElement(doc, element, requiredAttribute)
	if err != nil {
		return "", err
	}
	log.Info(bn.Attr)
	return renderNode(bn), nil
}

func getElement(doc *html.Node, element string, requiredAttribute []html.Attribute) (*html.Node, error) {
	var elementNode *html.Node
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == element && checkElementAttribute(requiredAttribute, node.Attr) {
			elementNode = node
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}
	crawler(doc)
	if elementNode != nil {
		return elementNode, nil
	}
	return nil, fmt.Errorf("missing <%s> in the node tree", element)
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	_ = html.Render(w, n)
	return buf.String()
}

func bundleSentences(articleText string) string {
	updatedArticleText := ""
	splitedByDot := strings.Split(articleText, ".")
	for _, sentence := range splitedByDot {
		updatedArticleText += strings.ReplaceAll(sentence, "\n", " ") + ".\n"
	}
	return updatedArticleText
}

func checkElementAttribute(attrToCheck []html.Attribute, elementAttr []html.Attribute) bool {
	for _, checkAttr := range attrToCheck {
		for _, elemAttr := range elementAttr {
			if checkAttr.Key == elemAttr.Key && checkAttr.Val == elemAttr.Val {
				return true
			}
		}
	}
	return false
}

type articleParseRequest struct {
	ArticleHtmlDom           string
	TextTagsToParse          []string
	WhitelistAttributeValues []html.Attribute
}

type articleParseResponse struct {
	ArticleText string
}

/*
type storeFileMetaStruct struct {
	SpiderName       string
	BucketName       string
	ArticleReference string
	Newspaper        string
	Uploader         s3manager.Uploader
}

type storeFileStruct struct {
	Filename   string
	FileEnding string
	File       string
}
*/
