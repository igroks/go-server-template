package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func task(c *gin.Context) {
	var requestMsg MatchRequest

	if err := c.ShouldBindJSON(&requestMsg); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	responseMsg := MatchResponse{
		Message: "Seu request foi recebido! Sua menssagem foi: " + *requestMsg.Query,
	}

	c.JSON(http.StatusOK, &responseMsg)

}
