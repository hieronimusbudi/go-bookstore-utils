package eventsutils

type ItemCreatedContext struct {
	ID                string  `json:"id"`
	Seller            int64   `json:"seller"`
	Title             string  `json:"title"`
	Price             float32 `json:"price"`
	AvailableQuantity int     `json:"available_quantity"`
	SoldQuantity      int     `json:"sold_quantity"`
	Status            string  `json:"status"`
}
