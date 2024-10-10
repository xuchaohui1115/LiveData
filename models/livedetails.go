package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"time"
)

type LiveDetails struct {
	LiveRoomId    int    `orm:"live_room_id" json:"live_room_id"`       // 直播间ID，主键
	LiveTitle     string `orm:"live_title" json:"live_title"`           // 直播标题
	LiveStartTime int    `orm:"live_start_time" json:"live_start_time"` // 开播时间（Unix 时间戳）
	HostId        int    `orm:"host_id" json:"host_id"`                 // 主播ID，用于关联主播
	CreatedAt     int    `orm:"created_at" json:"created_at"`           // 创建时间（Unix 时间戳)
}

func (*LiveDetails) TableName() string {
	return "live_details"
}

// CreateLiveDetails 新增直播详情
func CreateLiveDetails(liveDetail *LiveDetails) (int64, error) {
	o := orm.NewOrm()
	liveDetail.CreatedAt = int(time.Now().Unix())
	id, err := o.Insert(liveDetail)
	return id, err
}

// ReadLiveDetails 查询直播详情
func ReadLiveDetails(id int) (*LiveDetails, error) {
	o := orm.NewOrm()
	liveDetail := &LiveDetails{LiveRoomId: id}
	err := o.Read(liveDetail)
	return liveDetail, err
}

// UpdateLiveDetails 更新直播详情
func UpdateLiveDetails(liveDetail *LiveDetails) error {
	o := orm.NewOrm()
	_, err := o.Update(liveDetail)
	return err
}

// DeleteLiveDetails 删除直播详情
func DeleteLiveDetails(id int) error {
	o := orm.NewOrm()
	liveDetail := LiveDetails{LiveRoomId: id}
	_, err := o.Delete(&liveDetail)
	return err
}
