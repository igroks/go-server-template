package api

type MatchRequest struct {
	Query *string `json:"text" binding:"required"`
}

type MatchResponse struct {
	Message string `json:"msg"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}
