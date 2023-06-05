package repository

import (
	domain "hexa-design/domain/model"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
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

func NewProductRepositoryDb(db *gorm.DB) domain.ProductRepository {
	db.AutoMigrate(&domain.Product{})
	mockData(db)
	return productRepositoryDB{db}
}

func (r productRepositoryDB) GetProducts() (products []domain.Product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err

}