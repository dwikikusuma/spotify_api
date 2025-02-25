package memberships

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spotify_api/internal/model/memberhsips"
)

func (h *Handler) SignUp(c *gin.Context) {
	var req memberhsips.SignUpRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	err = h.service.SignUp(req)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}
