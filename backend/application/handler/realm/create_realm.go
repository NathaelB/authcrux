package realm

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathaelb/authcrux/application/utils"
	"github.com/nathaelb/authcrux/domain/port"
)

type CreateRealmHandler struct {
	realmService port.RealmService
}

type CreateRealmRequest struct {
	Name string `json:"name" validate:"required"`
}

func NewCreateRealmHandler(realmService port.RealmService) *CreateRealmHandler {
	return &CreateRealmHandler{
		realmService: realmService,
	}
}

func (h *CreateRealmHandler) Handle(c *fiber.Ctx) error {
	var request CreateRealmRequest

	if err := utils.BindAndValidate(c, &request); err != nil {
		return err
	}

	realm, err := h.realmService.CreateRealm(request.Name)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessResponse(realm))
}
