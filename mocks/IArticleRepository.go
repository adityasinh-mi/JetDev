// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import database "jetdev-task/shared/database"
import mock "github.com/stretchr/testify/mock"
import model "jetdev-task/model"
import utils "jetdev-task/shared/utils"
import uuid "github.com/satori/go.uuid"

import v1Response "jetdev-task/resources/api/response/v1"

// IArticleRepository is an autogenerated mock type for the IArticleRepository type
type IArticleRepository struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: conn, request
func (_m *IArticleRepository) CreateArticle(conn database.IConnection, request model.Article) (model.Article, error) {
	ret := _m.Called(conn, request)

	var r0 model.Article
	if rf, ok := ret.Get(0).(func(database.IConnection, model.Article) model.Article); ok {
		r0 = rf(conn, request)
	} else {
		r0 = ret.Get(0).(model.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(database.IConnection, model.Article) error); ok {
		r1 = rf(conn, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticleById provides a mock function with given fields: conn, articleId
func (_m *IArticleRepository) GetArticleById(conn database.IConnection, articleId uuid.UUID) (model.Article, error) {
	ret := _m.Called(conn, articleId)

	var r0 model.Article
	if rf, ok := ret.Get(0).(func(database.IConnection, uuid.UUID) model.Article); ok {
		r0 = rf(conn, articleId)
	} else {
		r0 = ret.Get(0).(model.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(database.IConnection, uuid.UUID) error); ok {
		r1 = rf(conn, articleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticleComments provides a mock function with given fields: conn, articleId
func (_m *IArticleRepository) GetArticleComments(conn database.IConnection, articleId uuid.UUID) ([]model.Comment, error) {
	ret := _m.Called(conn, articleId)

	var r0 []model.Comment
	if rf, ok := ret.Get(0).(func(database.IConnection, uuid.UUID) []model.Comment); ok {
		r0 = rf(conn, articleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(database.IConnection, uuid.UUID) error); ok {
		r1 = rf(conn, articleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArticles provides a mock function with given fields: conn, paginate
func (_m *IArticleRepository) GetArticles(conn database.IConnection, paginate *utils.PageAttr) ([]v1Response.ArticlesResponse, error) {
	ret := _m.Called(conn, paginate)

	var r0 []v1Response.ArticlesResponse
	if rf, ok := ret.Get(0).(func(database.IConnection, *utils.PageAttr) []v1Response.ArticlesResponse); ok {
		r0 = rf(conn, paginate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1Response.ArticlesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(database.IConnection, *utils.PageAttr) error); ok {
		r1 = rf(conn, paginate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
