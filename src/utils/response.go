package utils

//CustomResponse for every Endpoint
type CustomResponse struct {
	Data    string `json: "data"`
	Success bool   `json: "success"`
}
