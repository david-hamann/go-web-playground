package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	SetupRouter(router)
	router.Run(":8080")

}
