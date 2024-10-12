package domain

// TranslationRequest represents a request to translate a given prompt.
type TranslationRequest struct {
	Endpoint Endpoint `json:"endpoint" bson:"endpoint"`
	Prompt   string   `json:"prompt" bson:"prompt"`
}

// Translation represents a translation in the system. It includes the prompt, the request body definition and the translation itself.
type Translation struct {
	ID       string   `json:"id" bson:"_id"`
	Prompt   string   `json:"prompt" bson:"prompt"`
	Endpoint Endpoint `json:"endpoint" bson:"endpoint"`

	// Todo: remove
	Translation interface{} `json:"translation" bson:"translation"`
}
