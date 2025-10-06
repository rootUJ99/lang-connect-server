package language

import "github.com/jackc/pgx/v5/pgtype"

type LangCodeJSON struct {
	Name        string `json:"name"`
	Code        string `json:"code"`
	NativeName  string `json:"native_name"`
	Description pgtype.Text `json:"description"`
	Script pgtype.Text `json:"script"`
	Metadata    map[string]string `json:"metadata"`
}
type LangCodeJSONWithEmpty struct {
	Name        string            `json:"name,omitempty"`
	Code        string            `json:"code,omitempty"`
	Description pgtype.Text `json:"description,omitempty"`
	Script pgtype.Text `json:"script,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type GenericJSONMesage struct {
	Message string `json:"message"`
}
