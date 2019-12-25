package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/e"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/util"
	"github.com/spf13/viper"
)

func ParseRequest(c *gin.Context) error {
	header := c.Request.Header.Get("Authorization")
	secret := viper.GetString("jwt_secret")
	
	if len(header) == 0 {
		return e.ErrMissingHeader
	}
	
	var token string
	_, _ = fmt.Sscanf(header, "Bearer %s", &token)
	return util.ParseToken(token, secret)
}
