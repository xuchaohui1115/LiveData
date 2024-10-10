package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"time"
)

type LiveData struct {
	Id               int     `orm:"id" json:"id"`                                 // 主键ID
	HostNickname     string  `orm:"host_nickname" json:"host_nickname"`           // 主播昵称
	HostId           int     `orm:"host_id" json:"host_id"`                       // 主播ID
	LiveStatus       string  `orm:"live_status" json:"live_status"`               // 直播状态
	HostStatus       string  `orm:"host_status" json:"host_status"`               // 主播状态
	Remarks          string  `orm:"remarks" json:"remarks"`                       // 备注
	LiveDuration     int     `orm:"live_duration" json:"live_duration"`           // 直播时长（以分钟为单位）
	PaymentAmount    float64 `orm:"payment_amount" json:"payment_amount"`         // 支付金额
	ViewCount        int     `orm:"view_count" json:"view_count"`                 // 观看次数
	ClickThroughRate float64 `orm:"click_through_rate" json:"click_through_rate"` // 点击率
	ConversionRate   float64 `orm:"conversion_rate" json:"conversion_rate"`       // 成交率
	TransactionCount int     `orm:"transaction_count" json:"transaction_count"`   // 成交人数
	TransactionItems int     `orm:"transaction_items" json:"transaction_items"`   // 成交件数
	YesterdaySales   float64 `orm:"yesterday_sales" json:"yesterday_sales"`       // 昨日成交
	PauseCount       int     `orm:"pause_count" json:"pause_count"`               // 暂停次数
	PauseRecordId    int     `orm:"pause_record_id" json:"pause_record_id"`       // 暂停记录表ID，用于关联
	CreatedAt        int     `orm:"created_at" json:"created_at"`                 // 创建时间（Unix 时间戳）
	UpdatedAt        int     `orm:"updated_at" json:"updated_at"`                 // 修改时间（Unix 时间戳）
}

func (*LiveData) TableName() string {
	return "live_data"
}

// CreateLiveData 新增直播数据
func CreateLiveData(liveData *LiveData) (int64, error) {
	o := orm.NewOrm()
	liveData.CreatedAt = int(time.Now().Unix())
	liveData.UpdatedAt = int(time.Now().Unix())
	id, err := o.Insert(liveData)
	return id, err
}

// ReadLiveData 查询直播数据
func ReadLiveData(id int) (*LiveData, error) {
	o := orm.NewOrm()
	liveData := &LiveData{Id: id}
	err := o.Read(liveData)
	return liveData, err
}

// UpdateLiveData 更新直播数据
func UpdateLiveData(liveData *LiveData) error {
	o := orm.NewOrm()
	liveData.UpdatedAt = int(time.Now().Unix())
	_, err := o.Update(liveData)
	return err
}

// DeleteLiveData 删除直播数据
func DeleteLiveData(id int) error {
	o := orm.NewOrm()
	liveData := LiveData{Id: id}
	_, err := o.Delete(&liveData)
	return err
}
