package routes

import (
	"LiveData/controller"
	"LiveData/middleware"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitOperationLogRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	operationLogController := controller.NewOperationLogController()
	router := r.Group("/log")
	// 开启jwt认证中间件
	router.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	router.Use(middleware.CasbinMiddleware())
	{
		router.GET("/operation/list", operationLogController.GetOperationLogs)
		router.DELETE("/operation/delete/batch", operationLogController.BatchDeleteOperationLogByIds)
	}
	return r
}
