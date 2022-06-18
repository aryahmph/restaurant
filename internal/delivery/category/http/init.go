package http

import (
	httpCommon "github.com/aryahmph/restaurant/common/http"
	jwtCommon "github.com/aryahmph/restaurant/common/jwt"
	domainCategory "github.com/aryahmph/restaurant/internal/domain/category"
	categoryUc "github.com/aryahmph/restaurant/internal/usecase/category"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPCategoryDelivery struct {
	categoryUsecase categoryUc.Usecase
}

func NewHTTPCategoryDelivery(router *gin.RouterGroup, categoryUsecase categoryUc.Usecase, jwtManager *jwtCommon.JWTManager) HTTPCategoryDelivery {
	handler := HTTPCategoryDelivery{categoryUsecase: categoryUsecase}

	router.GET("", handler.list)
	//router.GET("/:id", handler.get)

	router.Use(httpCommon.MiddlewareJWT(jwtManager))
	router.POST("", handler.create)
	router.PUT("/:id", handler.update)
	router.DELETE("/:id", handler.delete)

	return handler
}

func (h HTTPCategoryDelivery) list(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := h.categoryUsecase.List(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	var data []httpCommon.Category
	for _, category := range categories {
		data = append(data, h.mapCategoryDomainToResponse(category))
	}

	c.PureJSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h HTTPCategoryDelivery) create(c *gin.Context) {
	inputUserID, ok := c.Get("user_id")
	if !ok {
		c.Error(ErrorUserID)
		return
	}
	adminID, ok := inputUserID.(string)
	if !ok {
		c.Error(ErrorUserID)
		return
	}

	ctx := c.Request.Context()
	var requestBody httpCommon.Category
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	id, err := h.categoryUsecase.Create(ctx, h.mapCategoryBodyToDomain(requestBody), adminID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}

func (h HTTPCategoryDelivery) update(c *gin.Context) {
	inputUserID, ok := c.Get("user_id")
	if !ok {
		c.Error(ErrorUserID)
		return
	}
	adminID, ok := inputUserID.(string)
	if !ok {
		c.Error(ErrorUserID)
		return
	}

	ctx := c.Request.Context()
	var requestBody httpCommon.UpdateCategory
	requestBody.ID = c.Param("id")
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	category := domainCategory.Category{ID: requestBody.ID, Name: requestBody.Name}
	id, err := h.categoryUsecase.Update(ctx, category, adminID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}

func (h HTTPCategoryDelivery) delete(c *gin.Context) {
	inputUserID, ok := c.Get("user_id")
	if !ok {
		c.Error(ErrorUserID)
		return
	}
	adminID, ok := inputUserID.(string)
	if !ok {
		c.Error(ErrorUserID)
		return
	}

	ctx := c.Request.Context()
	var request httpCommon.Delete
	if err := c.BindUri(&request); err != nil {
		return
	}

	id, err := h.categoryUsecase.Delete(ctx, request.ID, adminID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}
