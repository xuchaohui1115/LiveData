package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"time"
)

type Hosts struct {
	Id             int    `orm:"id" json:"id"`
	Username       string `orm:"username" json:"username"`               // 用户名
	Password       string `orm:"password" json:"password"`               // 密码（建议加密存储）
	Nickname       string `orm:"nickname" json:"nickname"`               // 昵称
	Phone          string `orm:"phone" json:"phone"`                     // 手机号码
	Email          string `orm:"email" json:"email"`                     // 电子邮件
	ProfilePicture string `orm:"profile_picture" json:"profile_picture"` // 头像URL
	Bio            string `orm:"bio" json:"bio"`                         // 个人简介
	FollowersCount int    `orm:"followers_count" json:"followers_count"` // 粉丝数量
	CreatedAt      int    `orm:"created_at" json:"created_at"`           // 创建时间（Unix 时间戳）
	UpdatedAt      int    `orm:"updated_at" json:"updated_at"`           // 更新时间（Unix 时间戳）
}

func (*Hosts) TableName() string {
	return "hosts"
}

// CreateHosts 新增主播
func CreateHosts(host *Hosts) (int64, error) {
	o := orm.NewOrm()
	host.CreatedAt = int(time.Now().Unix())
	host.UpdatedAt = int(time.Now().Unix())
	id, err := o.Insert(host)
	return id, err
}

// ReadHosts 查询主播
func ReadHosts(id int) (*Hosts, error) {
	o := orm.NewOrm()
	host := &Hosts{Id: id}
	err := o.Read(host)
	return host, err
}

// UpdateHosts 更新主播
func UpdateHosts(host *Hosts) error {
	o := orm.NewOrm()
	host.UpdatedAt = int(time.Now().Unix())
	_, err := o.Update(host)
	return err
}

// DeleteHosts 删除主播
func DeleteHosts(id int) error {
	o := orm.NewOrm()
	host := Hosts{Id: id}
	_, err := o.Delete(&host)
	return err
}
