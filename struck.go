package backendgcf

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	// Email		 string             	`bson:"email,omitempty" json:"email,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type RegisterStruct struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	// PhoneNumber string `json:"phone_number" bson:"phone_number"`
	// Email       string `json:"email" bson:"email"`
}
