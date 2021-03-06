/*
config 配置
middleware 应用中间件
model 数据模型
utils
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gin-blog/model"
	"gin-blog/pkg/setting"
	"gin-blog/router"
)

func main() {
	r := router.InitRouter()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", setting.Server.Port),
		Handler:        r,
		ReadTimeout:    time.Duration(setting.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe Error:", err.Error())
	}
	defer model.CloseDB()
}
