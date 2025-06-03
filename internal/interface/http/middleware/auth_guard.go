package middlewares

import (
	"final_project/internal/pkg/enums"
	"final_project/internal/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthGuard(c *gin.Context) {
	var jwt string = c.GetHeader("Authorization")

	if err := helpers.CheckJWT(c.Request.Context(), jwt); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   enums.ErrUnauthorized,
			"message": err.Error(),
		})

		c.Abort()

		return
	}

	JWTSubject := helpers.GetTokenSubject(jwt)

	c.Set("userID", JWTSubject.Id)
	c.Set("device", JWTSubject.Device)

	c.Next()
}
