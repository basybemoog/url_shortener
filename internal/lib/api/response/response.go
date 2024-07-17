package response

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "ok"
	StatusError = "error"
)

func OK(msg string) Response {
	return Response{
		Status: StatusOK,
		Error:  msg,
	}
}

func ERROR(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}

}
