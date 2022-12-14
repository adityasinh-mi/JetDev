package v1Ctl

import (
	v1req "jetdev-task/resources/api/request/v1"
	v1Service "jetdev-task/services/api/v1"
	u "jetdev-task/shared/common"
	"jetdev-task/shared/log"
	"jetdev-task/shared/utils"
	msg "jetdev-task/shared/utils/message"

	valid "jetdev-task/validator/api"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type ArticleCtl struct {
	ArticleService v1Service.IArticleservice
	APIValidator   valid.IAPIValidatorService
}

//CreateArticle is made for creating an article
// @router /api/v1/articles [post]
func (ac *ArticleCtl) CreateArticle(c *gin.Context) {
	log.GetLog().Info("INFO : ", "Article Controller Called(CreateArticle).")
	var req v1req.ArticleRequest

	//decode the request body into struct and failed if any error occur
	if err := c.BindJSON(&req); err != nil {
		log.GetLog().Info("ERROR : ", err.Error())
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidRequest))
		return
	}

	// Struct field validation
	if resp, ok := ac.APIValidator.ValidateStruct(req, "ArticleRequest"); !ok {
		log.GetLog().Info("ERROR : ", "Struct validation error")
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, resp))
		return
	}

	//call service
	resp := ac.ArticleService.CreateArticle(req)
	statusCode := u.GetHTTPStatusCode(resp["res_code"])

	//return response using api helper
	u.Respond(c.Writer, statusCode, resp)
	log.GetLog().Info("INFO : ", "Article Created Successfully.")

}

//GetArticleContent is made for getting an article contents
// @router /api/v1/articles/:articleId [get]
func (ac *ArticleCtl) GetArticlesContent(c *gin.Context) {
	log.GetLog().Info("INFO : ", "Article Controller Called(GetArticlesContent).")

	paramId := c.Param("articleId")
	articleId, err := uuid.FromString(paramId)
	if err != nil {
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidArticleId))
		return
	}

	//call service
	resp := ac.ArticleService.GetArticleById(articleId)
	statusCode := u.GetHTTPStatusCode(resp["res_code"])

	//return response using api helper
	u.Respond(c.Writer, statusCode, resp)
	log.GetLog().Info("INFO : ", "Article Fetched Successfully")

}

//GetArticles is made for Getting all the articles
// @router /api/v1/list-articles [get]
func (ac *ArticleCtl) GetArticles(c *gin.Context) {
	log.GetLog().Info("INFO : ", "Article Controller Called(GetArticles).")

	pageStr := c.DefaultQuery("page", "1")
	page, size, err := utils.PageAttributes(pageStr)
	if err != nil {
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidPageInput))
		return
	}

	pageAttr := utils.PageAttr{
		Page: page,
		Size: size,
	}

	//call service
	resp := ac.ArticleService.GetArticles(pageAttr)
	statusCode := u.GetHTTPStatusCode(resp["res_code"])

	//return response using api helper
	u.Respond(c.Writer, statusCode, resp)
	log.GetLog().Info("INFO : ", "Articles Fetched Successfully")

}

//GetArticleComments is made for getting all the comments that are posted on an article
// @router /api/v1/article/comments/:articleId
func (ac *ArticleCtl) GetArticleComments(c *gin.Context) {
	log.GetLog().Info("INFO : ", "Article Controller Called(GetArticleComments).")
	paramId := c.Param("articleId")

	articleId, err := uuid.FromString(paramId)
	if err != nil {
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidArticleId))
		return
	}

	//call service
	resp := ac.ArticleService.GetArticleComments(articleId)
	statusCode := u.GetHTTPStatusCode(resp["res_code"])

	//return response using api helper
	u.Respond(c.Writer, statusCode, resp)
	log.GetLog().Info("INFO : ", "Comments Fetched Successfully")

}
