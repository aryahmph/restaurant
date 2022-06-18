package user

import (
	"context"
	errorCommon "github.com/aryahmph/restaurant/common/error"
	jwtCommon "github.com/aryahmph/restaurant/common/jwt"
	passCommon "github.com/aryahmph/restaurant/common/password"
	userDomain "github.com/aryahmph/restaurant/internal/domain/user"
	userRepo "github.com/aryahmph/restaurant/internal/repository/user"
)

type userUsecaseImpl struct {
	userRepository  userRepo.Repository
	passwordManager *passCommon.PasswordHashManager
	jwtManager      *jwtCommon.JWTManager
}

func NewUserUsecaseImpl(userRepository userRepo.Repository, passwordManager *passCommon.PasswordHashManager,
	jwtManager *jwtCommon.JWTManager) userUsecaseImpl {
	return userUsecaseImpl{userRepository: userRepository, passwordManager: passwordManager, jwtManager: jwtManager}
}

func (u userUsecaseImpl) Register(ctx context.Context, user userDomain.User, adminID string) (id string, err error) {
	// Check admin role
	_, err = u.checkAdminByID(ctx, adminID)
	if err != nil {
		return id, err
	}

	// Check username existence
	err = u.usernameAlreadyExist(ctx, user.Username)
	if err != nil {
		return id, err
	}

	// Hash password
	hashPassword, err := u.passwordManager.HashPassword(user.PasswordHash)
	if err != nil {
		return id, err
	}
	user.PasswordHash = hashPassword

	// Save
	id, err = u.userRepository.Insert(ctx, user)
	if err != nil {
		return id, err
	}
	return id, err
}

func (u userUsecaseImpl) List(ctx context.Context, adminID string) (users []userDomain.User, err error) {
	// Check admin role
	_, err = u.checkAdminByID(ctx, adminID)
	if err != nil {
		return users, err
	}

	return u.userRepository.FindAll(ctx)
}

func (u userUsecaseImpl) GetByID(ctx context.Context, id string, adminID string) (user userDomain.User, err error) {
	// Check same user
	userByID, err := u.checkSameUserByID(ctx, id)
	if err != nil {
		// Check admin role
		_, err = u.checkAdminByID(ctx, adminID)
	}
	return userByID, nil
}

func (u userUsecaseImpl) Update(ctx context.Context, user userDomain.User, adminID string) (id string, err error) {
	// Check same user
	userByID, err := u.checkSameUserByID(ctx, user.ID)
	if err != nil {
		// Check admin role
		_, err = u.checkAdminByID(ctx, adminID)
		if err != nil {
			return id, err
		}
	}

	// Check username existence
	_, err = u.userRepository.FindByUsername(ctx, user.Username)
	if err == nil {
		return id, errorCommon.NewInvariantError("username already exist")
	}

	// Update filter
	if user.Username != userByID.Username {
		// Check username existence
		err := u.usernameAlreadyExist(ctx, user.Username)
		if err != nil {
			return id, err
		}
	}

	// Update
	return u.userRepository.Update(ctx, user)
}

func (u userUsecaseImpl) Delete(ctx context.Context, id string, adminID string) (rid string, err error) {
	// Check admin role
	_, err = u.checkAdminByID(ctx, adminID)
	if err != nil {
		return rid, err
	}

	// Check id existence
	_, err = u.userRepository.FindByID(ctx, id)
	if err != nil {
		return rid, err
	}

	// Delete
	return u.userRepository.Delete(ctx, id)
}

func (u userUsecaseImpl) checkAdminByID(ctx context.Context, id string) (admin userDomain.User, err error) {
	admin, err = u.userRepository.FindByID(ctx, id)
	if err != nil {
		return admin, err
	}

	if admin.Role != userDomain.Admin {
		return admin, errorCommon.NewForbiddenError("restricted user action")
	}
	return admin, nil
}

func (u userUsecaseImpl) checkSameUserByID(ctx context.Context, id string) (user userDomain.User, err error) {
	userByID, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return user, err
	}

	if userByID.ID != id {
		return user, errorCommon.NewForbiddenError("restricted user action")
	}
	return userByID, nil
}

func (u userUsecaseImpl) usernameAlreadyExist(ctx context.Context, username string) (err error) {
	_, err = u.userRepository.FindByUsername(ctx, username)
	if err == nil {
		return errorCommon.NewInvariantError("username already exist")
	}
	return nil
}
