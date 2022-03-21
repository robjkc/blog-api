package models

import (
	"fmt"

	"github.com/robjkc/blog-api/db"
)

type Comment struct {
	ID              int    `db:"id"`
	PostID          int    `db:"post_id"`
	ParentCommentID int    `db:"parent_comment_id"`
	Author          string `db:"author"`
	Content         string `db:"content"`
	Level           int    `db:"level"`
}

func GetComments(con *db.DbConnection, postId int) ([]Comment, error) {

	comments := []Comment{}
	err := con.ExecuteSelect(&comments, `select c.id,
		c.post_id,
		cc.parent_comment_id,
		c.author,
		c.content,
		c.level
		from comments c join child_comments cc on c.id = cc.child_comment_id
		where post_id = :postId order by c.level, c.create_date`, db.Args{"postId": postId})
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func AddComment(con *db.DbConnection, postId int, parentCommentId int, author string, content string) error {

	if parentCommentId == 0 {
		err := con.ExecuteUpdate(`insert into comments (post_id, author, content, level) values (:postId, :author, :content, 1)`,
			db.Args{"postId": postId, "author": author, "content": content})
		if err != nil {
			return err
		}
		// No parent comment id provided so just use the newly inserted comment id as the parent.
		err = con.ExecuteUpdate(`insert into child_comments (parent_comment_id, child_comment_id) values (currval('comments_id_seq'), currval('comments_id_seq'))`,
			db.Args{})
		if err != nil {
			return err
		}
	} else {
		fmt.Println("test!")

		err := con.ExecuteUpdate(`insert into comments (post_id, author, content, level) select :postId, :author, :content, level + 1 from comments where id = :parentCommentId`,
			db.Args{"postId": postId, "author": author, "content": content, "parentCommentId": parentCommentId})
		if err != nil {
			fmt.Println("fail!")
			return err
		}
		// A parent comment id was provided so just use it as the parent comment id.
		err = con.ExecuteUpdate(`insert into child_comments (parent_comment_id, child_comment_id) values (:parentCommentId, currval('comments_id_seq'))`,
			db.Args{"parentCommentId": parentCommentId})
		if err != nil {
			return err
		}
	}

	return nil
}
