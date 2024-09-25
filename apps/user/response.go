package user

type UserListResponse struct {
	Username      string `json:"username"`
	Fullname      string `json:"fullname"`
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsName       string `json:"tps_name"`
}

func NewUserListResponseFromEntity(users []User) []UserListResponse {
	var userList = []UserListResponse{}

	for _, inbox := range users {
		userList = append(userList, inbox.ToUserListResponse())
	}

	return userList
}
