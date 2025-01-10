package helpers

type MapResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WebResponse(code int, message string, data interface{}) MapResponse {
	return MapResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type PaginationMapResponse struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	NextPage bool        `json:"next_page"`
	Data     interface{} `json:"data, omitempty"`
}

func FindAllWebResponse(code int, message string, data interface{}, nextpage bool) PaginationMapResponse {
	return PaginationMapResponse{
		Code:     code,
		Message:  message,
		Data:     data,
		NextPage: nextpage,
	}
}
