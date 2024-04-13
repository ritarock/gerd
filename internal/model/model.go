package model

import "database/sql"

type Table string

type Column map[ColumnKey]ColumnValue
type ColumnKey string
type ColumnValue string

type MetaColumn struct {
	Field   sql.NullString `db:"Field"`
	Type    sql.NullString `db:"Type"`
	Null    sql.NullString `db:"Null"`
	Key     sql.NullString `db:"Key"`
	Default sql.NullString `db:"Default"`
	Extra   sql.NullString `db:"Extra"`
}

type MetaReference struct {
	TableName      string `db:"TABLE_NAME"`
	ConstraintType string `db:"CONSTRAINT_TYPE"`
}

func (t Table) String() string {
	return string(t)
}

func (v ColumnValue) String() string {
	return string(v)
}
