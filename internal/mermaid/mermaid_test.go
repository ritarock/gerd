package mermaid

import (
	"io"
	"os"
	"testing"

	"github.com/ritarock/gerd/internal/model"
	"github.com/stretchr/testify/assert"
)

func assertReadData(t *testing.T, file *os.File, expected string) {
	t.Helper()
	b, _ := io.ReadAll(file)
	assert.Equal(t, expected, string(b))
}

func TestCreate(t *testing.T) {
	tmpFile, err := os.CreateTemp("./", "testfile")
	if err != nil {
		t.Fatalf("failed TestWrite")
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	err = Create(tmpFile.Name())
	assert.NoError(t, err)

	assertReadData(t, tmpFile, "erDiagram\n")
}

func TestCreateTable(t *testing.T) {
	type args struct {
		tableName model.Table
		column    model.Column
	}

	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "pass",
			args: args{
				tableName: "samples",
				column: map[model.ColumnKey]model.ColumnValue{
					model.ColumnKey("id"):   model.ColumnValue("int"),
					model.ColumnKey("name"): model.ColumnValue("varchar(45)"),
				},
			},
			expected: "\nsamples {\n\tint id\n\tvarchar(45) name\n}\n\n",
		},
	}

	for _, test := range tests {
		tmpFile, err := os.CreateTemp("./", "testfile")
		if err != nil {
			t.Fatalf("failed TestWrite")
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		err = CreateTable(test.args.tableName, test.args.column, tmpFile.Name())
		assert.NoError(t, err)
		assertReadData(t, tmpFile, test.expected)
	}
}

func TestCreateReference(t *testing.T) {
	type args struct {
		tableName model.Table
		reference model.MetaReference
	}

	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "pass",
			args: args{
				tableName: "samples",
				reference: model.MetaReference{
					TableName:      "sample_categories",
					ConstraintType: "FOREIGN KEY",
				},
			},
			expected: "samples ||--|{ sample_categories: \"\"\n",
		},
	}

	for _, test := range tests {
		tmpFile, err := os.CreateTemp("./", "testfile")
		if err != nil {
			t.Fatalf("failed TestWrite")
		}
		defer os.Remove(tmpFile.Name())
		defer tmpFile.Close()

		err = CreateReference(test.args.tableName, test.args.reference, tmpFile.Name())
		assert.NoError(t, err)
		assertReadData(t, tmpFile, test.expected)
	}
}
