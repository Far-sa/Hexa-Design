package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"hexa-design/internal/ports"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type productHandlerRedis struct {
	productSvc  ports.ProductService
	redisClient *redis.Client
}

func NewProductHandlerRedis(productSvc ports.ProductService, redisClient *redis.Client) ports.ProductHandler {
	return productHandlerRedis{productSvc, redisClient}
}

func (h productHandlerRedis) GetProducts(c *fiber.Ctx) error {

	key := "handler::GetProducts"

	//* REDIS GET
	if resJson, err := h.redisClient.Get(context.Background(), key).Result(); err == nil {
		fmt.Println("REDIS ----> GET")
		c.Set("Content-Type", "Application/json")
		return c.SendString(resJson)
	}
	//
	products, err := h.productSvc.GetProducts()
	if err != nil {
		return err
	}
	response := fiber.Map{
		"status":   "200",
		"products": products,
	}

	//* REDIS SET
	if data, err := json.Marshal(response); err == nil {
		h.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}
	fmt.Println("SET -----> REDIS")
	return c.JSON(response)
}
