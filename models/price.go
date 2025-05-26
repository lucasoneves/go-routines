package models

import "time"

type ProductDetail struct {
	StorageName string
	Value       float64
	TimeStamp   time.Time
}
