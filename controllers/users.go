package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/heikergil/go_rest/models"
)

type CreateUserInput struct {
  Name  string `json:"name" binding:"required"`
}

type UpdateUserInput struct {
  Name  string `json:"name"` 
}

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
  
	c.JSON(http.StatusOK, gin.H{"data": users})
  }


  // POST /users
// Create new user
func CreateUser(c *gin.Context) {
  // Validate input
  var input CreateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create user
  user := models.User{Name: input.Name}
  models.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{"data": user})
}


// GET /users/:id
// Find a user
func FindUser(c *gin.Context) {  // Get model if exist
  var user models.User

  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": user})
}



// PATCH /users/:id
// Update user info
func UpdateUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
    return
  }

  // Validate input
  var input UpdateUserInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Model(&user).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": user})
}


// DELETE /user/:id
// Delete user
func DeleteUser(c *gin.Context) {
  // Get model if exist
  var user models.User
  if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "user not found!"})
    return
  }

  models.DB.Delete(&user)

  c.JSON(http.StatusOK, gin.H{"data": true})
}