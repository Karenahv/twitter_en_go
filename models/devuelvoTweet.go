package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempy"`
	UserID  string             `bson:"userid" json:"userId,omitempy"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempy"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempy"`
}
