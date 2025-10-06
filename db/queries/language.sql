-- name: ListLanguage :many
SELECT * FROM language; 


-- name: CreateLanguage :one
INSERT INTO language(
	code, name, native_name, script, description, metadata
) VALUES (
	$1, $2, $3, $4, $5, $6
)RETURNING *;

-- name: GetLanguageById :one
SELECT * FROM language 
	WHERE language_id = $1;

-- name: UpdateLanguageById :one
UPDATE language SET 
	code = $2, 
	name = $3,
	native_name = $4, 
	script = $5,
	description = $6,
	metadata = $7
	WHERE language_id = $1
 RETURNING *;

-- name: DeleteLanguageById :one
DELETE FROM language 
	WHERE language_id = $1
RETURNING *;


