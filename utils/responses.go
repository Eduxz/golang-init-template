package utils

//CustomResponse for every Endpoint
type CustomResponse struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

//ErrorResponse for every Endpoint
type ErrorResponse struct {
	Err    interface{} `json:"error"`
	Status string      `json:"status"`
}

// ResponseSuccess for every endpoint response
func ResponseSuccess(data interface{}, status string) CustomResponse {
	return CustomResponse{data, status}
}

// ResponseError for every endpoint response
func ResponseError(err interface{}, status string) ErrorResponse {
	return ErrorResponse{err, status}
}
