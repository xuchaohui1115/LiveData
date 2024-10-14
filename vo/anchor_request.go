package vo

// 获取主播列表结构体
type AnchorListRequest struct {
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
}
