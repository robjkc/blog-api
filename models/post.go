package models

import (
	"github.com/robjkc/blog-api/db"
)

type Post struct {
	ID      int    `db:"id"`
	Title   string `db:"title"`
	Author  string `db:"author"`
	Content string `db:"content"`
}

func GetPosts(con *db.DbConnection) ([]Post, error) {

	posts := []Post{}
	err := con.ExecuteSelect(&posts, `select id,
		title,
		author,
		content
		from posts`, db.Args{})
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func AddPost(con *db.DbConnection, author string, title string, content string) error {
	err := con.ExecuteUpdate(`insert into posts (author, title, content) values (:author, :title, :content)`,
		db.Args{"author": author, "title": title, "content": content})
	if err != nil {
		return err
	}

	return nil
}
