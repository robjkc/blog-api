# blog-api

**Steps to run**

1. Start a preloaded postgres database using docker-compose (from the root directory). The preloaded data is in the init.sql file. You will need to install docker-compose.
```
docker-compose up
```

2. Start the blog-api from the command line.
```
go run .
```

**Steps to run with both the database and api using docker-compose (Optional)**

1. Update db_connection.go by commenting out line 23 and adding a comment to line 24. This will change the host to "postgres" from "localhost".
```

	// Use the following line if you want to run the api with docker-compose. It sets the host to postgres instead of localhost.
	//db, err := sqlx.Connect("postgres", "user=postgres password=postgres host=postgres dbname=postgres sslmode=disable")
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres host=localhost dbname=postgres sslmode=disable")
```

2. Build the blog-api docker image.
```
sudo docker build -t blog-api .
```

3. Remove the comments from the docker-compose.yml (Lines starting with #).

4. Run the docker-compose file. This will start both the database and the api.
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

**Database tables**
```
create table if not exists posts (
	id serial primary key,
	title varchar(100) not null,
	author varchar(100) not null,
	content varchar(1000) not null,
    create_date timestamp default now()
);

create table if not exists comments (
	id serial primary key,
	post_id int references posts(id) on delete cascade,
	top_level int default 1 not null,
	author varchar(100) not null,
	content varchar(1000) not null,
    create_date timestamp default now()
);

create table if not exists child_comments (
	parent_comment_id int references comments(id) on delete cascade,
	child_comment_id int references comments(id) on delete cascade
);
```
