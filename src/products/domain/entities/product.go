package entities

type Product struct {
	ID    int32   `json:"id" gorm:"column:id_product;primaryKey;autoIncrement"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
