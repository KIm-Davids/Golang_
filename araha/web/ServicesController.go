package web

import (
	"Golang_/araha/services"
	"github.com/gin-gonic/gin"
)

func ServicesController(c *gin.Context) {

	services.ProcessPaymentService()
}
