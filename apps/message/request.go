package message

type CreateMessageRequestPayload struct {
	Id             string `json:"id"`
	ReceiverNumber string `json:"receiver_number"`
	Message        string `json:"message"`
	Processed      bool   `json:"processed"`
	// UpdatedAt    time.Time `json:"updated_at"`
	// CreatedAt    string    `json:"created_at"`
}

type StatusMessageRequestPayload struct {
	Processed string `query:"processed" json:"processed"`
}

type UploadInboxRequestPayload struct {
	Id           string `json:"id"`
	SenderNumber string `json:"sender_number"`
	Message      string `json:"message"`
	// UpdatedAt    time.Time `json:"updated_at"`
	// CreatedAt    string    `json:"created_at"`
}
