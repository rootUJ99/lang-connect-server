package words 

type LangCodeJSON struct {
	Word string `json:"word"`
	LangCode string `json:"lang_code"` 
}
type LangCodeJSONWithEmpty struct {
	Word string `json:"word,omitempty"`
	LangCode string `json:"lang_code,omitempty"` 
}

type GenericJSONMesage struct {
	Message string `json:"message"`
}


