package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Comment struct {
	ID        uuid.UUID  `gorm:"type:varchar(50);primaryKey" json:"id"`
	NickName  string     `gorm:"type:varchar(30)" json:"nick_name" validate:"required"`
	Content   string     `gorm:"type:text" json:"content" validate:"required"`
	ParentId  *uuid.UUID `gorm:"type:varchar(50);omitempty" json:"parent_id,omitempty"`
	ArticleID uuid.UUID  `gorm:"type:varchar(50)" json:"article_id"`
	CreatedAt time.Time  `json:"created_at"`
	Comments  []Comment  `gorm:"one2many:comments_comments;foreignKey:parent_id;AssociationForeignKey:id" json:"comments"`

	// Comments  []Comment  `gorm:"preload:true;foreignKey:parent_id;AssociationForeignKey:id" json:"comments"`
}

func (c *Comment) TableName() string {
	return "comments"
}

func (c *Comment) TimeStamp() {
	c.CreatedAt = time.Now()
}
