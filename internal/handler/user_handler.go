package handler

import (
    "net/http"

    "github.com/example/gin-rest-api/internal/model"
    "github.com/example/gin-rest-api/internal/repository"
    "github.com/example/gin-rest-api/internal/response"
    "github.com/example/gin-rest-api/internal/service"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(c *gin.Context) {
    response.Success(c, http.StatusOK, h.service.GetAll())
}

func (h *UserHandler) GetByID(c *gin.Context) {
    user, err := h.service.GetByID(c.Param("id"))
    if err != nil {
        response.Error(c, http.StatusNotFound, "user not found")
        return
    }

    response.Success(c, http.StatusOK, user)
}

func (h *UserHandler) Create(c *gin.Context) {
    var req model.CreateUserRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, err.Error())
        return
    }

    response.Success(c, http.StatusCreated, h.service.Create(req))
}

func (h *UserHandler) Update(c *gin.Context) {
    var req model.UpdateUserRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, err.Error())
        return
    }

    user, err := h.service.Update(c.Param("id"), req)
    if err == repository.ErrNotFound {
        response.Error(c, http.StatusNotFound, "user not found")
        return
    }
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "failed to update user")
        return
    }

    response.Success(c, http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
    if err := h.service.Delete(c.Param("id")); err == repository.ErrNotFound {
        response.Error(c, http.StatusNotFound, "user not found")
        return
    }

    c.Status(http.StatusNoContent)
}
