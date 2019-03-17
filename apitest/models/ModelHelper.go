package models

import (
	"database/sql"
	"reflect"
)

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 sql.NullInt64

// Scan implements the Scanner interface for NullInt64
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ni = NullInt64{i.Int64, false}
	} else {
		*ni = NullInt64{i.Int64, true}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (n *NullInt64) Value() int64 {
	if !n.Valid {
		return 0
	}
	return n.Int64
}

type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}
	return nil
}

// Value implements the driver Valuer interface.
func (ns NullString) Value() string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}

type NullBool sql.NullBool

// Scan implements the Scanner interface.
func (n *NullBool) Scan(value interface{}) error {
	if value == nil {
		n.Bool, n.Valid = false, false
		return nil
	}
	n.Valid = true
	return nil
}

// Value implements the driver Valuer interface.
func (n NullBool) Value() bool {
	if !n.Valid {
		return false
	}
	return n.Bool
}
