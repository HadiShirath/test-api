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

type ExportDataCSVResponse struct {
	KecamatanName string `json:"kecamatan_name"`
	KelurahanName string `json:"kelurahan_name"`
	TpsName       string `json:"tps_name"`
	FullName      string `json:"fullname"`
	Username      string `json:"username"`
	Password      string `json:"password_decoded"`
	CodeUnique    string `json:"code_unique"`
}

func NewExportDataCSVResponseFromEntity(users []User) []ExportDataCSVResponse {
	var exportDataCSVList = []ExportDataCSVResponse{}

	for _, user := range users {
		exportDataCSVList = append(exportDataCSVList, user.ToExportDataCSVResponse())
	}

	return exportDataCSVList
}
