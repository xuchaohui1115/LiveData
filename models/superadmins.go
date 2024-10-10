package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	"time"
)

type SuperAdmins struct {
	Id        int    `orm:"id" json:"id"`                 // 主键ID
	Username  string `orm:"username" json:"username"`     // 用户名，唯一
	Password  string `orm:"password" json:"password"`     // 密码（建议加密存储）
	Nickname  string `orm:"nickname" json:"nickname"`     // 昵称
	Phone     string `orm:"phone" json:"phone"`           // 手机号码
	Status    string `orm:"status" json:"status"`         // 状态
	CreatedAt int    `orm:"created_at" json:"created_at"` // 创建时间（Unix 时间戳）
	UpdatedAt int    `orm:"updated_at" json:"updated_at"` // 修改时间（Unix 时间戳）
}

func (*SuperAdmins) TableName() string {
	return "super_admins"
}

// CreateSuperAdmin 新增超级管理员
func CreateSuperAdmin(admin *SuperAdmins) (int64, error) {
	o := orm.NewOrm()
	admin.CreatedAt = int(time.Now().Unix())
	admin.UpdatedAt = int(time.Now().Unix())
	id, err := o.Insert(admin)
	return id, err
}

// ReadSuperAdmin 查询超级管理员
func ReadSuperAdmin(id int) (*SuperAdmins, error) {
	o := orm.NewOrm()
	admin := &SuperAdmins{Id: id}
	err := o.Read(admin)
	return admin, err
}

// UpdateSuperAdmin 更新超级管理员
func UpdateSuperAdmin(admin *SuperAdmins) error {
	o := orm.NewOrm()
	admin.UpdatedAt = int(time.Now().Unix())
	_, err := o.Update(admin)
	return err
}

// DeleteSuperAdmin 删除超级管理员
func DeleteSuperAdmin(id int) error {
	o := orm.NewOrm()
	admin := SuperAdmins{Id: id}
	_, err := o.Delete(&admin)
	return err
}
