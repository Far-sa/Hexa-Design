package repository

import (
	ports "hexa-design/internal/ports/repository"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDb(db *gorm.DB) ports.ProductRepository {
	db.AutoMigrate(&ports.Product{})
	mockData(db)
	return productRepositoryDB{db}
}

func (r productRepositoryDB) GetProducts() (products []ports.Product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err

}