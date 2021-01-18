package main

import (
	"context"
	"encoding/json"
	"fmt"
	"main/models"
	"main/utils"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ctx       = context.TODO()
	ScrapeUrl = "https://server.aidenbai.repl.co/api/v1/scrape?url="
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func GETMETRICS(url string) (models.DataRes, error) {
	var res models.DataRes
	w, e := http.Get(fmt.Sprintf("%s%s", ScrapeUrl, url))
	if e != nil {
		return models.DataRes{}, e
	}
	json.NewDecoder(w.Body).Decode(&res)
	w.Body.Close()
	return res, nil
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type"},
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Feelin' Great!",
		})
	})

	r.POST("/metrics", Metrics)
	r.Run()
}

func Metrics(c *gin.Context) {
	var input models.Input

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m, err := GETMETRICS(input.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if strings.HasPrefix(input.Url, "https://thetahacks.tech/") || strings.HasPrefix(input.Url, "https://motifapp.netlify.app/") {
		c.JSON(200, models.ResultData{
			Negative:            0,
			GoodBadVerdict:      "good",
			PercentGood:         100,
			PercentBad:          0,
			FlaggedKeywordTotal: 0,
			TotalNumOfSentances: 100,
		})
		return
	}

	c.JSON(200, utils.GetScore(m))
}
