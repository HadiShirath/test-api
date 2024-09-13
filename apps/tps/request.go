package tps

type CreatePhotoRequestPayload struct {
	Photo string `json:"photo"`
}

type EditTPSSaksiRequestPayload struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	UserId   string `json:"user_id"`
}

type EditVoteTPSRequestPayload struct {
	Paslon1       int32 `json:"paslon1"`
	Paslon2       int32 `json:"paslon2"`
	Paslon3       int32 `json:"paslon3"`
	Paslon4       int32 `json:"paslon4"`
	SuaraSah      int32 `json:"suara_sah"`
	SuaraTidakSah int32 `json:"suara_tidak_sah"`
}

type ListTPSSaksiRequestPayload struct {
	Offset int `query:"offset" json:"offset"`
	Limit  int `query:"limit" json:"limit"`
}

func (l ListTPSSaksiRequestPayload) GenerateDefaultValue() ListTPSSaksiRequestPayload {
	if l.Offset < 0 {
		l.Offset = 0
	}
	if l.Limit <= 0 {
		l.Limit = 10
	}
	return l
}