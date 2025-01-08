package entity

type Product struct {
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       uint        `json:"price" binding:"required"`
}

type ProductRespository interface {
	Create(*Product)
	List() []Product
	Get(string) Product
}
