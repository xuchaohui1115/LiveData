package routes

import (
	"LiveData/controller"
	"LiveData/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// InitAnchorRoutes 注册主播路由
func InitAnchorRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	anchoorController := controller.NewAnchorController()
	router := r.Group("/anchor")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/list", anchoorController.GetAnchor)
	}
	return r
}
