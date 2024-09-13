package auth

type RegisterRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

type LoginRequestPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
