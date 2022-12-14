package v1Request

type ArticleRequest struct {
	NickName string `gorm:"varchar(30)" json:"nick_name" validate:"required"`
	Title    string `gorm:"varchar(30)" json:"title" validate:"required"`
	Content  string `gorm:"varchar(255)" json:"content" validate:"required"`
}

type CommentRequest struct {
	NickName string `json:"nick_name" validate:"required"`
	Content  string `json:"content" validate:"required"`
}
