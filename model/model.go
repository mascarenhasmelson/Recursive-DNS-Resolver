package model

type QueryResponse struct {
	Records []string `json:"records"`
	Error   string   `json:"error"`
}
