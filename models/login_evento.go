package models

import "time"

type LoginEvento struct {
	Usuario   string    `bson:"usuario"`
	TimeStamp time.Time `bson:"time_stamp"`
	Status    string    `bson:"status"`
}
