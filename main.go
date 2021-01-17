package main

import (
	"context"
	"encoding/json"
	"fmt"
	"main/models"
	"main/utils"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx       = context.TODO()
	C, e      = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://motifapp:<password>@cluster0.unmcm.mongodb.net/motif?retryWrites=true&w=majority"))
	Client    = C.Database("motif")
	ScrapeUrl = "https://server.aidenbai.repl.co/api/v1/scrape?url="
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	utils.Error(e)
	e = C.Ping(ctx, nil)
	utils.Error(e)
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
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Feelin' Great!",
		})
	})

	r.GET("/metrics", Metrics)
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

	c.JSON(200, utils.GetScore(m))
}
