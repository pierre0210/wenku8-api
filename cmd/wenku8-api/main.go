package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pierre0210/wenku8-api/internal/api"
	"github.com/pierre0210/wenku8-api/internal/database"
)

func main() {
	port := flag.Int("p", 5000, "Port")
	flag.Parse()

	database.InitRedis()

	//err := godotenv.Load()
	//util.ErrorHandler(err, true)

	router := gin.Default()

	novelRouter := router.Group("/novel")
	novelRouter.GET("/chapter/:aid/:vol/:ch", api.HandleGetChapter)
	novelRouter.GET("/volume/:aid/:vol", api.HandleGetVolume)
	novelRouter.GET("/index/:aid", api.HandleGetIndex)

	router.Run(fmt.Sprintf(":%d", *port))
}
