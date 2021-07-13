package controller

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/paramonies/ozon-link-shortener/internal/app/model"
	"github.com/paramonies/ozon-link-shortener/internal/app/service"
	"github.com/paramonies/ozon-link-shortener/internal/app/utils"

	_ "github.com/paramonies/ozon-link-shortener/docs"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

type Controller struct {
	service service.Service
}

func NewController(service service.Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/short", c.shortLink)
	router.POST("/long", c.getLongLink)
	return router
}

// @Summary создать или получить сокращенную ссылку
// @Tags ShortLink
// @Description Создать новую или получить существующую сокращенную ссылку
// @ID get-short-link
// @Accept  json
// @Produce  json
// @Param input body InputShortLink true "long http link"
// @Success 200 {object} model.ClientLink
// @Success 201 {object} model.ClientLink
// @Failure 400 {object} GetShortLinkMessage400
// @Failure 500 {object} GetShortLinkMessage500
// @Router /short [post]
func (c *Controller) shortLink(ctx *gin.Context) {
	var input model.ClientLink

	if err := ctx.BindJSON(&input); err != nil {
		SendErrorResponse(ctx, http.StatusBadRequest, "invalid response body")
		return
	}

	_, err := url.ParseRequestURI(input.Url)
	if err != nil {
		SendErrorResponse(ctx, http.StatusBadRequest, "invalid http link format")
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

// @Summary получить полную ссылку
// @Tags LongLink
// @Description Получить полную ссылку по сокращенному id
// @ID get-long-link
// @Accept  json
// @Produce  json
// @Param input body InputLongLink true "short id"
// @Success 200 {object} model.ClientLink
// @Failure 400 {object} GetShortLinkMessage400
// @Failure 500 {object} GetShortLinkMessage500
// @Router /long [post]
func (c *Controller) getLongLink(ctx *gin.Context) {
	var input model.ClientLink

	if err := ctx.BindJSON(&input); err != nil {
		SendErrorResponse(ctx, http.StatusBadRequest, "invalid response body")
		return
	}

	if !utils.IsShortValid(input.Url) {
		SendErrorResponse(ctx, http.StatusBadRequest, "invalid short url id format")
		return
	}

	longUrl, err := c.service.GetLongLink(input.Url)
	if err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"url": longUrl,
	})
}
