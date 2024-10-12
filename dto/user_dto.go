package dto

import "LiveData/model"

// 返回给前端的当前用户信息
type UserInfoDto struct {
	ID           int64         `json:"id"`
	Username     string        `json:"username"`
	Mobile       string        `json:"mobile"`
	Avatar       string        `json:"avatar"`
	Nickname     string        `json:"nickname"`
	Introduction string        `json:"introduction"`
	Roles        []*model.Role `json:"roles"`
}

func ToUserInfoDto(user model.User) UserInfoDto {
	return UserInfoDto{
		ID:           user.ID,
		Username:     user.Username,
		Mobile:       user.Mobile,
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		Introduction: user.Introduction,
		Roles:        user.Roles,
	}
}

// 返回给前端的用户列表
type UsersDto struct {
	ID           int64  `json:"ID"`
	Username     string `json:"username"`
	Mobile       string `json:"mobile"`
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Status       int    `json:"status"`
	Creator      string `json:"creator"`
	RoleIds      []int  `json:"roleIds"`
}

func ToUsersDto(userList []*model.User) []UsersDto {
	var users []UsersDto
	for _, user := range userList {
		userDto := UsersDto{
			ID:           user.ID,
			Username:     user.Username,
			Mobile:       user.Mobile,
			Avatar:       user.Avatar,
			Nickname:     user.Nickname,
			Introduction: user.Introduction,
			Status:       int(user.Status),
			Creator:      user.Creator,
		}
		roleIds := make([]int, 0)
		for _, role := range user.Roles {
			roleIds = append(roleIds, int(role.ID))
		}
		userDto.RoleIds = roleIds
		users = append(users, userDto)
	}

	return users
}
