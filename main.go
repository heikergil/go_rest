package main

import (
  "net/http"
  "github.com/gin-gonic/gin"

  "github.com/heikergil/go_rest/models"
  "github.com/heikergil/go_rest/controllers"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  

  models.ConnectDatabase()


  r.GET("/users", controllers.FindUsers)
  r.POST("/users", controllers.CreateUser)
  r.GET("/users/:id", controllers.FindUser)
  r.PATCH("/users/:id", controllers.UpdateUser)
  r.DELETE("/users/:id", controllers.DeleteUser)

  r.Run()
}


