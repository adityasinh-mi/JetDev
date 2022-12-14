package v1Service

import (
	"errors"
	"jetdev-task/mocks"
	"jetdev-task/model"
	v1Request "jetdev-task/resources/api/request/v1"
	"jetdev-task/shared/utils/message"
	"net/http"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewCommentService(t *testing.T) {
	NewCommentService()
}

func TestCommentService_PostCommentOnArticle_Success(t *testing.T) {

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"res_code": http.StatusOK,
			"message":  message.CommentSuccess,
		},
	}

	articleRepoMock := new(mocks.IArticleRepository)
	commentRepoMock := new(mocks.ICommentRepository)

	articleRepoMock.On("GetArticleById", mock.Anything, mock.Anything).Return(model.Article{}, nil)
	commentRepoMock.On("PostCommentOnArticle", mock.Anything, mock.Anything).Return(model.Comment{}, nil)

	var srvc commentService
	srvc.CommentRepo = commentRepoMock
	srvc.ArticleRepo = articleRepoMock

	var articleId uuid.UUID
	var request v1Request.CommentRequest
	data := srvc.PostCommentOnArticle(request, articleId)

	assert.Equal(t, response["meta"], data["meta"])
	assert.NotNil(t, data["data"])
}

func TestCommentService_PostCommentOnArticle_Err(t *testing.T) {

	response := map[string]interface{}{
		"res_code": http.StatusBadRequest,
		"message":  message.InvalidArticleId,
	}

	articleRepoMock := new(mocks.IArticleRepository)
	commentRepoMock := new(mocks.ICommentRepository)

	articleRepoMock.On("GetArticleById", mock.Anything, mock.Anything).Return(model.Article{}, errors.New("error"))
	commentRepoMock.On("PostCommentOnArticle", mock.Anything, mock.Anything).Return(model.Comment{}, nil)

	var srvc commentService
	srvc.CommentRepo = commentRepoMock
	srvc.ArticleRepo = articleRepoMock

	var articleId uuid.UUID
	var request v1Request.CommentRequest
	data := srvc.PostCommentOnArticle(request, articleId)

	assert.Equal(t, response["meta"], data["meta"])
	assert.Nil(t, data["data"])
}

func TestCommentService_PostCommentOnComment_Success(t *testing.T) {

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"res_code": http.StatusOK,
			"message":  message.CommentSuccess,
		},
		"data": model.Comment{
			ID:       uuid.NewV1(),
			NickName: "MetaComment",
		},
	}

	commentRepoMock := new(mocks.ICommentRepository)

	commentRepoMock.On("GetCommentById", mock.Anything, mock.Anything).Return(model.Comment{ID: uuid.NewV1()}, nil)
	commentRepoMock.On("PostCommentOnArticle", mock.Anything, mock.Anything).Return(model.Comment{}, nil)

	var srvc commentService
	srvc.CommentRepo = commentRepoMock

	var commentId uuid.UUID
	var request v1Request.CommentRequest
	data := srvc.PostCommentOnComment(request, commentId)

	assert.Equal(t, response["meta"], data["meta"])
	assert.NotNil(t, data["data"])
}

func TestCommentService_PostCommentOnComment_Err(t *testing.T) {

	response := map[string]interface{}{
		"res_code": http.StatusBadRequest,
		"message":  message.InvalidCommentId,
	}

	commentRepoMock := new(mocks.ICommentRepository)

	commentRepoMock.On("GetCommentById", mock.Anything, mock.Anything).Return(model.Comment{ID: uuid.Nil}, errors.New("error"))
	commentRepoMock.On("PostCommentOnArticle", mock.Anything, mock.Anything).Return(model.Comment{}, nil)

	var srvc commentService
	srvc.CommentRepo = commentRepoMock

	var commentId uuid.UUID
	var request v1Request.CommentRequest
	data := srvc.PostCommentOnComment(request, commentId)

	assert.Equal(t, response["meta"], data["meta"])
	assert.Nil(t, data["data"])
}
