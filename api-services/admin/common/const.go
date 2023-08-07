package common

const (
	DbTypeShop     = 1
	DbTypeProduct  = 2
	DbTypeCategory = 3
	DbTypeUser     = 4
)

const (
	CurrentUser = "user"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
