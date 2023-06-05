package adapters

import (
	domain "hexa-design/domain/repository"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
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