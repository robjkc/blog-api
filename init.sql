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

-- hold previous comment id, used while populating the database.
create table if not exists parent_comment (
    id int not null,
	comment_id int not null
);

-- Add post
insert into posts(title, author, content) values ('Hello', 'Robert', 'Just wanted to say hello');

-- Add comment - 1 (parent 1)
insert into comments(author, content, post_id) values('Scott', 'Nice Post', currval('posts_id_seq'));
insert into parent_comment(id, comment_id) values(1, currval('comments_id_seq'));
insert into child_comments(parent_comment_id, child_comment_id) values(currval('comments_id_seq'), currval('comments_id_seq'));

-- Add comment - 2 (parent 1)
insert into comments(author, content, post_id) values('TW', 'I agree', currval('posts_id_seq'));
insert into child_comments(parent_comment_id, child_comment_id) select comment_id, currval('comments_id_seq') from parent_comment where id = 1;
update parent_comment set comment_id = currval('comments_id_seq') where id = 1;

-- Add comment - 3 (parent 2)
insert into comments(author, content, post_id, top_level) values('Robert', 'Thanks', currval('posts_id_seq'), 0);
insert into child_comments(parent_comment_id, child_comment_id) select comment_id, currval('comments_id_seq') from parent_comment where id = 1;
update parent_comment set comment_id = currval('comments_id_seq') where id = 1;

-- Add comment - 4 (parent 3)
insert into comments(author, content, post_id, top_level) values('Scott', 'You are welcome', currval('posts_id_seq'), 0);
insert into child_comments(parent_comment_id, child_comment_id) select comment_id, currval('comments_id_seq') from parent_comment where id = 1;
--update parent_comment set comment_id = currval('comments_id_seq') where id = 1;

-- Add comment - 5 (parent 3)
insert into comments(author, content, post_id, top_level) values('TW', 'Same here', currval('posts_id_seq'), 0);
insert into child_comments(parent_comment_id, child_comment_id) select comment_id, currval('comments_id_seq') from parent_comment where id = 1;
update parent_comment set comment_id = currval('comments_id_seq') where id = 1;


-- Add post
insert into posts(title, author, content) values ('New NBA Team', 'Scott', 'We just signed another NBA team!');

-- Add comment - 6 (parent 6)
insert into comments(author, content, post_id) values('Robert', 'Great News!', currval('posts_id_seq'));
update parent_comment set comment_id = currval('comments_id_seq') where id = 1;
insert into child_comments(parent_comment_id, child_comment_id) select comment_id, currval('comments_id_seq') from parent_comment where id = 1;

-- Add comment - 7 (parent 7)
insert into comments(author, content, post_id) values('TW', 'Thats awesome!', currval('posts_id_seq'));
update parent_comment set comment_id = currval('comments_id_seq') where id = 1;
insert into child_comments(parent_comment_id, child_comment_id) select comment_id, currval('comments_id_seq') from parent_comment where id = 1;

-- Add comment - 8 (parent 7)
insert into comments(author, content, post_id, top_level) values('Scott', 'We need to sign some more', currval('posts_id_seq'), 0);
insert into child_comments(parent_comment_id, child_comment_id) select comment_id, currval('comments_id_seq') from parent_comment where id = 1;
update parent_comment set comment_id = currval('comments_id_seq') where id = 1;
