package ports


type Product struct{
	ID int
	Name string
	Quantity int
}


type ProductRepository interface{
	GetProducts()([]Product,error)
}

type ProductService interface{
	GetProducts()([]Product,error)
}