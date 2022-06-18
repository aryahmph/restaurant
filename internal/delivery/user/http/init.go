package http

import (
	httpCommon "github.com/aryahmph/restaurant/common/http"
	jwtCommon "github.com/aryahmph/restaurant/common/jwt"
	domainUser "github.com/aryahmph/restaurant/internal/domain/user"
	userUc "github.com/aryahmph/restaurant/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPUserDelivery struct {
	userUsecase userUc.Usecase
}

func NewHTTPUserDelivery(router *gin.RouterGroup, userUsecase userUc.Usecase, jwtManager *jwtCommon.JWTManager) HTTPUserDelivery {
	router.Use(httpCommon.MiddlewareJWT(jwtManager))
	handler := HTTPUserDelivery{userUsecase: userUsecase}

	router.POST("", handler.register)
	router.GET("", handler.list)
	router.GET("/:id", handler.get)
	router.PUT("/:id", handler.update)
	router.DELETE("/:id", handler.delete)

	return handler
}

func (h HTTPUserDelivery) register(c *gin.Context) {
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
	var requestBody httpCommon.AddUser
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	id, err := h.userUsecase.Register(ctx, h.mapUserBodyToDomain(requestBody), adminID)
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

func (h HTTPUserDelivery) list(c *gin.Context) {
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
	users, err := h.userUsecase.List(ctx, adminID)
	if err != nil {
		c.Error(err)
		return
	}

	var data []httpCommon.User
	for _, user := range users {
		data = append(data, h.mapUserDomainToResponse(user))
	}

	c.PureJSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (h HTTPUserDelivery) get(c *gin.Context) {
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
	id := c.Param("id")

	user, err := h.userUsecase.GetByID(ctx, id, adminID)
	if err != nil {
		c.Error(err)
		return
	}

	c.PureJSON(http.StatusOK, gin.H{
		"data": h.mapUserDomainToResponse(user),
	})
}

func (h HTTPUserDelivery) update(c *gin.Context) {
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
	var requestBody httpCommon.UpdateUser
	requestBody.ID = c.Param("id")
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	user := domainUser.User{
		ID:       requestBody.ID,
		Username: requestBody.Username,
	}
	user.SetUserRoleString(requestBody.Role)

	id, err := h.userUsecase.Update(ctx, user, adminID)
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

func (h HTTPUserDelivery) delete(c *gin.Context) {
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

	id, err := h.userUsecase.Delete(ctx, request.ID, adminID)
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
