package ports

import "github.com/gofiber/fiber/v2"


type Product struct{
	ID int
	Name string
	Quantity int
}


type ProductRepository interface{
	GetProducts()([]Product,error)
}

type ProductService interface{
	GetProducts()([]Product,error)
}


type ProductHandler interface{
	GetProducts(c *fiber.Ctx)error
}