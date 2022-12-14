package v1Service

import (
	"errors"
	"jetdev-task/mocks"
	"jetdev-task/model"
	v1Request "jetdev-task/resources/api/request/v1"
	v1Response "jetdev-task/resources/api/response/v1"
	"jetdev-task/shared/utils"
	"jetdev-task/shared/utils/message"
	"net/http"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewArticleService(t *testing.T) {
	NewArticleService()
}

func TestArticleService_CreateArticle_Success(t *testing.T) {
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"res_code": http.StatusOK,
			"message":  message.ArticleCreated,
		},
	}

	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("CreateArticle", mock.Anything, mock.Anything).Return(model.Article{
		ID:       uuid.NewV1(),
		NickName: "RandomName",
	}, nil)

	var srvc ArticleService

	srvc.ArticleRepo = articleRepoMock

	var request v1Request.ArticleRequest
	data := srvc.CreateArticle(request)

	assert.Equal(t, response["meta"], data["meta"])
	assert.NotNil(t, data["data"])
}

func TestArticleService_CreateArticle_Error(t *testing.T) {
	response := map[string]interface{}{
		"res_code": http.StatusBadRequest,
		"message":  message.InvalidRequest,
	}

	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("CreateArticle", mock.Anything, mock.Anything).Return(model.Article{}, errors.New("error"))

	var srvc ArticleService

	srvc.ArticleRepo = articleRepoMock

	var request v1Request.ArticleRequest
	data := srvc.CreateArticle(request)

	assert.Equal(t, response, data)
}

func TestArticleService_GetArticleById_Success(t *testing.T) {
	articleResp := map[string]interface{}{
		"meta": map[string]interface{}{
			"res_code": 200,
			"message":  message.ArticleFetched,
		},
	}

	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("GetArticleById", mock.Anything, mock.Anything).Return(model.Article{
		ID:    uuid.NewV1(),
		Title: "TestTitle",
	}, nil)

	var articleId = uuid.NewV1()
	var srvc ArticleService
	srvc.ArticleRepo = articleRepoMock
	resp := srvc.GetArticleById(articleId)

	assert.Equal(t, resp["meta"].(map[string]interface{})["res_code"].(int), articleResp["meta"].(map[string]interface{})["res_code"].(int))
	assert.Equal(t, resp["meta"].(map[string]interface{})["message"].(string), articleResp["meta"].(map[string]interface{})["message"].(string))
}

func TestArticleService_GetArticleById_Err(t *testing.T) {
	articleResponse := map[string]interface{}{
		"res_code": 400,
		"message":  message.InvalidArticleId,
	}
	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("GetArticleById", mock.Anything, mock.Anything).Return(model.Article{}, errors.New("error"))

	var articleId uuid.UUID
	var srvc ArticleService
	srvc.ArticleRepo = articleRepoMock

	err := srvc.GetArticleById(articleId)
	assert.Equal(t, err, articleResponse)

}

func TestArticleService_GetArticles_WithData_Success(t *testing.T) {
	var page utils.PageAttr
	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"res_code":   http.StatusOK,
			"message":    message.ArticlesFetched,
			"pagination": page,
		},
	}

	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("GetArticles", mock.Anything, mock.Anything).Return([]v1Response.ArticlesResponse{}, nil)

	var srvc ArticleService

	srvc.ArticleRepo = articleRepoMock

	data := srvc.GetArticles(page)

	assert.Equal(t, response["meta"], data["meta"])
	assert.NotNil(t, data["data"])
}

func TestArticleService_GetArticles_Error(t *testing.T) {
	response := map[string]interface{}{
		"res_code": http.StatusBadRequest,
		"message":  message.InvalidPageInput,
	}

	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("GetArticles", mock.Anything, mock.Anything).Return([]v1Response.ArticlesResponse{}, errors.New(message.InvalidPageInput))

	var srvc ArticleService

	srvc.ArticleRepo = articleRepoMock

	var page utils.PageAttr
	data := srvc.GetArticles(page)

	assert.Equal(t, response, data)
}

func TestArticleService_GetArticleComments_Success(t *testing.T) {

	response := map[string]interface{}{
		"meta": map[string]interface{}{
			"res_code": http.StatusOK,
			"message":  message.CommentsFetched,
		},
	}

	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("GetArticleById", mock.Anything, mock.Anything).Return(model.Article{ID: uuid.NewV1()}, nil)
	articleRepoMock.On("GetArticleComments", mock.Anything, mock.Anything).Return([]model.Comment{}, nil)

	var articleId uuid.UUID
	var srvc ArticleService
	srvc.ArticleRepo = articleRepoMock
	data := srvc.GetArticleComments(articleId)
	assert.Equal(t, response["meta"], data["meta"])
	assert.NotNil(t, data["data"])
}

func TestArticleService_GetArticleComments_Error(t *testing.T) {

	response := map[string]interface{}{
		"res_code": http.StatusBadRequest,
		"message":  message.InvalidArticleId,
	}

	articleRepoMock := new(mocks.IArticleRepository)

	articleRepoMock.On("GetArticleById", mock.Anything, mock.Anything).Return(model.Article{}, errors.New("error"))
	articleRepoMock.On("GetArticleComments", mock.Anything, mock.Anything).Return([]model.Comment{}, nil)

	var articleId = uuid.NewV1()
	var srvc ArticleService
	srvc.ArticleRepo = articleRepoMock
	data := srvc.GetArticleComments(articleId)

	assert.Equal(t, response["meta"], data["meta"])
	assert.Nil(t, data["data"])
}
