package repository

import (
	"hexa-design/internal/domain/model"
	ports "hexa-design/internal/ports"

	"gorm.io/gorm"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDb(db *gorm.DB) ports.ProductRepository {
	db.AutoMigrate(&model.Product{})
	mockData(db)
	return productRepositoryDB{db}
}

func (r productRepositoryDB) GetProducts() (products []model.Product, err error) {
	err = r.db.Order("quantity desc").Limit(30).Find(&products).Error
	return products, err

}