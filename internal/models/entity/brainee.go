package entity

import "time"

type Brainee struct {
	Id        int       `gorm:"primaryKey;column:id"`
	Text      string    `gorm:"column:text"`
	Author    string    `gorm:"column:author"`
	Brand     string    `gorm:"column:brand"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
