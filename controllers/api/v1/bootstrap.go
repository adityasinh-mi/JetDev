package v1Ctl

import (
	v1Service "jetdev-task/services/api/v1"
	validator "jetdev-task/validator/api"
)

func ArticleController(validatorService validator.IAPIValidatorService, articleService v1Service.IArticleservice) *ArticleCtl {
	articleCtl := ArticleCtl{
		ArticleService: articleService,
		APIValidator:   validatorService,
	}

	return &articleCtl
}
func CommentController(validatorService validator.IAPIValidatorService, commentService v1Service.ICommentService) *CommentCtl {
	commentCtl := CommentCtl{
		CommentService: commentService,
		APIValidator:   validatorService,
	}

	return &commentCtl
}
