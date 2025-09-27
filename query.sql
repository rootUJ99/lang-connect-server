-- name: ListCommonWords :many
SELECT * FROM common_words; 


-- name: CreateCommonWord :one
INSERT INTO common_words(
	word, lang_code	
) VALUES (
	$1, $2
)RETURNING *;

-- name: GetCommonWordById :one
SELECT * FROM common_words 
	WHERE id = $1;

-- name: UpdateCommonWordById :one
UPDATE common_words SET word = $2, 
	lang_code =  $3
	WHERE id = $1
 RETURNING *;

-- name: DeleteCommonWordById :one
DELETE FROM common_words 
	WHERE id = $1
RETURNING *;
