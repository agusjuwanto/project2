package service

import (
    "github.com/example/gin-rest-api/internal/model"
    "github.com/example/gin-rest-api/internal/repository"
    "github.com/google/uuid"
)

type UserService interface {
    GetAll() []model.User
    GetByID(id string) (model.User, error)
    Create(req model.CreateUserRequest) model.User
    Update(id string, req model.UpdateUserRequest) (model.User, error)
    Delete(id string) error
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) GetAll() []model.User {
    return s.repo.FindAll()
}

func (s *userService) GetByID(id string) (model.User, error) {
    return s.repo.FindByID(id)
}

func (s *userService) Create(req model.CreateUserRequest) model.User {
    user := model.User{
        ID:    uuid.NewString(),
        Name:  req.Name,
        Email: req.Email,
    }
    return s.repo.Create(user)
}

func (s *userService) Update(id string, req model.UpdateUserRequest) (model.User, error) {
    user := model.User{
        ID:    id,
        Name:  req.Name,
        Email: req.Email,
    }
    return s.repo.Update(user)
}

func (s *userService) Delete(id string) error {
    return s.repo.Delete(id)
}
