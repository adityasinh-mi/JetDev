package v1Service

import (
	"jetdev-task/model"
	v1repo "jetdev-task/repository/api/orm/v1"
	v1req "jetdev-task/resources/api/request/v1"
	u "jetdev-task/shared/common"
	"jetdev-task/shared/database"
	"jetdev-task/shared/log"
	"jetdev-task/shared/utils"
	msg "jetdev-task/shared/utils/message"

	"net/http"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type IArticleservice interface {
	CreateArticle(req v1req.ArticleRequest) map[string]interface{}
	GetArticleById(articleId uuid.UUID) map[string]interface{}
	GetArticles(page utils.PageAttr) map[string]interface{}
	GetArticleComments(articleId uuid.UUID) map[string]interface{}
}

type ArticleService struct {
	ArticleRepo v1repo.IArticleRepository
}

func NewArticleService() IArticleservice {
	articleRepo := v1repo.NewArticleWriter()
	return &ArticleService{
		ArticleRepo: articleRepo,
	}
}

//CreateArticle is a service used for creating an article
func (as *ArticleService) CreateArticle(req v1req.ArticleRequest) map[string]interface{} {

	log.GetLog().Info("INFO : ", "Article Service Called(CreateArticle).")
	conn := database.NewConnection()
	var article model.Article

	//adding the request data to article model
	article.Content = req.Content
	article.NickName = req.NickName
	article.Title = req.Title
	article.TimeStamp()

	article.ID = uuid.NewV1()

	//Call article repository
	resp, err := as.ArticleRepo.CreateArticle(conn, article)
	if err != nil {
		log.GetLog().Info("ERROR(from repo) : ", err.Error())
		return u.ResponseErrorWithCode(http.StatusBadRequest, msg.InvalidRequest)
	}

	response := u.ResponseSuccessWithArray(msg.ArticleCreated, resp)
	return response

}

//GetArticleById is a service used for getting the contents of an article by it articleId
func (as *ArticleService) GetArticleById(ArticleId uuid.UUID) map[string]interface{} {
	log.GetLog().Info("INFO : ", "Article Service Called(GetArticleById).")

	conn := database.NewConnection()

	//call article repository
	resp, err := as.ArticleRepo.GetArticleById(conn, ArticleId)
	if err != nil {
		return u.ResponseErrorWithCode(http.StatusBadRequest, msg.InvalidArticleId)
	}

	response := u.ResponseSuccessWithArray(msg.ArticleFetched, resp)
	return response

}

//GetArticles is a service used for getting  the list of all the articles
func (as *ArticleService) GetArticles(page utils.PageAttr) map[string]interface{} {
	log.GetLog().Info("INFO : ", "Article Service Called(GetArticles).")

	conn := database.NewConnection()

	//calling article repository
	resp, err := as.ArticleRepo.GetArticles(conn, &page)
	if err != nil {
		return u.ResponseErrorWithCode(http.StatusBadRequest, err.Error())
	}

	response := u.ResponseSuccessWithArray(msg.ArticlesFetched, resp)
	response["meta"].(map[string]interface{})["pagination"] = page
	return response

}

//GetArticleComments is a service used for getting all the comments which have been posted on an article
func (as *ArticleService) GetArticleComments(articleId uuid.UUID) map[string]interface{} {
	log.GetLog().Info("INFO : ", "Article Service Called(GetArticleComments).")

	conn := database.NewConnection()

	//calling article repo
	resp, err := as.ArticleRepo.GetArticleById(conn, articleId)
	if resp.ID == uuid.Nil || err == gorm.ErrRecordNotFound {
		return u.ResponseErrorWithCode(http.StatusBadRequest, msg.InvalidArticleId)
	}
	if err != nil {
		return u.ResponseErrorWithCode(http.StatusInternalServerError, msg.InternalServer)
	}

	//calling article repo
	commentsRep, err := as.ArticleRepo.GetArticleComments(conn, articleId)
	if err != nil {
		return u.ResponseErrorWithCode(http.StatusInternalServerError, err.Error())
	}

	if resp.ID == uuid.Nil && err == gorm.ErrRecordNotFound {
		return u.ResponseErrorWithCode(http.StatusInternalServerError, msg.InternalServer)
	}

	response := u.ResponseSuccessWithArray(msg.CommentsFetched, commentsRep)
	return response
}
