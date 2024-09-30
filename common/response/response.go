package response

type TestMessageResponse struct {
	Body struct {
		Message1 string `json:"message_1"`
		Message2 string `json:"message_2"`
		Message3 string `json:"message_3"`
	}
}
