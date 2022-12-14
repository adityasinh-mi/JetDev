package v1Ctl

import (
	v1req "jetdev-task/resources/api/request/v1"
	v1Service "jetdev-task/services/api/v1"
	u "jetdev-task/shared/common"
	"jetdev-task/shared/log"

	msg "jetdev-task/shared/utils/message"
	valid "jetdev-task/validator/api"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CommentCtl struct {
	CommentService v1Service.ICommentService
	APIValidator   valid.IAPIValidatorService
}

//PostCommentOnArticle is made for posting a comment on an article
// @router /api/v1/comment/:articleId
func (cc *CommentCtl) PostCommentOnArticle(c *gin.Context) {
	log.GetLog().Info("INFO : ", "Comment Controller Called(PostCommentOnArticle).")
	var req v1req.CommentRequest

	paramId := c.Param("articleId")

	articleId, err := uuid.FromString(paramId)
	if err != nil {
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidArticleId))
		return
	}

	//decode the request body into struct and failed if any error occur
	if err := c.BindJSON(&req); err != nil {
		log.GetLog().Info("ERROR : ", err.Error())
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidRequest))
		return
	}

	// Struct field validation
	if resp, ok := cc.APIValidator.ValidateStruct(req, "CommentRequest"); !ok {
		log.GetLog().Info("ERROR : ", "Struct validation error")
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, resp))
		return
	}

	//call service
	resp := cc.CommentService.PostCommentOnArticle(req, articleId)
	statusCode := u.GetHTTPStatusCode(resp["res_code"])

	//return response using api helper
	u.Respond(c.Writer, statusCode, resp)
	log.GetLog().Info("INFO : ", "Comment Posted Successfully.")
}

//PostCommentOnComment is made for posting a comment on an article's comment
// @router /api/v1//subcomment/:commentId
func (cc *CommentCtl) PostCommentOnComment(c *gin.Context) {
	log.GetLog().Info("INFO : ", "Comment Controller Called(PostCommentOnComment).")
	var req v1req.CommentRequest

	paramId := c.Param("commentId")

	commentId, err := uuid.FromString(paramId)
	if err != nil {
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidArticleId))
		return
	}

	//decode the request body into struct and failed if any error occur
	if err := c.BindJSON(&req); err != nil {
		log.GetLog().Info("ERROR : ", err.Error())
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, msg.InvalidRequest))
		return
	}

	// Struct field validation
	if resp, ok := cc.APIValidator.ValidateStruct(req, "CommentRequest"); !ok {
		log.GetLog().Info("ERROR : ", "Struct validation error")
		u.Respond(c.Writer, http.StatusBadRequest, u.ResponseErrorWithCode(u.CodeBadRequest, resp))
		return
	}

	//call service
	resp := cc.CommentService.PostCommentOnComment(req, commentId)
	statusCode := u.GetHTTPStatusCode(resp["res_code"])

	//return response using api helper
	u.Respond(c.Writer, statusCode, resp)
	log.GetLog().Info("INFO : ", "Comment Posted Successfully.")

}
