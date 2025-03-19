package utils

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

type HTTPResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(data interface{}) HTTPResponse {
	return HTTPResponse{
		Success: true,
		Data:    data,
	}
}

func SendSuccess(c *fiber.Ctx, statusCode int, data interface{}) error {
	return c.Status(statusCode).JSON(SuccessResponse(data))
}

func SendError(c *fiber.Ctx, statusCode int, message string, field string) error {
	return c.Status(statusCode).JSON(NewErrorResponse(message, field))
}

func BindAndValidate(c *fiber.Ctx, req interface{}) error {
	if err := c.BodyParser(req); err != nil {
		return err
	}

	validator := NewValidator()
	validationsErrors := validator.Validate(req)

	if len(validationsErrors) > 0 {
		messageJsonString, _ := json.Marshal(validationsErrors)
		return fiber.NewError(fiber.StatusUnprocessableEntity, string(messageJsonString))
	}

	return nil
}
