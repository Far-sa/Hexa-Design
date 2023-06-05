package handler

import (
	ports "hexa-design/internal/ports"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productSvc ports.ProductService
}

func NewProductHandler(productSvc ports.ProductService) ports.ProductHandler {
	return productHandler{productSvc: productSvc}
}

func (h productHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.productSvc.GetProducts()
	if err != nil {
		return err
	}

	response := fiber.Map{
		"status":   "200",
		"products": products,
	}
	return c.JSON(response)
}
