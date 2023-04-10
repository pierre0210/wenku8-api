package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pierre0210/wenku8-api/internal/util"
)

func main() {
	port := flag.Int("p", 5000, "Port")
	flag.Parse()

	err := godotenv.Load()
	util.ErrorHandler(err, true)

	router := gin.Default()

	novelRouter := router.Group("/novel")
	novelRouter.GET("/chapter/:aid/:vol/:ch")
	novelRouter.GET("/index/:aid")

	router.Run(fmt.Sprintf("localhost:%d", *port))
}
