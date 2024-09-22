package message

import (
	"context"
	"database/sql"
	"nbid-online-shop/infra/response"

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

func (r repository) GetInboxList(ctx context.Context) (inboxs []Inbox, err error) {
	query := `
	SELECT
		id, sender_number, message, created_at, updated_at
	FROM 
		inbox
	ORDER BY
		created_at DESC
	`

	err = r.db.SelectContext(ctx, &inboxs, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}

	return

}

func (r repository) GetOutboxList(ctx context.Context, model StatusMessage) (outboxs []Outbox, err error) {

	var query string

	if model.Processed != "" {
		query = `
		SELECT
			id, receiver_number, message, processed, created_at, updated_at
		FROM 
			outbox
		WHERE
			processed=$1
		ORDER BY
			created_at ASC
		`
		err = r.db.SelectContext(ctx, &outboxs, query, model.Processed)
	} else {
		query = `
			SELECT
				id, receiver_number, message, processed, created_at, updated_at
			FROM 
				outbox
			ORDER BY
				created_at DESC
			`
		err = r.db.SelectContext(ctx, &outboxs, query)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}

	return

}

func (r repository) CreateMessage(ctx context.Context, model Outbox) (err error) {
	query := `
		INSERT INTO outbox (
			id, receiver_number, message, processed
		) VALUES (
			:id, :receiver_number, :message, :processed
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) UploadInbox(ctx context.Context, model Inbox) (err error) {
	query := `
		INSERT INTO inbox (
			id, sender_number, message
		) VALUES (
			:id, :sender_number, :message
		)
	`

	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}
