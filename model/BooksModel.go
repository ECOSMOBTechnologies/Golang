package model

type Books struct {
	Key string `json:"key" bson:"key"`
	Val int    `json:"val" bson:"val"`
}
