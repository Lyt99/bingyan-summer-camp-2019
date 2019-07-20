package serializer

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}
