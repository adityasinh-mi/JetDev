package model

import (
	"jetdev-task/shared/database"

	_ "github.com/go-sql-driver/mysql"
)

func AutoMigrate() {
	conn := database.NewConnection()

	conn.GetDB().AutoMigrate(
		// For auto migrate database tables, need to add model below

		&Article{},
		&Comment{},
	)
	conn.GetDB().Model(&Comment{}).AddForeignKey("article_id", "articles(id)", "CASCADE", "CASCADE")
	conn.GetDB().Model(&Comment{}).AddForeignKey("parent_id", "comments(id)", "CASCADE", "CASCADE")

}
