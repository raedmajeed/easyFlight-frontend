package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raedmajeed/frontend/config"
	"github.com/raedmajeed/frontend/middleware"
	"github.com/raedmajeed/frontend/pkg/api/handlers"
	"net/http"
)

type FrontendHandler struct {
	ctx *gin.Engine
	cf  *config.Parameters
}

func NewFrontendRoutes(ctx *gin.Engine, cf *config.Parameters) {
	handler := FrontendHandler{
		ctx: ctx,
		cf:  cf,
	}
	apiVersion := ctx.Group("/api/v1")
	{
		frontendBooking := apiVersion.Group("/booking/confirm/online/payment")
		{
			frontendBooking.GET("", handler.OnlinePaymentRender)
			frontendBooking.GET("/success", handler.OnlinePaymentSuccess)
			frontendBooking.GET("/success/render", handler.OnlinePaymentSuccessRender)
		}
	}
}

func (handler *FrontendHandler) UserAuthenticate(ctx *gin.Context) {
	email, err := middleware.ValidateToken(ctx, *handler.cf, "user")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":  err.Error(),
			"status": http.StatusUnauthorized,
		})
		return
	}
	ctx.Set("registered_email", email)
	ctx.Next()
}

func (handler *FrontendHandler) OnlinePaymentRender(ctx *gin.Context) {
	handlers.OnlinePaymentRender(ctx, handler.cf)
}

func (handler *FrontendHandler) OnlinePaymentSuccess(ctx *gin.Context) {
	handlers.OnlinePaymentSuccess(ctx, handler.cf)
}

func (handler *FrontendHandler) OnlinePaymentSuccessRender(ctx *gin.Context) {
	handlers.OnlinePaymentSuccessRender(ctx, handler.cf)
}
