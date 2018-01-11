// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
)

// PortalRole represents a row from 'public.portal_role'.
type PortalRole struct {
	ID       int    `json:"id"`        // id
	RoleName string `json:"role_name"` // role_name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PortalRole exists in the database.
func (pr *PortalRole) Exists() bool {
	return pr._exists
}

// Deleted provides information if the PortalRole has been deleted from the database.
func (pr *PortalRole) Deleted() bool {
	return pr._deleted
}

// Insert inserts the PortalRole to the database.
func (pr *PortalRole) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.portal_role (` +
		`role_name` +
		`) VALUES (` +
		`$1` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, pr.RoleName)
	err = db.QueryRow(sqlstr, pr.RoleName).Scan(&pr.ID)
	if err != nil {
		return err
	}

	// set existence
	pr._exists = true

	return nil
}

// Update updates the PortalRole in the database.
func (pr *PortalRole) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.portal_role SET (` +
		`role_name` +
		`) = ( ` +
		`$1` +
		`) WHERE id = $2`

	// run query
	XOLog(sqlstr, pr.RoleName, pr.ID)
	_, err = db.Exec(sqlstr, pr.RoleName, pr.ID)
	return err
}

// Save saves the PortalRole to the database.
func (pr *PortalRole) Save(db XODB) error {
	if pr.Exists() {
		return pr.Update(db)
	}

	return pr.Insert(db)
}

// Upsert performs an upsert for PortalRole.
//
// NOTE: PostgreSQL 9.5+ only
func (pr *PortalRole) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if pr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.portal_role (` +
		`id, role_name` +
		`) VALUES (` +
		`$1, $2` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, role_name` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.role_name` +
		`)`

	// run query
	XOLog(sqlstr, pr.ID, pr.RoleName)
	_, err = db.Exec(sqlstr, pr.ID, pr.RoleName)
	if err != nil {
		return err
	}

	// set existence
	pr._exists = true

	return nil
}

// Delete deletes the PortalRole from the database.
func (pr *PortalRole) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pr._exists {
		return nil
	}

	// if deleted, bail
	if pr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.portal_role WHERE id = $1`

	// run query
	XOLog(sqlstr, pr.ID)
	_, err = db.Exec(sqlstr, pr.ID)
	if err != nil {
		return err
	}

	// set deleted
	pr._deleted = true

	return nil
}

// GetAllPortalRoles returns all rows from 'public.portal_role',
// ordered by "created_at" in descending order.
func GetAllPortalRoles(db XODB) ([]*PortalRole, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.portal_role`

	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*PortalRole
	for q.Next() {
		pr := PortalRole{}

		// scan
		err = q.Scan(&pr.ID, &pr.RoleName)
		if err != nil {
			return nil, err
		}

		res = append(res, &pr)
	}

	return res, nil
}

// GetChunkedPortalRoles returns pagingated rows from 'public.portal_role',
// ordered by "created_at" in descending order.
func GetChunkedPortalRoles(db XODB, limit int, offset int) ([]*PortalRole, error) {
	const sqlstr = `SELECT ` +
		`*` +
		`FROM public.portal_role LIMIT $1 OFFSET $2`

	q, err := db.Query(sqlstr, limit, offset)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*PortalRole
	for q.Next() {
		pr := PortalRole{}

		// scan
		err = q.Scan(&pr.ID, &pr.RoleName)
		if err != nil {
			return nil, err
		}

		res = append(res, &pr)
	}

	return res, nil
}

// PortalRoleByID retrieves a row from 'public.portal_role' as a PortalRole.
//
// Generated from index 'portal_role_pkey'.
func PortalRoleByID(db XODB, id int) (*PortalRole, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, role_name ` +
		`FROM public.portal_role ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	pr := PortalRole{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&pr.ID, &pr.RoleName)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

// PortalRoleByRoleName retrieves a row from 'public.portal_role' as a PortalRole.
//
// Generated from index 'portal_role_role_name_key'.
func PortalRoleByRoleName(db XODB, roleName string) (*PortalRole, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, role_name ` +
		`FROM public.portal_role ` +
		`WHERE role_name = $1`

	// run query
	XOLog(sqlstr, roleName)
	pr := PortalRole{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, roleName).Scan(&pr.ID, &pr.RoleName)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}