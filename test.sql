-- Queries
-- The query to retrieve all the child comments for comment #3 will be as follows
SELECT * FROM Comments AS c
JOIN child_comments cc ON c.id = cc.child_comment_id
WHERE
   cc.parent_comment_id = 3;


-- The query to retrieve all the parent comments for comment #6 will be as follows
SELECT c.*
FROM
   Comments AS c
JOIN child_comments cc ON c.id = pc.child_comment_id
WHERE
   cc.child_comment_id = 6;
