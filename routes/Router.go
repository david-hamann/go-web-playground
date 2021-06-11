package routes

import (
    "../controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
    
    v1 := r.group("/v1")
    {
        v1.GET("data", controllers.GetData)
        //v1.POST("data", contollers.PostData)
        //v1.PUT("data", contollers.PutData)
    }

}

