package controller

import (
	"github.com/Waelson/go-demo-user/internal/service"
	"net/http"
)

type UserController interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

func (c *userController) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get"))
}

func (c *userController) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post"))
}

func (c *userController) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update"))
}

func (c *userController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}

func NewUserController(userService service.UserService) UserController {
	result := &userController{
		userService: userService,
	}
	return result
}
