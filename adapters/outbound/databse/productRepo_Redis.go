package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	domain "hexa-design/domain/repository"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type productRepositoryRedis struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductRepositoryRedis(db *gorm.DB, redisClient *redis.Client) domain.ProductRepository {
	db.AutoMigrate(&domain.Product{})
	mockData(db)
	return productRepositoryRedis{db, redisClient}
}

func mockData(db *gorm.DB) error {

	var count int64
	db.Model(&domain.Product{}).Count(&count)
	if count > 0 {
		return nil
	}

	seed := rand.NewSource(time.Now().Unix())
	random := rand.New(seed)
	products := []domain.Product{}
	for i := 0; i < 500; i++ {
		products = append(products, domain.Product{
			Name:     "",
			Quantity: random.Intn(100),
		})
	}
	return db.Create(&products).Error
}

func (r productRepositoryRedis) GetProducts() (products []domain.Product, err error) {

	key := "repository::GetProducts"

	// Redis Get
	productsJson, err := r.redisClient.Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(productsJson), &products)
		if err == nil {
			fmt.Println("Get to Redis")
			return products, nil
		}
	}
	// Database
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	if err != nil {
		return nil, err
	}
	// Redis Set ==> Marshal data
	data, err := json.Marshal(products)
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
	if err != nil {
		return nil, err
	}
	fmt.Println("Set to Redis DB")
	return products, err
}


