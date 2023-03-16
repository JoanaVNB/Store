package domain

type Order struct{
	ID string `json:"id" gorm:"varchar(150); not null"`
	User_id string `json:"user_id" gorm:"varchar(150); not null"`
	Item_description string `json:"item_description" gorm:"varchar(30); not null"`
	Item_quantity int `json:"item_quantity" gorm:"int; not null"`
	Item_price int `json:"item_price" gorm:"int; not null"`
	Total_value int `json:"total_value" gorm:"int; not null"`
}

func (Order) TableName() string{
	return "order"
}
