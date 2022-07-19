package http

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/krulsaidme0w/golang_pet_project_3/pkg/context"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/httputil"
	userservice "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/models"
)

type UserHandler struct {
	userUseCase userservice.UserUseCase
	adapter     *httputil.FiberFrameworkAdapter
}

func NewUserHandler(userUseCase userservice.UserUseCase, adapter *httputil.FiberFrameworkAdapter) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
		adapter:     adapter,
	}
}

// SaveUser godoc
// @Summary      SaveUser
// @Description  save new user
// @Tags         user
// @Accept       application/json
// @Produce      application/json
// @Param 		 data  body 	models.UserRequest true "user request"
// @Success      200  {object}  httputil.ReturnType
// @Failure      400  {object}  httputil.Error  "Cannot save user"
// @Router 		 /user [post]
func (h *UserHandler) SaveUser(c *fiber.Ctx) error {
	ctx, cancel := context.InitContext()
	defer cancel()

	userRequest := new(models.UserRequest)

	if err := h.adapter.GetPostBody(c, &userRequest); err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	err := h.userUseCase.Save(ctx, userRequest)
	if err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	h.adapter.Success(c, httputil.ReturnType{
		Status:  "ok",
		Message: "new user saved",
	})

	return nil
}

// GetUser godoc
// @Summary      GetUser
// @Description  get user by id
// @Tags         user
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true   "id of user to get"
// @Success      200  {object}  models.User
// @Failure      400  {object}  httputil.Error  "User not found"
// @Router 		 /user/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.InitContext()
	defer cancel()

	userID := c.Params("id")

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	user, err := h.userUseCase.Get(ctx, id)
	if err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	h.adapter.Success(c, user)

	return nil
}

// UpdateUser godoc
// @Summary      UpdateUser
// @Description  update user
// @Tags         user
// @Accept       application/json
// @Produce      application/json
// @Param 		 data  body 	models.User true "updated user"
// @Success      200  {object}  httputil.ReturnType
// @Failure      400  {object}  httputil.Error  "Cannot update user"
// @Router 		 /user/update [post]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	ctx, cancel := context.InitContext()
	defer cancel()

	user := new(models.User)

	if err := h.adapter.GetPostBody(c, &user); err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	if _, err := h.userUseCase.Get(ctx, user.ID); err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	err := h.userUseCase.Update(ctx, user)
	if err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	h.adapter.Success(c, httputil.ReturnType{
		Status:  "ok",
		Message: "user updated",
	})

	return nil
}

// GetUser godoc
// @Summary      DeleteUser
// @Description  get user by id
// @Tags         user
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true   "id of user to delete"
// @Success      200  {object}  models.User
// @Failure      400  {object}  httputil.Error  "User not found"
// @Router 		 /user/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.InitContext()
	defer cancel()

	userID := c.Params("id")

	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	err = h.userUseCase.Delete(ctx, id)
	if err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	h.adapter.Success(c, httputil.ReturnType{
		Status:  "ok",
		Message: "user deleted",
	})

	return nil
}
