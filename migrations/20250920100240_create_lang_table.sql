-- +goose Up
-- +goose StatementBegin
CREATE TABLE common_words(
	id BIGSERIAL PRIMARY KEY,
	word VARCHAR NOT NULL,
	lang_code VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP common_words;
-- +goose StatementEnd
