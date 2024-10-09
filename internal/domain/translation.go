package domain

// TranslationRequest represents a request to translate a given prompt.
type TranslationRequest struct {
	Prompt string `json:"prompt" bson:"prompt"`
}

// Translation represents a translation in the system. It includes the prompt and the
// corresponding endpoint to hit (which is the result of the translation).
type Translation struct {
	ID       string   `json:"id" bson:"_id"`
	Prompt   string   `json:"prompt" bson:"prompt"`
	Endpoint Endpoint `json:"endpoint" bson:"endpoint"`

	// Todo: remove
	Completion interface{} `json:"completion" bson:"completion"`
}
