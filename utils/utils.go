package utils

import (
	"log"
	"main/models"
)

func Error(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetScore(model models.DataRes) models.ResultData {
	pGood := (model.Classification.Teardown.Good * 100) / model.NumberOfSentences
	pBad := (model.Classification.Teardown.Bad * 100) / model.NumberOfSentences
	return models.ResultData{
		Negative:            model.Sentiment.Verdict,
		GoodBadVerdict:      model.Classification.Verdict,
		PercentGood:         pGood,
		PercentBad:          pBad,
		FlaggedKeywordTotal: model.StaticKeyword.Teardown.Total,
		TotalNumOfSentances: model.NumberOfSentences,
	}
}
