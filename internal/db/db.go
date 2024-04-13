package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ritarock/gerd/internal/model"
)

type dbx struct {
	db *sqlx.DB
}

func NewDbx() *dbx {
	return &dbx{
		db: nil,
	}
}

func (d *dbx) Connect(db, address, user, password string) error {
	config := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 address,
		DBName:               db,
		AllowNativePasswords: false,
	}

	c, err := sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		return err
	}
	d.db = c
	return nil
}

func (d *dbx) GetTables() ([]model.Table, error) {
	var tables []model.Table
	if err := d.db.Select(&tables, "SHOW TABLES"); err != nil {
		return nil, err
	}

	return tables, nil
}

func (d *dbx) GetDescribe(tableName model.Table) (model.Column, error) {
	query := "SHOW COLUMNS FROM " + tableName.String()
	metaColumns := []model.MetaColumn{}
	if err := d.db.Select(&metaColumns, query); err != nil {
		return nil, err
	}

	column := make(model.Column)
	for _, v := range metaColumns {
		key := model.ColumnKey(v.Field.String)
		value := model.ColumnValue(v.Field.String)
		column[key] = value
	}
	return column, nil
}

func (d *dbx) GetReferences(table model.Table) ([]model.MetaReference, error) {
	query := `
SELECT
	kcu.TABLE_NAME, tc.CONSTRAINT_TYPE
FROM
	information_schema.KEY_COLUMN_USAGE kcu
INNER JOIN
	information_schema.TABLE_CONSTRAINTS tc
ON
	kcu.CONSTRAINT_NAME = tc.CONSTRAINT_NAME
WHERE
	CONSTRAINT_TYPE = 'FOREIGN KEY'
AND
	REFERENCED_TABLE_NAME = '` + table.String() + "'"

	references := []model.MetaReference{}
	if err := d.db.Select(&references, query); err != nil {
		return nil, err
	}

	return references, nil
}
