package http

import (
	errorCommon "github.com/aryahmph/restaurant/common/error"
	"github.com/aryahmph/restaurant/common/jwt"
	"github.com/gin-gonic/gin"
)

func MiddlewareJWT(j *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= BEARER {
			c.Error(errorCommon.NewInvariantError("authorization header not valid"))
			c.Abort()
			return
		}
		tokenString := authHeader[BEARER:]
		id, err := j.VerifyToken(tokenString)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}

		c.Set("user_id", id)
		c.Next()
	}
}
