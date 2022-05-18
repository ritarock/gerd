package db

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ritarock/gerd/internal/db/types"
)

var mysqlConfig = mysql.Config{}

func Connection(user, passwd, addr, dbName string) (*sqlx.DB, error) {
	mysqlConfig = mysql.Config{
		User:                 user,
		Passwd:               passwd,
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               dbName,
		AllowNativePasswords: true,
	}
	client, err := sqlx.Connect("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetShowTables(client *sqlx.DB) ([]string, error) {
	tables := []string{}
	if err := client.Select(&tables, "SHOW TABLES"); err != nil {
		return nil, nil
	}

	return tables, nil
}

func GetDescribeTable(client *sqlx.DB, tableName string) (map[string]string, error) {
	query := "SHOW COLUMNS FROM " + tableName
	columnInfo := []types.ColumnInfo{}
	if err := client.Select(&columnInfo, query); err != nil {
		return nil, err
	}

	m := map[string]string{}
	for _, v := range columnInfo {
		m[v.Field.String] = v.Type.String
	}

	return m, nil
}

func GetReferences(client *sqlx.DB, table string) []types.ReferenceInfo {
	query := fmt.Sprintf(`
SELECT kcu.TABLE_NAME, tc.CONSTRAINT_TYPE
FROM
	information_schema.KEY_COLUMN_USAGE kcu
INNER JOIN information_schema.TABLE_CONSTRAINTS tc
	ON kcu.CONSTRAINT_NAME = tc.CONSTRAINT_NAME
WHERE
	CONSTRAINT_TYPE = 'FOREIGN KEY' AND REFERENCED_TABLE_NAME = '%v'`, table)

	referenceInfo := []types.ReferenceInfo{}
	if err := client.Select(&referenceInfo, query); err != nil {
		fmt.Println(err)
	}

	if len(referenceInfo) == 0 {
		return nil
	}

	tables := []types.ReferenceInfo{}
	for _, reference := range referenceInfo {
		tables = append(tables, reference)
	}

	return tables
}
