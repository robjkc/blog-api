# blog-api

**Steps to run**

***1. Start a preloaded postgres database using docker-compose ***


***Docker Compose Steps (from root directory):***


```
docker-compose up
```

**Example Endpoints**
```
# Get All Posts
GET http://localhost:8080/posts

# Get Single Post with threaded comments
GET http://localhost:8080/posts/1

# Add a new post
POST http://localhost:8080/posts
Body:
{
	"author": "Robert",
	"title": "Test",
	"content": "Testing "
}

# Add a new comment
POST http://localhost:8080/comments
Body:
{
	"post_id": 1,
	"author": "Robert",
	"content": "Test Comment"
}

# Add a new threaded comment
POST http://localhost:8080/comments
Body:
{
	"post_id": 1,
	"parent_comment_id": 2,
	"author": "Robert",
	"content": "Reply to Comment"
}

# Update a post
PUT http://localhost:8080/posts/1
Body:
{
	"author": "Robert",
	"title": "Test",
	"content": "Test Post Update"
}

# Update a comment
PUT http://localhost:8080/comments/1
Body:
{
	"author": "Robert",
	"content": "Test Comment Update"
}
```

**Endpoints w/Auth**
```
# Login a user
POST http://localhost:8080/login
{
	"username": "admin",
	"password": "admin"
}

# Add a new post with authorization (Use the auth token received from the Login)
POST http://localhost:8080/postsAuth
Header: Authorization Bearer <auth_token_from_login>
Body:
{
	"author": "Robert",
	"title": "Test",
	"content": "Auth Testing"
}
```

**Docker Build Step (Optional):**

```
sudo docker build -t blog-api .
```
