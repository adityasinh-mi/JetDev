package api

import (
	"context"
	"fmt"
	v1Ctl "jetdev-task/controllers/api/v1"
	v1Service "jetdev-task/services/api/v1"
	"jetdev-task/shared/config"
	"jetdev-task/shared/log"
	validator "jetdev-task/validator/api"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// IRoutes is
type IRoutes interface {
	Setup()
	Run()
	Close(ctx context.Context) error
}

// Routes is
type Routes struct {
	router *gin.Engine
	server *http.Server
	config config.IConfig

	articleCtl *v1Ctl.ArticleCtl
	commentCtl *v1Ctl.CommentCtl
}

// NewRouter is
func NewRouter(config config.IConfig) IRoutes {
	validation := validator.NewAPIValidatorService()
	articleSrv := v1Service.NewArticleService()
	articleCtl := v1Ctl.ArticleController(validation, articleSrv)
	commentSrv := v1Service.NewCommentService()
	commentCtl := v1Ctl.CommentController(validation, commentSrv)

	router := gin.Default()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.App().Port),
		Handler: router,
	}

	return &Routes{
		router,
		server,
		config,
		articleCtl,
		commentCtl,
	}
}

func (rt *Routes) Run() {
	log.GetLog().Info("", "service listen on "+rt.config.App().Port)

	err := rt.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.GetLog().Fatal("", "listen: %s\n", err)
	}
}

func (rt *Routes) Close(ctx context.Context) error {
	if rt.server != nil {
		return rt.server.Shutdown(ctx)
	}
	return nil
}

func (rt *Routes) Setup() {
	router := rt.router
	article := rt.articleCtl
	comment := rt.commentCtl

	rt.setupCors()
	rt.setupDefaultEndpoints()

	app := router.Group("/api/v1")

	app.GET("/articles/:articleId", article.GetArticlesContent)
	app.POST("/createArticle", article.CreateArticle)
	app.GET("/list-articles", article.GetArticles)
	app.POST("/comment/:articleId", comment.PostCommentOnArticle)
	app.POST("/subcomment/:commentId", comment.PostCommentOnComment)
	app.GET("/article/comments/:articleId", article.GetArticleComments)

}

func (rt *Routes) setupCors() {
	rt.router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		AllowAllOrigins: true,
		MaxAge:          12 * time.Hour,
	}))
}

func (rt *Routes) setupDefaultEndpoints() {
	rt.router.GET("/ping", func(c *gin.Context) {
		var msg string
		if rt.config.Env() == "production" {
			msg = fmt.Sprintf("Pong! I am %s. Version is %s.", rt.config.AppRegion(), rt.config.AppVersion())
		} else {
			msg = "pong"
		}
		c.JSON(200, gin.H{"message": msg})
	})

}
