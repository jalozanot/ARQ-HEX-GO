package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jalozanot/demoCeiba/domain/service"
)

const (
	CREATE = 1
	GET = 2
	GETS = 3
	DELETE = 4
	UPDATE = 5)

func Create(c *gin.Context) {
	service.ProxyHandler(CREATE,c)
}

func  Get(c *gin.Context) {
	service.ProxyHandler(GET,c)
}

func  Gets(c *gin.Context) {
	service.ProxyHandler(GETS,c)
}

func  Delete(c *gin.Context) {
	service.ProxyHandler(DELETE,c)
}

func  Update(c *gin.Context) {
	service.ProxyHandler(UPDATE,c)
}


