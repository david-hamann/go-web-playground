package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)


func GetData(c *gin.Context) {
    c.JSON(http.StatusOK, "foo"
}


/*
func PostData(c *gin.Context) {
}

func PutData(c *gin.Context) {
}
*/

        //v1.GET("data", controllers.getData)
        //v1.POST("data", contollers.postData)
        //v1.PUT("data", contollers.putData)


