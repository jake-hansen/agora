package services

import "github.com/jake-hansen/agora/database/domain"

type userService struct {
	repo	domain.UserRepository
}

func NewUserService(repository domain.UserRepository) domain.UserService {
	return &userService{repo: repository}
}

func (u *userService) Create(user *domain.User) (uint, error) {
	return u.repo.Create(user)
}

func (u *userService) GetAll() ([]*domain.User, error) {
	return u.repo.GetAll()
}

func (u *userService) GetByID(ID uint) (*domain.User, error) {
	return u.repo.GetByID(ID)
}

func (u *userService) Update(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *userService) Delete(ID uint) error {
	return u.repo.Delete(ID)
}

