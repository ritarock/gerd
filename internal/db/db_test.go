package db

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ritarock/gerd/internal/model"
	"github.com/stretchr/testify/assert"
)

func NewMockDB() (*dbx, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	dbMock := sqlx.NewDb(db, "sqlmock")

	return &dbx{db: dbMock}, mock
}

func Test_dbx_GetTables(t *testing.T) {
	tests := []struct {
		name     string
		mockSql  func(mock sqlmock.Sqlmock)
		expected []model.Table
	}{
		{
			name: "pass",
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SHOW TABLES").
					WillReturnRows(sqlmock.NewRows([]string{
						"Tables_in_sample",
					},
					).AddRow("table1").AddRow("table2"))
			},
			expected: []model.Table{
				"table1", "table2",
			},
		},
	}

	for _, test := range tests {
		dbMock, mock := NewMockDB()
		defer dbMock.db.Close()

		test.mockSql(mock)
		tables, err := dbMock.GetTables()
		assert.NoError(t, err)
		assert.Equal(t, test.expected, tables)

	}
}

func Test_dbx_GetDescribe(t *testing.T) {
	tests := []struct {
		name     string
		mockSql  func(mock sqlmock.Sqlmock)
		expected model.Column
	}{
		{
			name: "pass",
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SHOW COLUMNS FROM samples").
					WillReturnRows(sqlmock.NewRows([]string{
						"Field",
						"Type",
						"Null",
						"Key",
						"Default",
						"Extra",
					}).
						AddRow("ID", "int(11)", "NO", "PRI", "NULL", "auto_increment").
						AddRow("Name", "char(35)", "NO", "", "", ""),
					)
			},
			expected: model.Column{
				model.ColumnKey("ID"):   model.ColumnValue("int(11)"),
				model.ColumnKey("Name"): model.ColumnValue("char(35)"),
			},
		},
	}

	for _, test := range tests {
		dbMock, mock := NewMockDB()
		defer dbMock.db.Close()

		test.mockSql(mock)
		column, err := dbMock.GetDescribe(model.Table("samples"))
		assert.NoError(t, err)
		assert.Equal(t, column, test.expected)
	}
}

func Test_dbx_GetReferences(t *testing.T) {
	tests := []struct {
		name     string
		mockSql  func(mock sqlmock.Sqlmock)
		expected []model.MetaReference
	}{
		{
			name: "pass",
			mockSql: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`
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
						REFERENCED_TABLE_NAME = 'samples'`,
				)).
					WillReturnRows(sqlmock.NewRows([]string{
						"TABLE_NAME",
						"CONSTRAINT_TYPE",
					}).
						AddRow("sample_categories", "FOREIGN KEY"),
					)
			},
			expected: []model.MetaReference{
				{
					TableName:      "sample_categories",
					ConstraintType: "FOREIGN KEY",
				},
			},
		},
	}

	for _, test := range tests {
		dbMock, mock := NewMockDB()
		defer dbMock.db.Close()

		test.mockSql(mock)
		references, err := dbMock.GetReferences(model.Table("samples"))
		assert.NoError(t, err)
		assert.Equal(t, test.expected, references)
	}
}
