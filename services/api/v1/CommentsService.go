package v1Service

import (
	"jetdev-task/model"
	v1repo "jetdev-task/repository/api/orm/v1"
	v1req "jetdev-task/resources/api/request/v1"
	u "jetdev-task/shared/common"
	"jetdev-task/shared/database"
	"jetdev-task/shared/log"
	msg "jetdev-task/shared/utils/message"

	"net/http"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//ICommentService interface for comment services
type ICommentService interface {
	PostCommentOnArticle(req v1req.CommentRequest, articleId uuid.UUID) map[string]interface{}
	PostCommentOnComment(req v1req.CommentRequest, commentId uuid.UUID) map[string]interface{}
}

type commentService struct {
	CommentRepo v1repo.ICommentRepository
	ArticleRepo v1repo.IArticleRepository
}

func NewCommentService() ICommentService {
	commentRepo := v1repo.NewCommentWriter()
	articleRepo := v1repo.NewArticleWriter()
	return &commentService{
		CommentRepo: commentRepo,
		ArticleRepo: articleRepo,
	}
}

//PostCommentOnArticle is a service used for posting a comment on an article
func (cs *commentService) PostCommentOnArticle(req v1req.CommentRequest, articleId uuid.UUID) map[string]interface{} {
	log.GetLog().Info("INFO : ", "Comment Service Called(PostCommentOnArticle).")
	conn := database.NewConnection()

	//calling article repo
	resp, err := cs.ArticleRepo.GetArticleById(conn, articleId)
	if err == gorm.ErrRecordNotFound && resp.ID == uuid.Nil {
		return u.ResponseErrorWithCode(http.StatusBadRequest, msg.InvalidArticleId)
	}
	if err != nil {
		return u.ResponseErrorWithCode(http.StatusInternalServerError, msg.InternalServer)
	}

	//request to comment model
	var comment model.Comment
	comment.ID = uuid.NewV1()
	comment.NickName = req.NickName
	comment.ParentId = nil
	comment.Content = req.Content
	comment.ArticleID = articleId
	comment.TimeStamp()

	//calling comment repo
	commentResp, err := cs.CommentRepo.PostCommentOnArticle(conn, comment)
	if err != nil {
		log.GetLog().Info("ERROR(from repo) : ", err.Error())
		return u.ResponseErrorWithCode(http.StatusInternalServerError, msg.InternalServer)
	}
	response := u.ResponseSuccessWithArray(msg.CommentSuccess, commentResp)
	return response

}

//PostCommentOnComment is a service used for posting a comment on an article's comment
func (cs *commentService) PostCommentOnComment(req v1req.CommentRequest, commentId uuid.UUID) map[string]interface{} {
	log.GetLog().Info("INFO : ", "Comment Service Called(PostCommentOnComment).")
	conn := database.NewConnection()

	//calling comment repo
	resp, err := cs.CommentRepo.GetCommentById(conn, commentId)
	if err == gorm.ErrRecordNotFound && resp.ID == uuid.Nil {
		return u.ResponseErrorWithCode(http.StatusBadRequest, msg.InvalidCommentId)
	}
	if err != nil {
		return u.ResponseErrorWithCode(http.StatusInternalServerError, msg.InternalServer)
	}

	//request to comment model
	var comment model.Comment
	comment.ID = uuid.NewV1()
	comment.NickName = req.NickName
	comment.Content = req.Content
	comment.ParentId = &resp.ID
	comment.ArticleID = resp.ArticleID
	comment.TimeStamp()

	//calling comment repo
	commentResp, err := cs.CommentRepo.PostCommentOnArticle(conn, comment)
	if err != nil {
		log.GetLog().Info("ERROR(from repo) : ", err.Error())
		return u.ResponseErrorWithCode(http.StatusInternalServerError, msg.InternalServer)
	}
	response := u.ResponseSuccessWithArray(msg.CommentSuccess, commentResp)
	return response

}
