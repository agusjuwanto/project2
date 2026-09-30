package repository

import (
    "errors"
    "sync"

    "github.com/example/gin-rest-api/internal/model"
)

var ErrNotFound = errors.New("user not found")

type UserRepository interface {
    FindAll() []model.User
    FindByID(id string) (model.User, error)
    Create(user model.User) model.User
    Update(user model.User) (model.User, error)
    Delete(id string) error
}

type userRepository struct {
    mu    sync.RWMutex
    users map[string]model.User
}

func NewUserRepository() UserRepository {
    return &userRepository{
        users: make(map[string]model.User),
    }
}

func (r *userRepository) FindAll() []model.User {
    r.mu.RLock()
    defer r.mu.RUnlock()

    result := make([]model.User, 0, len(r.users))
    for _, user := range r.users {
        result = append(result, user)
    }
    return result
}

func (r *userRepository) FindByID(id string) (model.User, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()

    user, ok := r.users[id]
    if !ok {
        return model.User{}, ErrNotFound
    }
    return user, nil
}

func (r *userRepository) Create(user model.User) model.User {
    r.mu.Lock()
    defer r.mu.Unlock()

    r.users[user.ID] = user
    return user
}

func (r *userRepository) Update(user model.User) (model.User, error) {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, ok := r.users[user.ID]; !ok {
        return model.User{}, ErrNotFound
    }

    r.users[user.ID] = user
    return user, nil
}

func (r *userRepository) Delete(id string) error {
    r.mu.Lock()
    defer r.mu.Unlock()

    if _, ok := r.users[id]; !ok {
        return ErrNotFound
    }

    delete(r.users, id)
    return nil
}
