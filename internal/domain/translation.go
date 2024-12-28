package domain

import "fmt"

const (
	TranslationTypeObject = "object"
)

// TranslationRequest represents a request to translate a given prompt.
type TranslationRequest struct {
	CurrentDate     string            `json:"current_date" bson:"current_date"`
	InputText       string            `json:"input_text" bson:"input_text"`
	Lang            string            `json:"lang" bson:"lang"`
	TranslationType string            `json:"translation_type" bson:"translation_type"`
	TargetObject    *ObjectDefinition `json:"target_object" bson:"target_object"`
}

func (tr *TranslationRequest) Validate() error {
	if tr.CurrentDate == "" {
		return fmt.Errorf("%w: current date is required", ErrorValidation)
	}
	if tr.InputText == "" {
		return fmt.Errorf("%w: input text is required", ErrorValidation)
	}
	switch tr.TranslationType {
	case "":
		return fmt.Errorf("%w: translate to is required", ErrorValidation)
	case TranslationTypeObject:
		if tr.TargetObject == nil {
			return fmt.Errorf("%w: target object is required for translation type '%s'", ErrorValidation, TranslationTypeObject)
		}
		if err := tr.TargetObject.Validate(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%w: translation type '%s' is not supported", ErrorValidation, tr.TranslationType)
	}

	return nil
}

// Translation represents a translation in the system. It includes the prompt, the request body definition and the translation itself.
type Translation struct {
	ID                 string             `json:"id" bson:"_id"`
	TranslationRequest TranslationRequest `json:"translation_request" bson:"translation_request"`
	MappedObject       interface{}        `json:"mapped_object" bson:"mapped_object"`
}
