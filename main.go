package main

import (
    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()
    app.Router(r)
    r.run(":8080")

}

