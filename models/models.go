package models

type DataRes struct {
	Metrics struct {
		Preliminary struct {
			NumOfSusWords int `json:"numberOfSusWords"`
			NumOfSusChars int `json:"numberOfSusCharacters"`
			TotalChars    int `json:"totalNumberOfCharacters"`
		} `json:"preliminary"`
		Nlp struct {
			SentimentData struct {
				Score       float32 `json:"score"`
				Comparative float32 `json:"comparative"`
			} `json:"sentimentData"`
			GenderBias struct {
				Verdict string `json:"verdict"`
				Female  int    `json:"female"`
				Male    int    `json:"male"`
			} `json:"genderBias"`
		} `json:"nlp"`
	} `json:"metrics"`
	Content string `json:"content"`
}

type ResultData struct {
	PercentSus       float64 `json:"percentsus"`
	GenderBias       string  `json:"genderbias"`
	NumberOfSusWords int     `json:"numofsuswords"`
	TotalWords       int     `json:"totalwords"`
	SentimentScore   float32 `json:"sentimentscore"`
}
