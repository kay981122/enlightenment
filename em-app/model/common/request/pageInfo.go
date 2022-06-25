package request

// 分页信息公用实体
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 页大小
	Keyword  string `json:"keyword" form:"keyword"`   // 关键字
}

type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
