package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/models"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/setting"
	"github.com/lawtech0902/gin_todo_demo/backend/routers"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/unknwon/com"
	"log"
	"net/http"
	"time"
)

var cfg = pflag.StringP("config", "c", "", "api server config file path.")

func init() {
	pflag.Parse()
	if err := setting.Init(*cfg); err != nil {
		panic(err)
	}
	log.Printf("config init successful!")
	
	models.Init()
	log.Printf("db init successful!")
}

func main() {
	gin.SetMode(viper.GetString("runmode"))
	log.Printf("gin run mode set to: %s", viper.GetString("runmode"))
	
	routersInit := routers.InitRouter()
	readTimeout := time.Duration(com.StrTo(viper.GetString("read_timeout")).MustInt()) * time.Second
	writeTimeout := time.Duration(com.StrTo(viper.GetString("write_timeout")).MustInt()) * time.Second
	
	server := &http.Server{
		Addr:           viper.GetString("addr"),
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	
	log.Printf("Start to listening on: %s", viper.GetString("addr"))
	
	_ = server.ListenAndServe()
}
