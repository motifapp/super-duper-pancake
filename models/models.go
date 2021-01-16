package models

type DataRes struct {
	Metrics struct {
		Preliminary struct {
			NumOfSusWords int     `json:"numberOfSusWords"`
			NumOfSusChars int     `json:"numberOfSusCharacters"`
			TotalChars    int     `json:"totalNumberOfCharacters"`
			PercentSus    float32 `json:"percentageOfSusToTotal"`
		} `json:"preliminary"`
		Nlp struct {
			SentimentData struct {
				Score       float32 `json:"score"`
				Comparative float32 `json:"comparative"`
			} `json:"sentimentData"`
			GenderBias string `json:"genderBias"`
		} `json:"nlp"`
	} `json:"metrics"`
	Content string `json:"content"`
}
