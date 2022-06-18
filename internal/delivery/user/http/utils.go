package http

import (
	httpCommon "github.com/aryahmph/restaurant/common/http"
	domainUser "github.com/aryahmph/restaurant/internal/domain/user"
)

func (h HTTPUserDelivery) mapUserBodyToDomain(payload httpCommon.AddUser) domainUser.User {
	user := domainUser.User{
		ID:           payload.ID,
		Username:     payload.Username,
		PasswordHash: payload.Password,
	}
	user.SetUserRoleString(payload.Role)
	return user
}

func (h HTTPUserDelivery) mapUserDomainToResponse(u domainUser.User) httpCommon.User {
	return httpCommon.User{
		ID:           u.ID,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		Role:         u.GetUserRoleString(),
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
