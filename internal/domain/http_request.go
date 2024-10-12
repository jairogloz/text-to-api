package domain

type HTTPRequest struct {
	Name        string              `json:"name" bson:"name"`
	Path        string              `json:"path" bson:"path"`
	Method      string              `json:"method" bson:"method"`
	RequestBody interface{}         `json:"request_body" bson:"request_body"`
	QueryParams map[string][]string `json:"query_params" bson:"query_params"`
	URLParams   map[string]string   `json:"url_params" bson:"url_params"`
}
