package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"time"
)

type Violations struct {
	Id                  int    `orm:"id" json:"id"`                                         // 主键ID
	HostNickname        string `orm:"host_nickname" json:"host_nickname"`                   // 主播昵称
	HostId              int    `orm:"host_id" json:"host_id"`                               // 主播ID
	SecurityCode        string `orm:"security_code" json:"security_code"`                   // 安全码
	AccountStatus       string `orm:"account_status" json:"account_status"`                 // 账号状态
	ViolationNumber     string `orm:"violation_number" json:"violation_number"`             // 违规编号
	ViolationTime       int    `orm:"violation_time" json:"violation_time"`                 // 违规时间（Unix 时间戳）
	ViolationReason     string `orm:"violation_reason" json:"violation_reason"`             // 违规原因
	ViolationImpact     string `orm:"violation_impact" json:"violation_impact"`             // 违规影响
	ViolationLiveRoomId string `orm:"violation_live_room_id" json:"violation_live_room_id"` // 违规对象是直播间的ID
	CreatedAt           int    `orm:"created_at" json:"created_at"`                         // 创建时间（Unix 时间戳）
}

func (*Violations) TableName() string {
	return "violations"
}

// CreateViolation 新增违规记录
func CreateViolation(violation *Violations) (int64, error) {
	o := orm.NewOrm()
	violation.CreatedAt = int(time.Now().Unix())
	id, err := o.Insert(violation)
	return id, err
}

// ReadViolation 查询违规记录
func ReadViolation(id int) (*Violations, error) {
	o := orm.NewOrm()
	violation := &Violations{Id: id}
	err := o.Read(violation)
	return violation, err
}

// UpdateViolation 更新违规记录
func UpdateViolation(violation *Violations) error {
	o := orm.NewOrm()
	_, err := o.Update(violation)
	return err
}

// DeleteViolation 删除违规记录
func DeleteViolation(id int) error {
	o := orm.NewOrm()
	violation := Violations{Id: id}
	_, err := o.Delete(&violation)
	return err
}
