package eventsutils

import "go.mongodb.org/mongo-driver/bson/primitive"

type ItemCreatedContext struct {
	ID                primitive.ObjectID `json:"id" bson:"_id"`
	Seller            int64              `json:"seller" bson:"seller"`
	Title             string             `json:"title" bson:"title"`
	Price             float32            `json:"price" bson:"price"`
	AvailableQuantity int                `json:"available_quantity" bson:"available_quantity"`
	SoldQuantity      int                `json:"sold_quantity" bson:"sold_quantity"`
	Status            string             `json:"status" bson:"status"`
}
