package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/app"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/e"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/util"
	"github.com/spf13/viper"
	"net/http"
)

type Key struct {
	Key string `json:"key"`
}

func Token(c *gin.Context) {
	var key Key
	
	c.BindJSON(&key)
	if key.Key != viper.GetString("key") {
		app.SendResponse(c, http.StatusBadRequest, e.ErrKeyIncorrect, nil)
		return
	}
	
	t, err := util.GenerateToken(key.Key, "")
	if err != nil {
		app.SendResponse(c, http.StatusInternalServerError, e.ErrToken, nil)
		return
	}
	
	app.SendResponse(c, http.StatusOK, nil, util.Token{Token: t})
}
