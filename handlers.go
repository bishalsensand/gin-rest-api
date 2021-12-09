package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//var sensandEngineers [9]string = [9]string{"Mukhtar", "Sergey", "Chris", "Zac", "CK", "Bishal", "Gary", "UV", "Julian"}

func isSensandEngineer(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    true,
	})
}
