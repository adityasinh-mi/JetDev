package v1Response

import (
	"jetdev-task/model"
	"time"

	uuid "github.com/satori/go.uuid"
)

type ArticleContentResponse struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}

type ArticlesResponse struct {
	ID        uuid.UUID `json:"id"`
	NickName  string    `json:"nick_name"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticlesComments struct {
	ID        uuid.UUID     `gorm:"type:varchar(50);primary_key"`
	NickName  string        `gorm:"varchar(30)" json:"nick_name" validate:"required"`
	Title     string        `gorm:"varchar(30)" json:"title" validate:"required"`
	Content   string        `gorm:"varchar(255)" json:"content" validate:"required"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Comments  []CommentResp `json:"comments"`
}

// type CommentResp struct {
// 	ID        uuid.UUID     `gorm:"type:varchar(50);primaryKey" json:"id"`
// 	NickName  string        `gorm:"varchar(30)" json:"nick_name" validate:"required"`
// 	Content   string        `gorm:"text" json:"content" validate:"required"`
// 	ParentId  uuid.UUID     `gorm:"type:varchar(50)"`
// 	ArticleID uuid.UUID     `gorm:"type:varchar(50)" json:"article_id"`
// 	CreatedAt time.Time     `json:"created_at"`
// 	Comments  []CommentResp `json:"comments"`
// }

type CommentResp struct {
	ID        uuid.UUID       `gorm:"type:varchar(50);primaryKey" json:"id"`
	NickName  string          `gorm:"varchar(30)" json:"nick_name" validate:"required"`
	Content   string          `gorm:"text" json:"content" validate:"required"`
	ParentId  *uuid.UUID      `gorm:"type:varchar(50)" json:"parent_id,omitempty"`
	ArticleID uuid.UUID       `gorm:"type:varchar(50)" json:"article_id"`
	CreatedAt time.Time       `json:"created_at"`
	Comment   []model.Comment `json:"comment,omitempty"gorm:"foreignKey:parent_id;AssociationForeignKey:id;omitempty"`
}
