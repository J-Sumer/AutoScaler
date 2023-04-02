package contracts

type BaseResponse struct {
	Error   *string `json:"error"` 
	Message *string `json:"message"`
}
