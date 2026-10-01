package response

type APIResponse struct {
	Data         interface{} `json:"data"`
	Meta         interface{} `json:"meta"`
	Error        string      `json:"error"`
	Success      bool        `json:"success"`
	ResponseCode string      `json:"response_code"`
	Message      string      `json:"message"`
}
