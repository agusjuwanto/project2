package route

import (
    "net/http"

    "github.com/example/gin-rest-api/internal/handler"
    "github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, userHandler *handler.UserHandler) {
    r.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "API is healthy",
        })
    })

    api := r.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.GET("", userHandler.GetAll)
            users.GET("/:id", userHandler.GetByID)
            users.POST("", userHandler.Create)
            users.PUT("/:id", userHandler.Update)
            users.DELETE("/:id", userHandler.Delete)
        }
    }
}
