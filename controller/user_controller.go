package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"teashop/model"
	"teashop/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type UserController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	userService service.UserService
}

// Get all userlist godoc
// @tags user-API
// @Summary Get userlist
// @Description output: user-item struct list
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/all [get]
func (c *userController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	tmp, err := c.userService.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "Get all successfuly",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get user by id godoc
// @tags user-API
// @Summary Get user by id
// @Description input: user's id => output: user-item struct
// @Accept json
// @Produce json
// @Param id path integer true "user id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/{id} [get]
func (c *userController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get param failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.userService.GetById(id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "Get by ID successfuly",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}


// Creat user item  godoc
// @tags user-API
// @Summary Create new user
// @Description input: user's struct
// @Accept json
// @Produce json
// @Param user body model.User true "Fill the form"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/create [post]
func (c *userController) Create(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var user *model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
		res = &model.Response{
			Data:    nil,
			Message: "Create user failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.userService.Create(user)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "Create user successfuly",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Update user item  godoc
// @tags user-API
// @Summary Update info user
// @Description input: user's struct to update
// @Accept json
// @Produce json
// @Param user body model.User true "Fill the form"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/update [put]
func (c *userController) Update(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "Update failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.userService.Update(user.Id, &user)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "Update user successfuly",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Delete user item godoc
// @tags user-API
// @Summary Delete one user
// @Description input: userID
// @Accept json
// @Produce json
// @Param id path integer true "user Id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /user/{id} [delete]
func (c *userController) Delete(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "Delete user failed",
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	err = c.userService.Delete(id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    nil,
			Message: "Delete user successfuly",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

func NewUserController() UserController {
	userService := service.NewUserService()
	return &userController{
		userService: userService,
	}
}
