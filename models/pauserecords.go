package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"time"
)

type PauseRecords struct {
	Id         int    `orm:"id" json:"id"`                     // 主键ID
	LiveDataId int    `orm:"live_data_id" json:"live_data_id"` // 直播数据表的ID，用于关联
	HostId     int    `orm:"host_id" json:"host_id"`           // 主播ID，用于关联
	PauseTime  int    `orm:"pause_time" json:"pause_time"`     // 暂停时间（Unix 时间戳）
	ResumeTime int    `orm:"resume_time" json:"resume_time"`   // 恢复时间（Unix 时间戳）
	Duration   int    `orm:"duration" json:"duration"`         // 暂停时长，以分钟为单位
	Reason     string `orm:"reason" json:"reason"`             // 暂停理由
	CreatedAt  int    `orm:"created_at" json:"created_at"`     // 创建时间（Unix 时间戳）
}

func (*PauseRecords) TableName() string {
	return "pause_records"
}

// CreatePauseRecord 新增暂停记录
func CreatePauseRecord(record *PauseRecords) (int64, error) {
	o := orm.NewOrm()
	record.CreatedAt = int(time.Now().Unix())
	id, err := o.Insert(record)
	return id, err
}

// ReadPauseRecord 查询暂停记录
func ReadPauseRecord(id int) (*PauseRecords, error) {
	o := orm.NewOrm()
	record := &PauseRecords{Id: id}
	err := o.Read(record)
	return record, err
}

// UpdatePauseRecord 更新暂停记录
func UpdatePauseRecord(record *PauseRecords) error {
	o := orm.NewOrm()
	_, err := o.Update(record)
	return err
}

// DeletePauseRecord 删除暂停记录
func DeletePauseRecord(id int) error {
	o := orm.NewOrm()
	record := PauseRecords{Id: id}
	_, err := o.Delete(&record)
	return err
}
