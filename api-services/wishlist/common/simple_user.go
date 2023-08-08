package common

type SimpleUser struct {
	SQLModel `json:",inline"`
	FullName string `json:"fullname" gorm:"column:fullname;"`
	Email    string `json:"email" gorm:"column:email;"`
	Role     string `json:"role" gorm:"column:role;"`
}

func (u *SimpleUser) GetUserId() int {
	return u.Id
}

func (u *SimpleUser) GetEmail() string {
	return u.Email
}

func (u *SimpleUser) GetRole() string {
	return u.Role
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
