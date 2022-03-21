# blog-api

**Docker Build Steps:**

sudo docker build -t blog-api .

**Docker Compose Steps (from root directory):**

docker-compose up

**Example Endpoints**
GET http://localhost:8080/posts

GET http://localhost:8080/posts/1

POST http://localhost:8080/posts
Body:
{
	"author": "Robert",
	"title": "Test",
	"content": "Testing "
}

POST http://localhost:8080/comments
Body:
{
	"post_id": 1,
	"author": "Robert",
	"content": "Test Comment"
}

POST http://localhost:8080/comments
Body:
{
	"post_id": 1,
	"parent_comment_id": 2,
	"author": "Robert",
	"content": "Reply to Comment"
}

PUT http://localhost:8080/posts/1
Body:
{
	"author": "Robert",
	"title": "Test",
	"content": "Test Post Update"
}

PUT http://localhost:8080/comments/1
Body:
{
	"author": "Robert",
	"content": "Test Comment Update"
}

**Endpoints w/Auth**
POST http://localhost:8080/login
{
	"username": "admin",
	"password": "admin"
}

POST http://localhost:8080/postsAuth
Header: Authorization Bearer <auth_token_from_login>
Body:
{
	"author": "Robert",
	"title": "Test",
	"content": "Auth Testing"
}
