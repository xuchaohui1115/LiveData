// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID           int64          `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	CreatedAt    time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"-"`
	Username     string         `gorm:"column:username;type:varchar(20);not null" json:"username"`
	Password     string         `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Mobile       string         `gorm:"column:mobile;type:varchar(11);not null" json:"mobile"`
	Avatar       string         `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	Nickname     string         `gorm:"column:nickname;type:varchar(20)" json:"nickname"`
	Introduction string         `gorm:"column:introduction;type:varchar(255)" json:"introduction"`
	Status       int64          `gorm:"column:status;type:tinyint(1);default:1;comment:'1正常, 2禁用'" json:"status"` // '1正常, 2禁用'
	Creator      string         `gorm:"column:creator;type:varchar(20)" json:"creator"`
	Roles        []*Role        `gorm:"many2many:user_roles" json:"roles"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
