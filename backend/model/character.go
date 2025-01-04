package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID           primitive.ObjectID `json:"id"                     bson:"_id"`
	Name         string             `json:"name,omitempty"         bson:"name,omitempty"`
	Level        int                `json:"level,omitempty"        bson:"level,omitempty"`
	Health       int                `json:"health,omitempty"       bson:"health,omitempty"`
	MentalHealth int                `json:"mentalhealth,omitempty" bson:"mentalhealth,omitempty"`
	Race         string             `json:"race,omitempty"         bson:"race,omitempty"`
	Gender       string             `json:"gender,omitempty"       bson:"gender,omitempty"`
	Height       string             `json:"height,omitempty"       bson:"height,omitempty"`
	Weight       string             `json:"weight,omitempty"       bson:"weight,omitempty"`
	Dodge        int                `json:"dodge,omitempty"        bson:"dodge,omitempty"`
	UserID       primitive.ObjectID `json:"user_id"                bson:"user_id"`
}
