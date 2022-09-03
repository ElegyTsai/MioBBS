package post

import (
	"gorm.io/gorm"
	"time"
)

type CommonPost struct {
	PosterUID int64
	PostID    int64
	Content   string

	CreatedAt time.Time //和gorm有关
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
