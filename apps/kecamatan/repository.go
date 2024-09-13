package kecamatan

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

func (r repository) GetVoterKecamatan(ctx context.Context, codeKecamatan string) (kecamatan Kecamatan, err error) {
	query := `
		SELECT
    k.kecamatan_name,
    SUM(t.paslon1) AS paslon1,
    SUM(t.paslon2) AS paslon2,
    SUM(t.paslon3) AS paslon3,
    SUM(t.paslon4) AS paslon4,
    SUM(t.suara_tidak_sah) AS suara_tidak_sah
		FROM
			tps t
		JOIN
			kecamatan k ON t.kecamatan_id = k.kecamatan_id
		WHERE
			k.code = $1
		GROUP BY
			k.kecamatan_name;
	`
	err = r.db.GetContext(ctx, &kecamatan, query, codeKecamatan)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetAllVoter(ctx context.Context) (kecamatan Kecamatan, err error) {
	query := `
		SELECT
    SUM(suara_sah + suara_tidak_Sah) AS total_suara,
    ROUND(
        (COUNT(CASE WHEN suara_sah IS NOT NULL AND suara_sah > 0 THEN 1 END) * 100.0 / COUNT(*))
    ) AS persentase
		FROM tps;
	`
	err = r.db.GetContext(ctx, &kecamatan, query)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) GetListKecamatan(ctx context.Context) (kecamatans []Kecamatan, err error) {
	query := `
		SELECT
			c.kecamatan_name,
			COALESCE(SUM(t.paslon1), 0) AS paslon1,
			COALESCE(SUM(t.paslon2), 0) AS paslon2,
			COALESCE(SUM(t.paslon3), 0) AS paslon3,
			COALESCE(SUM(t.paslon4), 0) AS paslon4,
			COALESCE(SUM(t.suara_sah), 0) AS suara_sah,
			COALESCE(SUM(t.suara_tidak_sah), 0) AS suara_tidak_sah,
			c.total_voters,
			c.total_tps,
			COUNT(CASE WHEN t.suara_sah > 0 THEN 1 END) AS sudah,
			COUNT(CASE WHEN t.suara_sah = 0 THEN 1 END) AS belum,
			ROUND(
				(COALESCE(SUM(t.suara_sah), 0) + COALESCE(SUM(t.suara_tidak_sah), 0)) * 100.0 / COALESCE(c.total_voters, 1)
			) AS pp,
			c.code
		FROM
			kecamatan c
		LEFT JOIN
			kelurahan k ON c.kecamatan_id = k.kecamatan_id
		LEFT JOIN
			tps t ON k.kelurahan_id = t.kelurahan_id
		GROUP BY
			c.kecamatan_id, c.kecamatan_name, c.total_voters, c.total_tps
		ORDER BY
			c.kecamatan_name;
	`

	err = r.db.SelectContext(ctx, &kecamatans, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}

func (r repository) GetListKelurahanFromKecamatan(ctx context.Context, codeKecamatan string) (kelurahans []Kelurahan, err error) {
	query := `
	    SELECT
			c.kecamatan_name,
			k.kelurahan_name,
			COALESCE(SUM(t.paslon1), 0) AS paslon1,
			COALESCE(SUM(t.paslon2), 0) AS paslon2,
			COALESCE(SUM(t.paslon3), 0) AS paslon3,
			COALESCE(SUM(t.paslon4), 0) AS paslon4,
			COALESCE(SUM(t.suara_sah), 0) AS suara_sah,
			COALESCE(SUM(t.suara_tidak_sah), 0) AS suara_tidak_sah,
			COALESCE(k.total_voters, 0) AS total_voters,
			COALESCE(k.total_tps, 0) AS total_tps,
			COUNT(CASE WHEN t.suara_sah > 0 THEN 1 END) AS sudah,
			COUNT(CASE WHEN t.suara_sah = 0 THEN 1 END) AS belum,
			ROUND(
				(COALESCE(SUM(t.suara_sah), 0) + COALESCE(SUM(t.suara_tidak_sah), 0)) * 100.0 / COALESCE(NULLIF(k.total_voters, 0), 1)
			) AS pp,
			k.code
		FROM
			kelurahan k
		LEFT JOIN
			tps t ON k.kelurahan_id = t.kelurahan_id
		LEFT JOIN
			kecamatan c ON k.kecamatan_id = c.kecamatan_id
		WHERE
			c.code = $1
		GROUP BY
			k.kelurahan_id, k.kelurahan_name, c.kecamatan_name, k.total_voters, k.total_tps, k.code
		ORDER BY
			k.kelurahan_name;
	`

	err = r.db.SelectContext(ctx, &kelurahans, query, codeKecamatan)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.ErrNotFound
		}
		return

	}
	return

}
