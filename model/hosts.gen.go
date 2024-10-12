// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHost = "hosts"

// Host 主播基础信息表
type Host struct {
	ID             int64  `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id,string"`
	Username       string `gorm:"column:username;type:varchar(50);not null;comment:用户名" json:"username"`                 // 用户名
	Password       string `gorm:"column:password;type:varchar(50);not null;comment:密码（建议加密存储）" json:"password"`          // 密码（建议加密存储）
	Nickname       string `gorm:"column:nickname;type:varchar(50);not null;comment:昵称" json:"nickname"`                  // 昵称
	Phone          string `gorm:"column:phone;type:varchar(15);not null;comment:手机号码" json:"phone"`                      // 手机号码
	Email          string `gorm:"column:email;type:varchar(100);not null;comment:电子邮件" json:"email"`                     // 电子邮件
	ProfilePicture string `gorm:"column:profile_picture;type:varchar(255);not null;comment:头像URL" json:"profilePicture"` // 头像URL
	Bio            string `gorm:"column:bio;type:text;not null;comment:个人简介" json:"bio"`                                 // 个人简介
	FollowersCount int64  `gorm:"column:followers_count;type:int(11);not null;comment:粉丝数量" json:"followersCount"`       // 粉丝数量
	CreatedAt      int64  `gorm:"column:created_at;comment:创建时间" json:"createdAt"`                                       // 创建时间（Unix 时间戳）
	UpdatedAt      int64  `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"`                                       // 更新时间（Unix 时间戳）
}

// TableName Host's table name
func (*Host) TableName() string {
	return TableNameHost
}
