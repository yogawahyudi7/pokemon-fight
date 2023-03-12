package common

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

func (r Response) Login(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "Berhasil Login.",
		Data:    data,
	}
}

func (r Response) Logout(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "Berhasil Logout.",
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

func (r Response) Unauthorized(code int) Response {

	switch code {
	case 2:
		return Response{
			Code:    401,
			Message: "Maaf, User Tidak Valid Mohon Untuk Mencoba Login Kembali.",
			Data:    nil,
		}
	case 3:
		return Response{
			Code:    401,
			Message: "Maaf, User Tidak Valid Mohon Login Menggunakan User Lain.",
			Data:    nil,
		}
	case 1:
		return Response{
			Code:    401,
			Message: "Maaf, Dimohon Untuk Menginputan Autorisasi Yang Valid",
			Data:    nil,
		}
	case 4:
		return Response{
			Code:    401,
			Message: "Maaf, Email Atau Password Salah.",
			Data:    nil,
		}
	case 5:
		return Response{
			Code:    401,
			Message: "Maaf, Level Belum Terdaftar Pada Server.",
			Data:    nil,
		}
	}

	return Response{
		Code:    401,
		Message: "Sorry, This Code Unassigned",
		Data:    nil,
	}
}
