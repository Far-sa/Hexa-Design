package ports

import (
	"hexa-design/internal/domain/model"

	"github.com/gofiber/fiber/v2"
)




type ProductRepository interface{
	GetProducts()([]model.Product,error)
}

type ProductService interface{
	GetProducts()([]model.Product,error)
}


type ProductHandler interface{
	GetProducts(c *fiber.Ctx)error
}