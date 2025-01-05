package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID           primitive.ObjectID          `json:"id"                     bson:"_id"`
	UserID       primitive.ObjectID          `json:"user_id"                bson:"user_id"`
	Name         string                      `json:"name"                   bson:"name"`
	Level        uint                        `json:"level,omitempty"        bson:"level,omitempty"`
	Health       int                         `json:"health,omitempty"       bson:"health,omitempty"`
	MentalHealth int                         `json:"mentalhealth,omitempty" bson:"mentalhealth,omitempty"`
	Mana         uint                        `json:"mana,omitempty"         bson:"mana,omitempty"`
	Race         string                      `json:"race,omitempty"         bson:"race,omitempty"`
	Gender       string                      `json:"gender,omitempty"       bson:"gender,omitempty"`
	Height       string                      `json:"height,omitempty"       bson:"height,omitempty"`
	Weight       string                      `json:"weight,omitempty"       bson:"weight,omitempty"`
	Dodge        uint                        `json:"dodge,omitempty"        bson:"dodge,omitempty"`
	Skills       map[string][]SkillValuePair `json:"skills,omitempty"       bson:"skills,omitempty"`
}
