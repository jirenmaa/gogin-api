package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jirenmaa/gogin-api/middlewares"

	router "github.com/jirenmaa/gogin-api/router"
	// gindump "github.com/tpkeeper/gin-dump"
)

func setupLogFile() {
	f, _ := os.Create("gogin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogFile()

	server := gin.Default()

	// gindump.Dump()
	server.Use(middlewares.BasicAuth())

	router.VideoRouter(server)

	server.Run(":8080")
}
