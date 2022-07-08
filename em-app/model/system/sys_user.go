package system

// 用户表model
type User struct {
	Id         string `gorm:"column:id" form:"id" json:"id"`
	Username   string `gorm:"column:username" form:"username" json:"username"`
	Password   string `gorm:"column:password" form:"password" json:"password"`
	Nickname   string `gorm:"column:nickname" form:"nickname" json:"nickname"`
	Phone      string `gorm:"column:phone" form:"phone" json:"phone"`
	Email      string `gorm:"column:email" form:"email" json:"email"`
	Status     string `gorm:"column:status" form:"status" json:"status"`
	CreateTime string `gorm:"column:create_time" form:"createTime" json:"createTime"`
	UpdateTime string `gorm:"column:update_time" form:"updateTime" json:"updateTime"`
}

func (User) TableName() string {
	return "sys_user"
}
