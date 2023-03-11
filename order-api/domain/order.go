package domain

type Order struct{
	ID string 
	User_id string `json:"user_id"`
	Item_description string `json:"item_description"`
	Item_quantity int `json:"item_quantity"`
	Item_price int `json:"item_price"`
	Total_value int `json:"total_value"`
}
