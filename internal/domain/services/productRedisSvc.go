package services

import (
	"context"
	"encoding/json"
	"fmt"
	ports "hexa-design/internal/ports"
	"time"

	"github.com/go-redis/redis/v8"
)

type productServiceRedis struct {
	productRepo ports.ProductRepository
	redisClient *redis.Client
}

func NewProductServiceRedis(productRepo ports.ProductRepository, redisClient *redis.Client) ports.ProductService {
	return productServiceRedis{productRepo, redisClient}
}

func (s productServiceRedis) GetProducts() (products []ports.Product, err error) {
	key := "service::GetProducts"

	// Redis Get
	productsJson, err := s.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("Get to Redis")
			return products, nil
		}
	}
	// Database
	productDB, err := s.productRepo.GetProducts()

	if err != nil {
		return nil, err
	}

	for _, p := range productDB {
		products = append(products, ports.Product{
			ID:       p.ID,
			Name:     p.Name,
			Quantity: p.Quantity,
		})
	}
	// Redis Set ==> Marshal data
	if data, err := json.Marshal(products); err == nil {
		s.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}
	fmt.Println("Set to Redis DB")
	return products, err
}
