package adapters

import (
	"github.com/fiatfour/go-unit/core"
	"github.com/gofiber/fiber/v2"
)

// ! Primary adapter (http_adapter.go)

type HttpOrderHandler struct {
	service core.OrderService
}

func NewHttpOrderHandler(service core.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{service: service}
}

func (h *HttpOrderHandler) CreateOrderHandler(c *fiber.Ctx) error {
	var order core.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.service.CreateOrder(order); err != nil {
		// Return an appropriate error message and status code
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
