package v1ORM

import (
	"database/sql"
	"fmt"
	"jetdev-task/model"
	v1Response "jetdev-task/resources/api/response/v1"
	"jetdev-task/shared/database"
	"jetdev-task/shared/log"
	"jetdev-task/shared/utils"
	"math"

	uuid "github.com/satori/go.uuid"
)

type IArticleRepository interface {
	CreateArticle(conn database.IConnection, request model.Article) (model.Article, error)
	GetArticleById(conn database.IConnection, articleId uuid.UUID) (model.Article, error)
	GetArticles(conn database.IConnection, paginate *utils.PageAttr) ([]v1Response.ArticlesResponse, error)
	GetArticleComments(conn database.IConnection, articleId uuid.UUID) ([]model.Comment, error)
}

type articleRepo struct {
	DB *sql.DB
}

func NewArticleWriter() IArticleRepository {
	return &articleRepo{}
}

func (ar *articleRepo) CreateArticle(conn database.IConnection, req model.Article) (model.Article, error) {
	log.GetLog().Info("INFO : ", "User Repo Called(SignUp).")

	result := conn.GetDB().Create(req)
	if result.Error != nil {
		return model.Article{}, result.Error
	}

	res, ok := result.Value.(model.Article)
	if !ok {
		return res, fmt.Errorf("Internal Server Error")
	}

	return res, nil

}

func (ar *articleRepo) GetArticleById(conn database.IConnection, articleId uuid.UUID) (model.Article, error) {
	log.GetLog().Info("INFO : ", "User Repo Called(GetArticleById).")

	var article model.Article
	err := conn.GetDB().Where("id = ?", articleId).First(&article).Error
	if err != nil {
		log.GetLog().Info("ERROR(query) : ", err.Error())
		return article, err
	}
	return article, nil

}

func (ar *articleRepo) GetArticles(conn database.IConnection, paginate *utils.PageAttr) ([]v1Response.ArticlesResponse, error) {
	log.GetLog().Info("INFO : ", "User Repo Called(GetArticles).")

	var articles []v1Response.ArticlesResponse
	var count int
	query := conn.GetDB().Table("articles").Select("*").Order(("created_at DESC"))

	query.Count(&count)

	query = query.Offset((paginate.Page - 1) * paginate.Size).Limit(paginate.Size)

	query.Find(&articles)

	paginate.Total = count

	lastPage := math.Ceil(float64(paginate.Total) / float64(paginate.Size))
	paginate.LastPage = int(lastPage)

	if paginate.Page > paginate.LastPage {
		return articles, fmt.Errorf("page is out of bound.")
	}

	return articles, nil

}

func (ar *articleRepo) GetArticleComments(conn database.IConnection, articleId uuid.UUID) ([]model.Comment, error) {
	log.GetLog().Info("INFO : ", "User Repo Called(GetArticleComments).")

	commentData, commentError := getParentComment(conn, articleId)
	if commentError != nil {
		return []model.Comment{}, commentError

	}
	return commentData, nil

}

func getParentComment(conn database.IConnection, articleId uuid.UUID) ([]model.Comment, error) {
	var parentComments []model.Comment
	rows, err := conn.GetDB().Raw("select * from comments where article_id = ? AND parent_id IS NULL ", articleId).Rows()
	if err != nil {
		return []model.Comment{}, err
	}

	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.ID, &comment.NickName, &comment.Content, &comment.ParentId, &comment.ArticleID, &comment.CreatedAt)
		if err != nil {
			return []model.Comment{}, err
		}

		parentComments = append(parentComments, comment)
	}

	for i, parentComment := range parentComments {

		comments, err := getChildrenComments(conn, parentComment.ID)
		if err != nil {
			return []model.Comment{}, err
		}

		parentComments[i].Comments = comments
	}

	return parentComments, nil
}

func getChildrenComments(conn database.IConnection, parentId uuid.UUID) ([]model.Comment, error) {

	rows, err := conn.GetDB().Raw("select * from comments where parent_id = ? ", parentId).Rows()
	if err != nil {
		return []model.Comment{}, err
	}

	var childComments []model.Comment
	for rows.Next() {
		var comment model.Comment
		err := rows.Scan(&comment.ID, &comment.NickName, &comment.Content, &comment.ParentId, &comment.ArticleID, &comment.CreatedAt)
		if err != nil {
			return []model.Comment{}, err
		}

		childComments = append(childComments, comment)
	}

	for i, childComment := range childComments {

		comments, err := getChildrenComments(conn, childComment.ID)
		if err != nil {
			return []model.Comment{}, err
		}

		childComments[i].Comments = comments
	}

	return childComments, nil

}
