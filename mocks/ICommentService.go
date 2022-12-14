// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import uuid "github.com/satori/go.uuid"
import v1Request "jetdev-task/resources/api/request/v1"

// ICommentService is an autogenerated mock type for the ICommentService type
type ICommentService struct {
	mock.Mock
}

// PostCommentOnArticle provides a mock function with given fields: req, articleId
func (_m *ICommentService) PostCommentOnArticle(req v1Request.CommentRequest, articleId uuid.UUID) map[string]interface{} {
	ret := _m.Called(req, articleId)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(v1Request.CommentRequest, uuid.UUID) map[string]interface{}); ok {
		r0 = rf(req, articleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// PostCommentOnComment provides a mock function with given fields: req, commentId
func (_m *ICommentService) PostCommentOnComment(req v1Request.CommentRequest, commentId uuid.UUID) map[string]interface{} {
	ret := _m.Called(req, commentId)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(v1Request.CommentRequest, uuid.UUID) map[string]interface{}); ok {
		r0 = rf(req, commentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}
