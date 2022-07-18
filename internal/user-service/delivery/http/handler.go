package userHttp

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/krulsaidme0w/golang_pet_project_3/pkg/context"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/httputil"
	userservice "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/models"
)

type Handler struct {
	userUseCase userservice.UserUseCase
	adapter     httputil.HttpFrameworkAdapter
}

func NewUserHandler(userUseCase userservice.UserUseCase) *Handler {
	return &Handler{
		userUseCase: userUseCase,
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
func (h *Handler) SaveUser(c *fiber.Ctx) error {
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
func (h *Handler) GetUser(c *fiber.Ctx) error {
	ctx, cancel := context.InitContext()
	defer cancel()

	userID := c.Params("id")

	user, err := h.userUseCase.Get(ctx, userID)
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
// @Param 		 data  body 	models.UserRequest true "user request"
// @Success      200  {object}  httputil.ReturnType
// @Failure      400  {object}  httputil.Error  "Cannot update user"
// @Router 		 /user [patch]
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	ctx, cancel := context.InitContext()
	defer cancel()

	userRequest := new(models.UserRequest)

	if err := h.adapter.GetPostBody(c, &userRequest); err != nil {
		h.adapter.Error(c, http.StatusBadRequest, err)
		return nil
	}

	err := h.userUseCase.Update(ctx, nil, userRequest)
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
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	ctx, cancel := context.InitContext()
	defer cancel()

	userID := c.Params("id")

	err := h.userUseCase.Delete(ctx, userID)
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
