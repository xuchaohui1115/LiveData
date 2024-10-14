package controller

import (
	"LiveData/common"
	"LiveData/repository"
	"LiveData/response"
	"LiveData/vo"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// IIllegalController 违规管理
type IIllegalController interface {
	GetIllegal(c *gin.Context) //获取违规列表
}

type IllegalController struct {
	IllegalRepository repository.IIllegalRepository
}

func (a IllegalController) GetIllegal(c *gin.Context) {
	var req vo.IllegalRequest
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
	illegal, total, err := a.IllegalRepository.GetIllegal(&req)
	if err != nil {
		response.Fail(c, nil, "获取违规记录列表失败: "+err.Error())
		return
	}
	response.Success(c, gin.H{"违规记录": illegal, "total": total}, "获取违规记录列表成功")
}

// 构造函数
func NewIllegalController() IIllegalController {
	illegalRepository := repository.NewIllegalRepository()
	illegalController := IllegalController{IllegalRepository: illegalRepository}
	return illegalController
}
