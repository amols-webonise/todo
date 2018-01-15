// Package exp contains the types for schema 'public'.
package exp

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/lib/pq"
)

// GooseDbVersion represents a row from 'public.goose_db_version'.
type GooseDbVersion struct {
	ID        int         `json:"id"`         // id
	VersionID int64       `json:"version_id"` // version_id
	IsApplied bool        `json:"is_applied"` // is_applied
	Tstamp    pq.NullTime `json:"tstamp"`     // tstamp

	// xo fields
	_exists, _deleted bool
}

type GooseDbVersionService interface {
	DoesGooseDbVersionExist(gdv *GooseDbVersion) (bool, error)
}

type GooseDbVersionServiceImpl struct {
}

// Exists determines if the GooseDbVersion exists in the database.
func (serviceImpl *GooseDbVersionServiceImpl) Exists(gdv *GooseDbVersion) (bool, error) {
	panic("not yet implemented")
}

// Deleted provides information if the GooseDbVersion has been deleted from the database.
func (gdv *GooseDbVersion) Deleted() bool {
	return gdv._deleted
}

// Insert inserts the GooseDbVersion to the database.
func (gdv *GooseDbVersion) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if gdv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.goose_db_version (` +
		`version_id, is_applied, tstamp` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	err = db.QueryRow(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp).Scan(&gdv.ID)
	if err != nil {
		return err
	}

	// set existence
	gdv._exists = true

	return nil
}

// Update updates the GooseDbVersion in the database.
func (gdv *GooseDbVersion) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !gdv._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if gdv._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.goose_db_version SET (` +
		`version_id, is_applied, tstamp` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp, gdv.ID)
	_, err = db.Exec(sqlstr, gdv.VersionID, gdv.IsApplied, gdv.Tstamp, gdv.ID)
	return err
}

// Save saves the GooseDbVersion to the database.
/*
	func (gdv *GooseDbVersion) Save(db XODB) error {
		if gdv.Exists() {
			return gdv.Update(db)
		}

		return gdv.Insert(db)
	}
*/

// Upsert performs an upsert for GooseDbVersion.
//
// NOTE: PostgreSQL 9.5+ only
func (gdv *GooseDbVersion) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if gdv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.goose_db_version (` +
		`id, version_id, is_applied, tstamp` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, version_id, is_applied, tstamp` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.version_id, EXCLUDED.is_applied, EXCLUDED.tstamp` +
		`)`

	// run query
	XOLog(sqlstr, gdv.ID, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	_, err = db.Exec(sqlstr, gdv.ID, gdv.VersionID, gdv.IsApplied, gdv.Tstamp)
	if err != nil {
		return err
	}

	// set existence
	gdv._exists = true

	return nil
}

// Delete deletes the GooseDbVersion from the database.
func (gdv *GooseDbVersion) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !gdv._exists {
		return nil
	}

	// if deleted, bail
	if gdv._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.goose_db_version WHERE id = $1`

	// run query
	XOLog(sqlstr, gdv.ID)
	_, err = db.Exec(sqlstr, gdv.ID)
	if err != nil {
		return err
	}

	// set deleted
	gdv._deleted = true

	return nil
}

// GetAllGooseDbVersions returns all rows from 'public.goose_db_version',
// ordered by "created_at" in descending order.
func GetAllGooseDbVersions(db XODB) ([]*GooseDbVersion, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.goose_db_version`

	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*GooseDbVersion
	for q.Next() {
		gdv := GooseDbVersion{}

		// scan
		err = q.Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
		if err != nil {
			return nil, err
		}

		res = append(res, &gdv)
	}

	return res, nil
}

// GetChunkedGooseDbVersions returns pagingated rows from 'public.goose_db_version',
// ordered by "created_at" in descending order.
func GetChunkedGooseDbVersions(db XODB, limit int, offset int) ([]*GooseDbVersion, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.goose_db_version LIMIT $1 OFFSET $2`

	q, err := db.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*GooseDbVersion
	for q.Next() {
		gdv := GooseDbVersion{}

		// scan
		err = q.Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
		if err != nil {
			return nil, err
		}

		res = append(res, &gdv)
	}

	return res, nil
}

// GooseDbVersionByID retrieves a row from 'public.goose_db_version' as a GooseDbVersion.
//
// Generated from index 'goose_db_version_pkey'.
func GooseDbVersionByID(db XODB, id int) (*GooseDbVersion, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, version_id, is_applied, tstamp ` +
		`FROM public.goose_db_version ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	gdv := GooseDbVersion{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&gdv.ID, &gdv.VersionID, &gdv.IsApplied, &gdv.Tstamp)
	if err != nil {
		return nil, err
	}

	return &gdv, nil
}
