package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Article struct {
	ID        uuid.UUID `gorm:"type:varchar(50);primary_key" json:"id"`
	NickName  string    `gorm:"type:varchar(30)" json:"nick_name" validate:"required"`
	Title     string    `gorm:"type:varchar(30)" json:"title" validate:"required"`
	Content   string    `gorm:"type:text" json:"content" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	Comment   []Comment `json:"comment,omitempty"gorm:"foreignKey:parent_id;AssociationForeignKey:id"`
}

func (a *Article) TableName() string {
	return "articles"
}

func (a *Article) TimeStamp() {
	a.CreatedAt = time.Now()
}
