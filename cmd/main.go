package main

import (
	"fmt"
	handler "hexa-design/internal/adapters/inbound/rest"
	repository "hexa-design/internal/adapters/outbound/databse"
	"hexa-design/internal/domain/services"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db := initDatabase()
	redisClient := initRedis()
	_ = redisClient

	//productRepo := repository.NewProductRepositoryRedis(db, redisClient)
	//productService := services.NewProductServiceRedis(productRepo, redisClient)
	productRepo := repository.NewProductRepositoryDb(db)
	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	app := fiber.New()

	app.Get("/products", productHandler.GetProducts)

	app.Listen(":8000")
}

func initDatabase() *gorm.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	db_host := os.Getenv("DB_HOST")

	dialect := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s", db_user, db_pass, db_host, db_name))
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
