package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zou2699/learnGin2/router/api/v1"
)

func InitRouter() *gin.Engine {
	/*
	    r.Use(gin.Logger())

	    r.Use(gin.Recovery())
		// env
	    gin.SetMode(setting.RunMode)
	*/
	r := gin.Default()
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}