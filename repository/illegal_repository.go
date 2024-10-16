package repository

import (
	"LiveData/common"
	"LiveData/model"
	"LiveData/vo"
	"fmt"
	"strings"
)

type IIllegalRepository interface {
	GetIllegal(v *vo.IllegalRequest) ([]*model.Violations, int64, error) // 获取违规列表
}

type IllegalRepository struct {
}

// NewIllegalRepository 构造函数
func NewIllegalRepository() IIllegalRepository {
	return IllegalRepository{}
}

// 获取主播列表
func (ar IllegalRepository) GetIllegal(req *vo.IllegalRequest) ([]*model.Violations, int64, error) {
	var list []*model.Violations
	db := common.DB.Model(&model.Violations{}).Order("created_at DESC")
	/*
		Nickname      string `json:"nickname" form:"nickname"` 主播昵称/UID
		IllegalNum    string `json:"illegalNum" form:"illegalNum"`
		IllegalType   string `json:"illegalType" form:"illegalType"`
		IllegalEffect string `json:"illegalEffect" form:"illegalEffect"`  //违规影响
		McnName       string `json:"mcnName" form:"mcnName"`
		Status        int    `json:"status" form:"status"`
	*/
	nickname := strings.TrimSpace(req.Nickname)
	if nickname != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", nickname)).Or("uid LIKE ?", fmt.Sprintf("%%%s%%", nickname))
	}
	illegalNum := strings.TrimSpace(req.IllegalNum)
	if illegalNum != "" {
		db = db.Where("violation_number LIKE ?", fmt.Sprintf("%%%s%%", illegalNum))
	}
	illegalType := strings.TrimSpace(req.IllegalType)
	if illegalType != "" {
		db = db.Where("violation_type LIKE ?", fmt.Sprintf("%%%s%%", illegalType))
	}
	illegalEffect := strings.TrimSpace(req.IllegalEffect)
	if illegalEffect != "" {
		db = db.Where("violation_effect LIKE ?", fmt.Sprintf("%%%s%%", illegalEffect))
	}
	mcnName := strings.TrimSpace(req.McnName)
	if mcnName != "" {
		db = db.Where("mcn_name LIKE ?", fmt.Sprintf("%%%s%%", mcnName))
	}
	status := req.Status
	if status != 0 {
		db = db.Where("account_status = ?", status)
	}
	// 当pageNum > 0 且 pageSize > 0 才分页
	//记录总条数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := int(req.PageNum)
	pageSize := int(req.PageSize)
	if pageNum > 0 && pageSize > 0 {
		//err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Preload("Roles").Find(&list).Error
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
