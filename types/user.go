package types

type User struct {
	ID        string `bson:"_id" json:"id"`
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
}
