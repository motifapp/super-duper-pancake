package models

type DataRes struct {
	Sentiment struct {
		Verdict float64 `json:"verdict"`
	} `json:"sentiment"`
	Classification struct {
		Verdict  string `json:"verdict"`
		Teardown struct {
			Good int `json:"good"`
			Bad  int `json:"bad"`
		} `json:"teardown"`
	} `json:"classification"`
	StaticKeyword struct {
		Verdict  float64 `json:"verdict"`
		Teardown struct {
			Total int `json:"total"`
		} `json:"teardown"`
	} `json:"staticKeyword"`
	GenderBias struct {
		Verdict  string `json:"verdict"`
		Teardown struct {
			Male   int `json:"male"`
			Female int `json:"female"`
		} `json:"teardown"`
	} `json:"genderBias"`
	NumberOfSentences int `json:"numberOfSentences"`
}

type Input struct {
	Url string `json:"url" form:"url" binding:"required"`
}

type ResultData struct {
	Negative            float64 `json:"negative"`
	GoodBadVerdict      string  `json:"goodBadVerdict"`
	PercentGood         int     `json:"percentGood"`
	PercentBad          int     `json:"percentBad"`
	FlaggedKeywordTotal int     `json:"flaggedKeywordTotal"`
	TotalNumOfSentances int     `json:"totalNumOfSentaces"`
}
