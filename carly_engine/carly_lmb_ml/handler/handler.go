package spider_ml

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	pkg "github.com/leonardpahlke/carly_pkg"
	"gopkg.in/neurosnap/sentences.v1/english"
)

/*
	1. split article text into sentences to analyze sentiment by sentence (neurosnap/sentences)
	2. analyze entities, key phrases and sentiment 					      (comprehend)
	3. store document in s3 bucket 									      (s3)
*/

/* Article Bucket Analytics Files:
FOLDER:<newspaper>
	FOLDER:<ARTICLE-REFERENCE>
		FILE:<LANGUAGE-CODE> (de or en)
		FILE:comprehend
		FILE:DOM
*/

func Handler(request pkg.CarlyEngineLmbMLEvent) (pkg.CarlyEngineLmbMLResponse, error) {
	rfc5646English := pkg.RFC_5646_ENGLISH

	spiderName, _ := pkg.CheckEnvNotEmpty(pkg.EnvSpiderName)
	bucketNameAnalytics, _ := pkg.CheckEnvNotEmpty(pkg.EnvArticleBucketAnalytics)
	mySession := session.Must(session.NewSession())

	/*
		1. split article text into sentences to analyze sentiment by sentence
	*/

	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}

	splittedTextBySentence := tokenizer.Tokenize(request.ArticleText)

	// Create a Comprehend client - to analyse sentiment, entities, key phrases,
	clientComprehend := comprehend.New(mySession)

	/*
		2. --- analyze entities, key phrases and sentiment
	*/

	// key phrases
	detectedKeyPhrases, err := clientComprehend.DetectKeyPhrases(&comprehend.DetectKeyPhrasesInput{
		LanguageCode: &rfc5646English,
		Text:         &request.ArticleText,
	})
	if err != nil {
		pkg.LogError(spiderName, "clientComprehend.DetectKeyPhrases error", err)
	}
	pkg.LogInfo(spiderName, "detectedKeyPhrases complete..")

	// entities
	detectedEntites, err := clientComprehend.DetectEntities(&comprehend.DetectEntitiesInput{
		LanguageCode: &rfc5646English,
		Text:         &request.ArticleText, //todo - detectedEntites
	})
	if err != nil {
		pkg.LogError(spiderName, "clientComprehend.DetectEntities error", err)
	}
	pkg.LogInfo(spiderName, "detectedEntites complete..")

	// sentiment
	var sentimentBySentence []pkg.BucketAnalytics_COMPREHEND_sentiment
	for _, sentence := range splittedTextBySentence {
		detectedSentiment, err := clientComprehend.DetectSentiment(&comprehend.DetectSentimentInput{
			LanguageCode: &rfc5646English,
			Text:         &sentence.Text,
		})
		if err != nil {
			pkg.LogError(spiderName, "clientComprehend.DetectSentiment error", err)
		}
		sentimentBySentence = append(sentimentBySentence, convBucketAnalytics_COMPREHEND_sentiment(sentence.Text, detectedSentiment))
	}
	pkg.LogInfo(spiderName, fmt.Sprintf("detectedSentiment complete.. %v", sentimentBySentence))

	/*
		2. --- store document in s3 bucket
	*/

	spiderMLComprehendDocumentJsonByteArray := pkg.MarshalStruct(pkg.BucketAnalytics_COMPREHEND{
		KeyPhrases:       convBucketAnalytics_COMPREHEND_keyphrases(detectedKeyPhrases),
		Entities:         convBucketAnalytics_COMPREHEND_entities(detectedEntites),
		Sentiment:        sentimentBySentence,
		ArticleReference: request.ArticleReference,
	})

	// Create a S3 client - to store the article text in the s3 bucket
	uploader := s3manager.NewUploader(mySession)
	documentS3Key := pkg.GetBucketKeyForAnalyticsBucket(request.Newspaper, request.ArticleReference, "comprehend", "json")

	// store article comprehend analytics results in s3 bucket
	articleUploadResponse, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucketNameAnalytics,
		Key:    &documentS3Key,
		Body:   strings.NewReader(string(spiderMLComprehendDocumentJsonByteArray)),
	})
	if err != nil {
		pkg.LogError(spiderName, "s3 upload error", err)
	}

	return pkg.CarlyEngineLmbMLResponse{
		ArticleReference: request.ArticleReference,
		Newspaper:        request.Newspaper,
		S3ArticleFileUrl: articleUploadResponse.Location,
	}, nil
}

// convert comprehend sentiment analysis output to internal format
func convBucketAnalytics_COMPREHEND_sentiment(sentence string, detectedSentimentOut *comprehend.DetectSentimentOutput) pkg.BucketAnalytics_COMPREHEND_sentiment {
	return pkg.BucketAnalytics_COMPREHEND_sentiment{
		Sentence:  sentence,
		Sentiment: *detectedSentimentOut.Sentiment,
		SentimentScore: pkg.BucketAnalytics_COMPREHEND_sentiment_scoredetails{
			Mixed:    *detectedSentimentOut.SentimentScore.Mixed,
			Negative: *detectedSentimentOut.SentimentScore.Negative,
			Neutral:  *detectedSentimentOut.SentimentScore.Neutral,
			Positive: *detectedSentimentOut.SentimentScore.Positive,
		},
	}
}

// convert comprehend key phrases analysis output to internal format
func convBucketAnalytics_COMPREHEND_keyphrases(detectedKeyPhrasesOut *comprehend.DetectKeyPhrasesOutput) []pkg.BucketAnalytics_COMPREHEND_key_phrases {
	var keyPhrases []pkg.BucketAnalytics_COMPREHEND_key_phrases
	for _, keyPhrase := range detectedKeyPhrasesOut.KeyPhrases {
		keyPhrases = append(keyPhrases, pkg.BucketAnalytics_COMPREHEND_key_phrases{
			BeginOffset: keyPhrase.BeginOffset,
			EndOffset:   keyPhrase.EndOffset,
			Score:       keyPhrase.Score,
			Text:        keyPhrase.Text,
		})
	}
	return keyPhrases
}

// convert comprehend entities analysis output to internal format
func convBucketAnalytics_COMPREHEND_entities(detectedEntitiesOut *comprehend.DetectEntitiesOutput) []pkg.BucketAnalytics_COMPREHEND_entities {
	var entities []pkg.BucketAnalytics_COMPREHEND_entities
	for _, entity := range detectedEntitiesOut.Entities {
		entities = append(entities, pkg.BucketAnalytics_COMPREHEND_entities{
			BeginOffset: entity.BeginOffset,
			EndOffset:   entity.EndOffset,
			Score:       entity.Score,
			Text:        entity.Text,
			Type:        entity.Type,
		})
	}
	return entities
}
