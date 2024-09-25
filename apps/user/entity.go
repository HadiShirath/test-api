package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Fullname        string `db:"fullname"`
	Username        string `db:"username"`
	Password        string `db:"password"`
	PasswordDecoded string `db:"password_decoded"`

	KecamatanName string `db:"kecamatan_name"`
	KelurahanName string `db:"kelurahan_name"`
	TpsName       string `db:"tps_name"`
}

func NewFromEditTPSSaksiRequest(req EditUserRequestPayload) User {
	return User{
		Fullname:        req.Fullname,
		Username:        req.Username,
		Password:        req.Password,
		PasswordDecoded: req.Password,
	}
}

func (u User) ToUserListResponse() UserListResponse {
	return UserListResponse{
		Username:      u.Username,
		Fullname:      u.Fullname,
		KecamatanName: u.KecamatanName,
		KelurahanName: u.KelurahanName,
		TpsName:       u.TpsName,
	}
}

func (u *User) EncryptPassword(salt int) (err error) {
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	u.Password = string(encryptPass)

	return nil
}
