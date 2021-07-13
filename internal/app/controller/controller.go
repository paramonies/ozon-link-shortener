package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paramonies/ozon-link-shortener/internal/app/model"
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

	router.POST("/short", c.shortLink)
	router.GET("/long", c.getLongLink)
	return router
}

func (c *Controller) shortLink(ctx *gin.Context) {
	var input model.ClientLink

	if err := ctx.BindJSON(&input); err != nil {
		SendErrorResponse(ctx, http.StatusBadRequest, "invalid response body")
		return
	}

	shortUrl := c.service.GetShortLink(input.Url)

	if shortUrl != "" {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"url": shortUrl,
		})
		return
	}

	shortLink, err := c.service.CreateLink(input.Url)
	if err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, shortLink)
}

func (c *Controller) getLongLink(ctx *gin.Context) {
	SendErrorResponse(ctx, 200, "/long Link")
}
