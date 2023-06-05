package domain

type Product struct{
	ID int
	Name string
	Quantity int
}


type ProductRepository interface{
	GetProducts()([]Product,error)
}