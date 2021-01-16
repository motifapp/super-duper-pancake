package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alexsasharegan/dotenv"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx  = context.TODO()
	err  = dotenv.Load()
	C, e = mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://motifapp:%s@cluster0.unmcm.mongodb.net/motif?retryWrites=true&w=majority", os.Getenv("PASSWORD")),
	))
	Client = C.Database("motif")
)

func init() {
	if err != nil {
		panic(err)
	}
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	if e != nil {
		panic(e)
	}
	if e := C.Ping(ctx, nil); e != nil {
		panic(e)
	}
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	r.Run()
}
