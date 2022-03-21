package models

import (
	"time"

	"github.com/robjkc/blog-api/db"
)

type Comment struct {
	ID              int       `db:"id"`
	PostID          int       `db:"post_id"`
	TopLevel        int       `db:"top_level"`
	ParentCommentID int       `db:"parent_comment_id"`
	Author          string    `db:"author"`
	Content         string    `db:"content"`
	CreateDate      time.Time `db:"create_date"`
}

func GetComments(con *db.DbConnection, postId int) ([]Comment, error) {

	comments := []Comment{}
	err := con.ExecuteSelect(&comments, `select c.id,
		c.post_id,
		c.top_level,
		cc.parent_comment_id,
		c.author,
		c.content,
		c.create_date
		from comments c join child_comments cc on c.id = cc.child_comment_id
		where post_id = :postId order by c.top_level desc, cc.parent_comment_id, c.create_date`, db.Args{"postId": postId})
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func AddComment(con *db.DbConnection, postId int, parentCommentId int, author string, content string) error {

	topLevel := 0
	if parentCommentId == 0 {
		// No parent so this is a top level comment.
		topLevel = 1
	}

	err := con.ExecuteUpdate(`insert into comments (post_id, author, content, top_level) values (:postId, :author, :content, :topLevel)`,
		db.Args{"postId": postId, "author": author, "content": content, "topLevel": topLevel})
	if err != nil {
		return err
	}

	if parentCommentId == 0 {
		// No parent comment id not provided so just use the newly inserted comment id as the parent.
		err = con.ExecuteUpdate(`insert into child_comments (parent_comment_id, child_comment_id) values (currval('comments_id_seq'), currval('comments_id_seq'))`,
			db.Args{})
		if err != nil {
			return err
		}
	} else {
		// A parent comment id was provided so just use it as the parent comment id.
		err = con.ExecuteUpdate(`insert into child_comments (parent_comment_id, child_comment_id) values (:parentCommentId, currval('comments_id_seq'))`,
			db.Args{"parentCommentId": parentCommentId})
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteComment(con *db.DbConnection, commentId int) error {
	err := con.ExecuteUpdate(`delete from comments where id = :commentId`,
		db.Args{"commentId": commentId})
	if err != nil {
		return err
	}

	return nil
}

func UpdateComment(con *db.DbConnection, id int, author string, content string) error {
	err := con.ExecuteUpdate(`update comments set author = :author, content = :content where id = :id`,
		db.Args{"id": id, "author": author, "content": content})
	if err != nil {
		return err
	}

	return nil
}
