package main

import (
	"github.com/gin-gonic/gin"
	//"fmt"
	"net/http"
)

func GetData(c *gin.Context) {
	//	c.String(http.StatusOK, "foo")
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}

func PutData(c *gin.Context) {

	//Sends a JSON response with no content
	//c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	//c.Writer.WriteHeader(http. StatusNoContent)
	c.JSON(http.StatusNoContent, nil)
}

func PostData(c *gin.Context) {
	//Sends a JSON response with no content
	//c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	//c.Writer.WriteHeader(http. StatusNoContent)
	c.JSON(http.StatusNoContent, nil)
}
