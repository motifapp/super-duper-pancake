package utils

import (
	"log"
	"main/models"
	"strings"
)

func Error(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func GetScore(model models.DataRes) models.ResultData {
	return models.ResultData{
		PercentSus:       float64(model.Metrics.Preliminary.NumOfSusWords) / float64(len(strings.Split(model.Content, " "))) * 100,
		GenderBias:       model.Metrics.Nlp.GenderBias.Verdict,
		NumberOfSusWords: model.Metrics.Preliminary.NumOfSusWords,
		TotalWords:       len(strings.Split(model.Content, " ")),
		SentimentScore:   model.Metrics.Nlp.SentimentData.Comparative,
	}
}
