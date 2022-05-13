package types

import "database/sql"

type ColumnInfo struct {
	Field   sql.NullString `db:"Field"`
	Type    sql.NullString `db:"Type"`
	Null    sql.NullString `db:"Null"`
	Key     sql.NullString `db:"Key"`
	Default sql.NullString `db:"Default"`
	Extra   sql.NullString `db:"Extra"`
}

type ReferenceInfo struct {
	TableName      string `db:"TABLE_NAME"`
	ConstraintType string `db:"CONSTRAINT_TYPE"`
}
