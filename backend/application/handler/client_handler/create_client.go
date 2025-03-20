package client_handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nathaelb/authcrux/application/utils"
	"github.com/nathaelb/authcrux/domain/client"
)

type CreateClientHandler struct {
	clientService client.ClientService
}

func NewCreateClientHandler(clientService client.ClientService) *CreateClientHandler {
	return &CreateClientHandler{
		clientService: clientService,
	}
}

type CreateClientRequest struct {
	ClientID string `json:"client_id" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (h *CreateClientHandler) Handle(c *fiber.Ctx) error {
	var request CreateClientRequest
	realmID := c.Params("realm_id")

	if realmID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "realm_id is required")
	}

	if err := utils.BindAndValidate(c, &request); err != nil {
		return err
	}

	client, err := h.clientService.CreateClient(realmID, request.ClientID, request.Name)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(client)
}
