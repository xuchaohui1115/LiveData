package routes

import (
	"LiveData/controller"
	"LiveData/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// InitIllegalRoutes 注册违规记录路由
func InitIllegalRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	illegalController := controller.NewIllegalController()
	router := r.Group("/illegal")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/list", illegalController.GetIllegal)
	}
	return r
}
