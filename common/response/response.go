package response

type TestMessageResponse struct {
	Body struct {
		Message1 string `json:"message_1"`
		Message2 string `json:"message_2"`
		Message3 string `json:"message_3"`
	}
}

type GenericResponse[args any] struct {
	Body BodyResponse[args]
}

type BodyResponse[args any] struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    args   `json:"data"`
}
