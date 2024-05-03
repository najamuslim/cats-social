package handler

import "github.com/gin-gonic/gin"

type CatHandlerInterface interface {
	GetCatById(c *gin.Context)
	AddCat(c *gin.Context)
}

type AuthHandlerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
