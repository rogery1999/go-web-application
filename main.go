package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Salute struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("Hello There!")

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		log.Println("New Change")
		ctx.JSON(http.StatusOK, Salute{Message: "Hello World!"})
	})

	r.Run("127.0.0.1:8000")
}
