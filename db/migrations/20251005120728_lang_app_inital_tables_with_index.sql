-- +goose Up
-- +goose StatementBegin

CREATE TABLE language(
	language_id SERIAL PRIMARY KEY,
	code VARCHAR NOT NULL,
	name VARCHAR NOT NULL UNIQUE,
	native_name VARCHAR NOT NULL,
	script VARCHAR,
	description VARCHAR,
	metadata JSONB,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE concept (
	concept_id BIGSERIAL PRIMARY KEY,
	description VARCHAR NOT NULL,
	source VARCHAR,
	metadata JSONB,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE term (
	term_id BIGSERIAL PRIMARY KEY,
	language_id INT NOT NULL REFERENCES language(language_id),
	base_form_id INT REFERENCES term(term_id),
	term_text VARCHAR NOT NULL,
	romanized_text VARCHAR NOT NULL,
	en_translation VARCHAR,
	part_of_speech VARCHAR,
	gender VARCHAR,
	plurality VARCHAR,
	metadata JSONB,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE sentence (
	sentence_id BIGSERIAL PRIMARY KEY,
	language_id INT NOT NULL REFERENCES language(language_id),
	source VARCHAR,
	sentence_text VARCHAR,
	romanized_sentence VARCHAR,
	en_translation VARCHAR,
	metadata JSONB,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE term_relation_map(
	term_id1 INT NOT NULL REFERENCES term(term_id) ON DELETE CASCADE,
	term_id2 INT NOT NULL REFERENCES term(term_id) ON DELETE CASCADE,
	relation VARCHAR NOT NULL,
	PRIMARY KEY(term_id1, term_id2, relation)
);

CREATE TABLE term_concept_map (
	term_id INT NOT NULL REFERENCES term(term_id) ON DELETE CASCADE,
	concept_id INT NOT NULL REFERENCES concept(concept_id) ON DELETE CASCADE,
	PRIMARY KEY(term_id, concept_id)
);

CREATE TABLE sentence_concept_map (
	sentence_id INT NOT NULL REFERENCES sentence(sentence_id) ON DELETE CASCADE,
	concept_id INT NOT NULL REFERENCES concept(concept_id) ON DELETE CASCADE,
	PRIMARY KEY(sentence_id, concept_id)
);

CREATE TABLE sentence_term_map (
	sentence_id INT NOT NULL REFERENCES sentence(sentence_id) ON DELETE CASCADE,
	term_id INT NOT NULL REFERENCES term(term_id) ON DELETE CASCADE,
	PRIMARY KEY(sentence_id, term_id)
);


CREATE INDEX idx_term_language_id ON term(language_id);
CREATE INDEX idx_term_base_form_id ON term(base_form_id);
CREATE INDEX idx_sentence_language_id ON sentence(language_id);
CREATE INDEX idx_term_term_text ON term(term_text);
CREATE INDEX idx_term_romanized_text ON term(romanized_text);
CREATE INDEX idx_sentence_sentence_text ON sentence(sentence_text);
CREATE INDEX idx_sentence_romanized_text ON sentence(romanized_sentence);
CREATE INDEX idx_term_relation_map_term_id_2 ON term_relation_map(term_id2);
CREATE INDEX idx_term_concept_map_concept_id ON term_concept_map(concept_id);
CREATE INDEX idx_sentence_concept_map_concept_id ON sentence_concept_map(concept_id);
CREATE INDEX idx_sentence_term_map_term_id ON sentence_term_map(term_id);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_term_language_id;
DROP INDEX IF EXISTS idx_term_base_form_id;
DROP INDEX IF EXISTS idx_sentence_language_id;
DROP INDEX IF EXISTS idx_term_term_text;
DROP INDEX IF EXISTS idx_term_romanized_text;
DROP INDEX IF EXISTS idx_sentence_sentence_text;
DROP INDEX IF EXISTS idx_sentence_romanized_text;
DROP INDEX IF EXISTS idx_term_relation_map_term_id_2;
DROP INDEX IF EXISTS idx_term_concept_map_concept_id;
DROP INDEX IF EXISTS idx_sentence_concept_map_concept_id;
DROP INDEX IF EXISTS idx_sentence_term_map_term_id;

DROP TABLE sentence_term_map;
DROP TABLE sentence_concept_map;
DROP TABLE term_concept_map;
DROP TABLE term_relation_map;
DROP TABLE sentence;
DROP TABLE term;
DROP TABLE concept;
DROP TABLE language;

-- +goose StatementEnd


