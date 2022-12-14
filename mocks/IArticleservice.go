// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import utils "jetdev-task/shared/utils"
import uuid "github.com/satori/go.uuid"
import v1Request "jetdev-task/resources/api/request/v1"

// IArticleservice is an autogenerated mock type for the IArticleservice type
type IArticleservice struct {
	mock.Mock
}

// CreateArticle provides a mock function with given fields: req
func (_m *IArticleservice) CreateArticle(req v1Request.ArticleRequest) map[string]interface{} {
	ret := _m.Called(req)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(v1Request.ArticleRequest) map[string]interface{}); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GetArticleById provides a mock function with given fields: articleId
func (_m *IArticleservice) GetArticleById(articleId uuid.UUID) map[string]interface{} {
	ret := _m.Called(articleId)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(uuid.UUID) map[string]interface{}); ok {
		r0 = rf(articleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GetArticleComments provides a mock function with given fields: articleId
func (_m *IArticleservice) GetArticleComments(articleId uuid.UUID) map[string]interface{} {
	ret := _m.Called(articleId)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(uuid.UUID) map[string]interface{}); ok {
		r0 = rf(articleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GetArticles provides a mock function with given fields: page
func (_m *IArticleservice) GetArticles(page utils.PageAttr) map[string]interface{} {
	ret := _m.Called(page)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(utils.PageAttr) map[string]interface{}); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}
