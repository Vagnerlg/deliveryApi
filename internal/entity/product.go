package entity

type Product struct {
	ID          interface{} `json:"id" bson:"_id,omitempty"`
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description" binding:"required"`
	Price       uint        `json:"price" binding:"required"`
	Category    string      `json:"category"`
	Tags        []string    `json:"tags"`
	Pictures    []string    `json:"pictures"`
	Options     []Option    `json:"options"`
	Extra       []Extra     `json:"extra"`
}

type Option struct {
	Kind  Kind   `json:"kind" binding:"required"`
	Rules Rules  `json:"rules"`
	Items []Item `json:"items" binding:"required"`
}

type Kind string

const (
	Select Kind = "select"
)

type Rules struct {
	Min        uint       `json:"min" binding:"required"`
	Max        uint       `json:"max" binding:"required"`
	Calculated Calculated `json:"calculated" binding:"required"`
}

type Calculated string

const (
	ByQuota Calculated = "byQuota"
	ByUnit  Calculated = "ByUnit"
)

type Item struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Pictures    []string `json:"pictures"`
	Price       uint     `json:"price" binding:"required"`
}

type Extra struct {
	Kind  string `json:"kind" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ProductRespository interface {
	Create(*Product)
	List() []Product
	Get(string) Product
}
