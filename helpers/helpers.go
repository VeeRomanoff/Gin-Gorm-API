package helpers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Message struct {
	StatusCode int         `json:"status_code"`
	Meta       interface{} `json:"meta"`
	Data       interface{} `json:"data"`
}

func ResponseJSON(ctx *gin.Context, statusCode int, data interface{}) {
	log.Println("status code :", statusCode)
	var message Message
	message.StatusCode = statusCode
	message.Meta = time.Now()
	message.Data = data

	ctx.JSON(http.StatusOK, message)
}
