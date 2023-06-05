package services

import ports "hexa-design/internal/ports/repository"


type productService struct{
	productRepo ports.ProductRepository
}


func NewProductService(productRepo ports.ProductRepository)ports.ProductService{
return productService{productRepo: productRepo}
}


func (s productService) GetProducts()(products []ports.Product, err error){
productDB ,err := s.productRepo.GetProducts()
if err != nil {
	return nil,err
}

for _, p := range productDB{
	products = append(products, ports.Product{
		ID: p.ID,
		Name: p.Name,
		Quantity: p.Quantity,
	})
}
return products,nil
}