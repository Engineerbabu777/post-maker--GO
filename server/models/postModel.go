package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Caption  string             `json:"caption"`
	ImageURL string             `json:"image_url"`
	Likes    int                `json:"likes"`
	Shares   int                `json:"shares"`
}
