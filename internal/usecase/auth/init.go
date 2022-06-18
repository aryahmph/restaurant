package auth

import (
	"context"
	"github.com/aryahmph/restaurant/common/jwt"
	passCommon "github.com/aryahmph/restaurant/common/password"
	userRepo "github.com/aryahmph/restaurant/internal/repository/user"
	"time"
)

type authUsecase struct {
	userRepository  userRepo.Repository
	passwordManager *passCommon.PasswordHashManager
	jwtManager      *jwt.JWTManager
}

func NewAuthUsecase(userRepository userRepo.Repository, passwordManager *passCommon.PasswordHashManager, jwtManager *jwt.JWTManager) authUsecase {
	return authUsecase{userRepository: userRepository, passwordManager: passwordManager, jwtManager: jwtManager}
}

func (u authUsecase) Login(ctx context.Context, username string, password string) (accessToken string, err error) {
	// Get database row
	user, err := u.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return accessToken, err
	}

	// Compare password input w/ database
	if err := u.passwordManager.CheckPasswordHash(password, user.PasswordHash); err != nil {
		return accessToken, err
	}

	return u.jwtManager.GenerateToken(user.ID, time.Hour*9)
}
