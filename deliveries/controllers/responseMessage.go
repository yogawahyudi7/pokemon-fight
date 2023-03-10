package controllers

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func (r Response) Saved(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "Data Berhasil Disimpan.",
		Data:    data,
	}
}

func (r Response) Found(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "Data Ditemukan.",
		Data:    data,
	}
}

func (r Response) BadRequest(message string) Response {
	return Response{
		Code:    400,
		Message: message,
		Data:    nil,
	}
}

func (r Response) NotFound(message string) Response {

	response := Response{
		Code:    404,
		Message: message,
		Data:    nil,
	}
	if message == "" {
		response.Message = "Data Tidak Ditemukan."
	}

	return response

}

func (r Response) InternalServerError(err string) Response {
	return Response{
		Code:    500,
		Message: "Maaf, Server Sedang Dalam Perbaikan Cobalah Beberapa Saat Lagi.",
		Data:    nil,
		Error:   err,
	}
}
