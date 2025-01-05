package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Skill struct {
	ID          primitive.ObjectID `json:"id"                    bson:"_id"`
	Name        string             `json:"name"                  bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
}

type SkillValuePair struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	SkillID primitive.ObjectID `json:"skill_id" bson:"skill_id"`
	Value   uint               `json:"value" bson:"value"`
}
