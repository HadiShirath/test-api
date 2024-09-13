package user

import (
	"nbid-online-shop/infra/response"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Fullname string `db:"fullname"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func NewFromEditTPSSaksiRequest(req EditUserRequestPayload) User {
	return User{
		Fullname: req.Fullname,
		Username: req.Username,
		Password: req.Password,
	}
}

func (u User) ValidatePassword() (err error) {
	if u.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(u.Password) < 8 {
		return response.ErrPasswordInvalidLength
	}

	return
}

func (u *User) EncryptPassword(salt int) (err error) {
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	u.Password = string(encryptPass)

	return nil
}

func (u User) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

func (u User) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(u.Password))
}
