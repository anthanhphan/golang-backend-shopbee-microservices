package usermodel

import (
	"shopbee/common"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	FullName        string        `json:"fullname" gorm:"column:fullname;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Addr            string        `json:"addr" gorm:"column:addr;"`
	Avatar          *common.Image `json:"avatar" gorm:"column:avatar;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DbTypeUser)
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	FullName        string `json:"fullname" gorm:"column:fullname;"`
	Password        string `json:"password" gorm:"column:password;"`
	Salt            string `json:"-" gorm:"column:salt;"`
}

func (u *UserCreate) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DbTypeUser)
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type UserUpdate struct {
	Email    string        `json:"email" gorm:"column:email;"`
	FullName string        `json:"fullname" gorm:"column:fullname;"`
	Phone    string        `json:"phone" gorm:"column:phone;"`
	Addr     string        `json:"addr" gorm:"column:addr;"`
	Role     string        `json:"role" gorm:"column:role;"`
	Avatar   *common.Image `json:"avatar" gorm:"column:avatar;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}
