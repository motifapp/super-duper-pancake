package models

type DataRes struct {
	Metrics struct {
		NumOfSusWords int `json:"numberOfSusWords"`
		NumOfSusChars int `json:"numberOfSusCharacters"`
		TotalChars    int `json:"totalNumberOfCharacters"`
		PercentSus    int `json:"percentageOfSusToTotal"`
	} `json:"metrics"`
	Content string `json:"content"`
}
