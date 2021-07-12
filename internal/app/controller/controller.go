package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paramonies/ozon-link-shortener/internal/app/service"
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/short", c.getShortLink)
	router.GET("/long", c.getLongLink)
	return router
}

func (c *Controller) getShortLink(ctx *gin.Context) {
	SendErrorResponse(ctx, 200, "/short Link")
}

func (c *Controller) getLongLink(ctx *gin.Context) {
	SendErrorResponse(ctx, 200, "/long Link")
}
