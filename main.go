package main

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx    = context.TODO()
	C, e   = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://motifapp:<password>@cluster0.unmcm.mongodb.net/motif?retryWrites=true&w=majority"))
	Client = C.Database("motif")
)

func init() {
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
