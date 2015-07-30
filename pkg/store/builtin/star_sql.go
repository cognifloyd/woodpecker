package builtin

// DO NOT EDIT
// code generated by go:generate

import (
	"database/sql"
	"encoding/json"
)

var _ = json.Marshal

// generic database interface, matching both *sql.Db and *sql.Tx
type starDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func getStar(db starDB, query string, args ...interface{}) (*Star, error) {
	row := db.QueryRow(query, args...)
	return scanStar(row)
}

func getStars(db starDB, query string, args ...interface{}) ([]*Star, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanStars(rows)
}

func createStar(db starDB, query string, v *Star) error {
	var v0 int64
	var v1 int64
	v0 = v.UserID
	v1 = v.RepoID

	res, err := db.Exec(query,
		&v0,
		&v1,
	)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func updateStar(db starDB, query string, v *Star) error {
	var v0 int64
	var v1 int64
	var v2 int64
	v0 = v.ID
	v1 = v.UserID
	v2 = v.RepoID

	_, err := db.Exec(query,
		&v1,
		&v2,
		&v0,
	)
	return err
}

func scanStar(row *sql.Row) (*Star, error) {
	var v0 int64
	var v1 int64
	var v2 int64

	err := row.Scan(
		&v0,
		&v1,
		&v2,
	)
	if err != nil {
		return nil, err
	}

	v := &Star{}
	v.ID = v0
	v.UserID = v1
	v.RepoID = v2

	return v, nil
}

func scanStars(rows *sql.Rows) ([]*Star, error) {
	var err error
	var vv []*Star
	for rows.Next() {
		var v0 int64
		var v1 int64
		var v2 int64
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
		)
		if err != nil {
			return vv, err
		}

		v := &Star{}
		v.ID = v0
		v.UserID = v1
		v.RepoID = v2
		vv = append(vv, v)
	}
	return vv, rows.Err()
}

const stmtStarSelectList = `
SELECT
 star_id
,star_user_id
,star_repo_id
FROM stars
`

const stmtStarSelectRange = `
SELECT
 star_id
,star_user_id
,star_repo_id
FROM stars
LIMIT ? OFFSET ?
`

const stmtStarSelect = `
SELECT
 star_id
,star_user_id
,star_repo_id
FROM stars
WHERE star_id = ?
`

const stmtStarSelectStarUserRepo = `
SELECT
 star_id
,star_user_id
,star_repo_id
FROM stars
WHERE star_user_id = ?
AND star_repo_id = ?
`

const stmtStarSelectCount = `
SELECT count(1)
FROM stars
`

const stmtStarInsert = `
INSERT INTO stars (
 star_user_id
,star_repo_id
) VALUES (?,?);
`

const stmtStarUpdate = `
UPDATE stars SET
 star_user_id = ?
,star_repo_id = ?
WHERE star_id = ?
`

const stmtStarDelete = `
DELETE FROM stars
WHERE star_id = ?
`

const stmtStarTable = `
CREATE TABLE IF NOT EXISTS stars (
 star_id	INTEGER PRIMARY KEY AUTOINCREMENT
,star_user_id	INTEGER
,star_repo_id	INTEGER
);
`

const stmtStarStarUserRepoIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS ux_star_user_repo ON stars (star_user_id,star_repo_id);
`