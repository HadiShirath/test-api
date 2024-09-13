package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) EditTPSSaksi(ctx context.Context, model User, userId string) (err error) {

	var query string

	if model.Password != "" {
		query = `
		UPDATE auth SET fullname=$1, username=$2, password=$3 WHERE public_id=$4
		`
		_, err = r.db.ExecContext(ctx, query, model.Fullname, model.Username, model.Password, userId)
	} else {
		query = `
			UPDATE auth SET fullname=$1, username=$2 WHERE public_id=$3
			`
		_, err = r.db.ExecContext(ctx, query, model.Fullname, model.Username, userId)
	}

	if err != nil {
		return
	}

	return
}
