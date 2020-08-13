package model

import "time"

// Todo モデル
type Todo struct {
	ID        uint64 `gorm:"primary_key"`
	Title     string
	Finished  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
