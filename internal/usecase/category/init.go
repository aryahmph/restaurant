package category

import (
	"context"
	"fmt"
	errorCommon "github.com/aryahmph/restaurant/common/error"
	categoryDomain "github.com/aryahmph/restaurant/internal/domain/category"
	userDomain "github.com/aryahmph/restaurant/internal/domain/user"
	categoryRepo "github.com/aryahmph/restaurant/internal/repository/category"
	userRepo "github.com/aryahmph/restaurant/internal/repository/user"
)

type categoryUsecaseImpl struct {
	userRepository     userRepo.Repository
	categoryRepository categoryRepo.Repository
}

func NewCategoryUsecaseImpl(userRepository userRepo.Repository, categoryRepository categoryRepo.Repository) categoryUsecaseImpl {
	return categoryUsecaseImpl{userRepository: userRepository, categoryRepository: categoryRepository}
}

func (u categoryUsecaseImpl) Create(ctx context.Context, category categoryDomain.Category, adminID string) (id string, err error) {
	// Check admin role
	_, err = u.checkAdminByID(ctx, adminID)
	if err != nil {
		return id, err
	}

	// Check name existence
	err = u.nameAlreadyExist(ctx, category.Name)
	if err != nil {
		return id, err
	}

	// Save
	id, err = u.categoryRepository.Insert(ctx, category)
	if err != nil {
		return id, err
	}
	return id, err
}

func (u categoryUsecaseImpl) List(ctx context.Context) (categories []categoryDomain.Category, err error) {
	return u.categoryRepository.FindAll(ctx)
}

func (u categoryUsecaseImpl) Update(ctx context.Context, category categoryDomain.Category, adminID string) (id string, err error) {
	// Check admin role
	_, err = u.checkAdminByID(ctx, adminID)
	if err != nil {
		return id, err
	}

	// Check name existence
	categoryByID, err := u.categoryRepository.FindByName(ctx, category.Name)
	if err == nil {
		return id, errorCommon.NewInvariantError("name already exist")
	}

	// Update filter
	if category.Name != categoryByID.Name {
		// Check username existence
		err := u.nameAlreadyExist(ctx, category.Name)
		if err != nil {
			return id, err
		}
	}

	// Update
	return u.categoryRepository.Update(ctx, category)
}

func (u categoryUsecaseImpl) Delete(ctx context.Context, categoryID string, adminID string) (rid string, err error) {
	// Check admin role
	_, err = u.checkAdminByID(ctx, adminID)
	if err != nil {
		fmt.Println("error di check admin")
		return rid, err
	}

	// Check id existence
	_, err = u.categoryRepository.FindByID(ctx, categoryID)
	if err != nil {
		fmt.Println("error di findbyid")
		return rid, err
	}

	// Delete
	return u.categoryRepository.Delete(ctx, categoryID)
}

func (u categoryUsecaseImpl) checkAdminByID(ctx context.Context, id string) (admin userDomain.User, err error) {
	admin, err = u.userRepository.FindByID(ctx, id)
	if err != nil {
		return admin, err
	}

	if admin.Role != userDomain.Admin {
		return admin, errorCommon.NewForbiddenError("restricted user action")
	}
	return admin, nil
}

func (u categoryUsecaseImpl) nameAlreadyExist(ctx context.Context, username string) (err error) {
	_, err = u.categoryRepository.FindByName(ctx, username)
	if err == nil {
		return errorCommon.NewInvariantError("name already exist")
	}
	return nil
}
