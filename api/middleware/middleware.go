package middleware

import (
	"api-gateway/api/auth"
	"errors"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func Check(c *gin.Context) {

	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})
		return
	}

	_, err := auth.ValidateAccessToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token provided",
		})
		return
	}
	c.Next()
}

func (casb *casbinPermission) GetRole(c *gin.Context) (string, int) {
	token := c.GetHeader("Authorization")
	if token == "" {
		return "unauthorized", http.StatusUnauthorized
	}
	_, role, _, _, _, _, _, err := auth.GetUserInfoFromAccessToken(token)
	if err != nil {
		return "error while reding role", 500
	}

	return role, 0
}

func (casb *casbinPermission) CheckPermission(c *gin.Context) (bool, error) {

	act := c.Request.Method
	sub, status := casb.GetRole(c)
	if status != 0 {
		return false, errors.New("error in get role")
	}
	obj := c.FullPath()

	ok, err := casb.enforcer.Enforce(sub, obj, act)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Internal server error",
		})
		c.Abort()
		return false, err
	}
	return ok, nil
}

func CheckPermissionMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := &casbinPermission{
		enforcer: enf,
	}

	return func(c *gin.Context) {
		result, err := casbHandler.CheckPermission(c)

		if err != nil {
			c.AbortWithError(500, err)
		}
		if !result {
			c.AbortWithStatusJSON(401, gin.H{
				"message": "Forbidden",
			})
		}

		c.Next()
	}
}
