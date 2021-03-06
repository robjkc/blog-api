package models

import (
	"time"

	"github.com/robjkc/blog-api/db"
)

type Post struct {
	ID         int       `db:"id"`
	Title      string    `db:"title"`
	Author     string    `db:"author"`
	Content    string    `db:"content"`
	CreateDate time.Time `db:"create_date"`
}

func GetPosts(con *db.DbConnection) ([]Post, error) {

	posts := []Post{}
	err := con.ExecuteSelect(&posts, `select id,
		title,
		author,
		content,
		create_date
		from posts`, db.Args{})
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func GetPost(con *db.DbConnection, id int) (Post, error) {

	posts := []Post{}
	err := con.ExecuteSelect(&posts, `select id,
		title,
		author,
		content,
		create_date
		from posts where id = :id`, db.Args{"id": id})
	if err != nil {
		return Post{}, err
	}

	return posts[0], nil
}

func AddPost(con *db.DbConnection, author string, title string, content string) error {
	err := con.ExecuteUpdate(`insert into posts (author, title, content) values (:author, :title, :content)`,
		db.Args{"author": author, "title": title, "content": content})
	if err != nil {
		return err
	}

	return nil
}

func UpdatePost(con *db.DbConnection, id int, author string, title string, content string) error {
	err := con.ExecuteUpdate(`update posts set author = :author, title = :title, content = :content where id = :id`,
		db.Args{"id": id, "author": author, "title": title, "content": content})
	if err != nil {
		return err
	}

	return nil
}

func DeletePost(con *db.DbConnection, postId int) error {
	err := con.ExecuteUpdate(`delete from posts where id = :postId`,
		db.Args{"postId": postId})
	if err != nil {
		return err
	}

	return nil
}
