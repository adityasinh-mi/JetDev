package v1ORM

import (
	"database/sql"
	"fmt"
	"jetdev-task/model"
	"jetdev-task/shared/database"
	"jetdev-task/shared/log"
	"jetdev-task/shared/utils/message"

	uuid "github.com/satori/go.uuid"
)

type ICommentRepository interface {
	PostCommentOnArticle(conn database.IConnection, req model.Comment) (model.Comment, error)
	GetCommentById(conn database.IConnection, commentId uuid.UUID) (model.Comment, error)
}

type commentRepo struct {
	DB *sql.DB
}

func NewCommentWriter() ICommentRepository {
	return &commentRepo{}
}

func (cr *commentRepo) PostCommentOnArticle(conn database.IConnection, req model.Comment) (model.Comment, error) {
	log.GetLog().Info("INFO : ", "Comment Repo Called(PostCommentOnArticle).")

	result := conn.GetDB().Create(req)
	if result.Error != nil {
		return model.Comment{}, result.Error
	}

	res, ok := result.Value.(model.Comment)
	if !ok {
		return res, fmt.Errorf(message.InternalServer)
	}

	return res, nil

}

func (cr *commentRepo) GetCommentById(conn database.IConnection, commentId uuid.UUID) (model.Comment, error) {
	log.GetLog().Info("INFO : ", "User Repo Called(SignUp).")

	var commentData model.Comment
	err := conn.GetDB().Where("id = ?", commentId).First(&commentData).Error
	if err != nil {
		log.GetLog().Info("ERROR(query) : ", err.Error())
		return commentData, err
	}
	return commentData, nil

}
