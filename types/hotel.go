package types

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	_ RoomType = iota
	Single
	Double
	Deluxe
)

type Hotel struct {
	ID       primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
}

type RoomType int

type Room struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      RoomType           `json:"type" bson:"type"`
	BasePrice float64            `json:"basePrice" bson:"basePrice"`
}
