package repository

import (
	"LiveData/common"
	"LiveData/model"
	"LiveData/vo"
	"fmt"
	"strings"
)

type IAnchorRepository interface {
	GetGetAnchor(req *vo.AnchorListRequest) ([]*model.LiveData, int64, error) // 获取用户列表
}

type AnchorRepository struct {
}

// NewAnchorRepository 构造函数
func NewAnchorRepository() IAnchorRepository {
	return AnchorRepository{}
}

// 获取主播列表
func (ar AnchorRepository) GetGetAnchor(req *vo.AnchorListRequest) ([]*model.LiveData, int64, error) {
	var list []*model.LiveData
	db := common.DB.Model(&model.LiveData{}).Order("created_at DESC")
	/*
		TbAccount   string `json:"tb_account" form:"tb_account"`
		Mobile      string `json:"mobile" form:"mobile"`
		Nickname    string `json:"nickname" form:"nickname"`
		Uid         int    `json:"uid" form:"uid"`
		Status      int    `json:"status" form:"status"`
		LoginStatus int    `json:"loginStatus" form:"loginStatus"`
		McnId       string `json:"mcnId" form:"mcnId"`
		BusinessId  int    `json:"businessId" form:"businessId"`
		PageNum     int    `json:"pageNum" form:"pageNum"`
		PageSize    int    `json:"pageSize" form:"pageSize"`
	*/
	tbAccount := strings.TrimSpace(req.TbAccount)
	if tbAccount != "" {
		db = db.Where("tb_account LIKE ?", fmt.Sprintf("%%%s%%", tbAccount))
	}
	nickname := strings.TrimSpace(req.Nickname)
	if nickname != "" {
		db = db.Where("host_nickname LIKE ?", fmt.Sprintf("%%%s%%", nickname))
	}
	mobile := strings.TrimSpace(req.Mobile)
	if mobile != "" {
		db = db.Where("phone LIKE ?", fmt.Sprintf("%%%s%%", mobile))
	}
	uid := req.Uid
	if uid != 0 {
		db = db.Where("uid = ?", uid)
	}
	//status := req.Status
	//if status != 0 {
	//	db = db.Where("status = ?", status)
	//}
	loginStatus := req.LoginStatus
	if loginStatus != 0 {
		db = db.Where("login_status = ?", loginStatus)
	}
	mcnId := strings.TrimSpace(req.McnId)
	if mcnId != "" {
		db = db.Where("mcn_name LIKE ?", fmt.Sprintf("%%%s%%", mcnId))
	}
	//todo 对接人 需要根据id找到用户

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
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Preload("Roles").Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
