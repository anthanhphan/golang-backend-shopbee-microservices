package common

type SimpleUser struct {
	SQLModel `json:",inline"`
	FullName string `json:"fullname" gorm:"column:fullname;"`
	Role     string `json:"role" gorm:"column:role;"`
	Avatar   *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
