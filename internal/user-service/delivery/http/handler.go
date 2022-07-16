package userHttp

import (
	"github.com/gofiber/fiber"
	userservice "github.com/krulsaidme0w/golang_pet_project_3/internal/user-service"
)

type Handler struct {
	userUseCase userservice.UseCase
}

func NewUserHandler(userUseCase userservice.UseCase) *Handler {
	return &Handler{
		userUseCase: userUseCase,
	}
}

// GetUser godoc
// @Summary      Get
// @Description  get user by id
// @Tags         user
// @Accept       application/json
// @Produce      application/json
// @Param        id   path      integer  true   "id of user to get"
// @Success      200  {object}  user-service.User
// @Failure      400  {object}  webUtils.Error  "Data is invalid"
// @Failure      404  {object}  webUtils.Error  "User not found"
// @Failure      405  {object}  webUtils.Error  "Method is not allowed"
func (h *Handler) GetUser(c *fiber.Ctx) error {
	return nil
}
