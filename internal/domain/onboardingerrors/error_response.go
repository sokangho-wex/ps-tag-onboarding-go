package onboardingerrors

type errorResponse struct {
	Error   string   `json:"error"`
	Details []string `json:"details,omitempty"`
}
