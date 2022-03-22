package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/heikergil/go_rest/models"
)

// GET /books
// Get all books
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
  
	c.JSON(http.StatusOK, gin.H{"data": users})
  }