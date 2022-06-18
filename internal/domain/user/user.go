package user

import "github.com/aryahmph/restaurant/internal/domain"

const (
	Waiter  role = "WAITER"
	Cashier role = "CASHIER"
	Admin   role = "ADMIN"
)

type (
	User struct {
		ID           string
		Username     string
		PasswordHash string
		Role         role

		Timestamp
	}

	role string

	Timestamp = domain.Timestamp
)

func (u *User) GetUserRoleString() string {
	return string(u.Role)
}

func (u *User) SetUserRoleString(r string) {
	switch r {
	case string(Waiter):
		u.Role = Waiter
	case string(Cashier):
		u.Role = Cashier
	case string(Admin):
		u.Role = Admin
	}
}
