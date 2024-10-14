package controller

import (
	"LiveData/common"
	"LiveData/repository"
	"LiveData/response"
	"LiveData/vo"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IAnchorController 主播管理
type IAnchorController interface {
	GetAnchor(c *gin.Context) //获取主播列表
}

type AnchorController struct {
	AnchorRepository repository.IAnchorRepository
}

func (a AnchorController) GetAnchor(c *gin.Context) {
	var req vo.AnchorListRequest
	// 参数绑定
	if err := c.ShouldBind(&req); err != nil {
		response.Fail(c, nil, err.Error())
		return
	}
	// 参数校验
	if err := common.Validate.Struct(&req); err != nil {
		errStr := err.(validator.ValidationErrors)[0].Translate(common.Trans)
		response.Fail(c, nil, errStr)
		return
	}

	// 获取
	anchors, total, err := a.AnchorRepository.GetGetAnchor(&req)
	if err != nil {
		response.Fail(c, nil, "获取主播列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"anchor": anchors, "total": total}, "获取主播列表成功")
}

// 构造函数
func NewAnchorController() IAnchorController {
	anchorRepository := repository.NewAnchorRepository()
	anchorController := AnchorController{AnchorRepository: anchorRepository}
	return anchorController
}
