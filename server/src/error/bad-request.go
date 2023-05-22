package error

type BadRequest struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
