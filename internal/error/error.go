package custom_errors

type APIError struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}

func (e *APIError) Error() string {
	return e.Detail
}

func NewAPIError(status int, problemType, title, detail string) *APIError {
	return &APIError{
		Status: status,
		Type:   problemType,
		Title:  title,
		Detail: detail,
	}
}
