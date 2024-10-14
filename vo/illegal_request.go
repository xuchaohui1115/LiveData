package vo

// 获取违规列表结构体
type IllegalRequest struct {
	Nickname      string `json:"nickname" form:"nickname"`
	IllegalNum    string `json:"illegalNum" form:"illegalNum"`
	IllegalType   string `json:"illegalType" form:"illegalType"`
	IllegalEffect string `json:"illegalEffect" form:"illegalEffect"`
	McnName       string `json:"mcnName" form:"mcnName"`
	Status        int    `json:"status" form:"status"`
	PageNum       int    `json:"pageNum" form:"pageNum"`
	PageSize      int    `json:"pageSize" form:"pageSize"`
}
