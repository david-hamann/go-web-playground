package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("data", GetData)
		//v1.GET("data", GetData)
		v1.POST("data", PostData)
		v1.PUT("data", PutData)
	}

}
