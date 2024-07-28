package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// OK sends a 200 OK response
func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

// Created sends a 201 Created response
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   data,
	})
}

// Accepted sends a 202 Accepted response
func Accepted(c *gin.Context, data interface{}) {
	c.JSON(http.StatusAccepted, gin.H{
		"status": "success",
		"data":   data,
	})
}

// BadRequest sends a 400 Bad Request response
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status":  "error",
		"message": message,
	})
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"status":  "error",
		"message": message,
	})
}

// NotFound sends a 404 Not Found response
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  "error",
		"message": message,
	})
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  "error",
		"message": message,
	})
}
