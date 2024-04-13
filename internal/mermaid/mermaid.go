package mermaid

import (
	"fmt"
	"os"
	"strings"

	"github.com/ritarock/gerd/internal/model"
)

const FILE_NAME = "mermaid.md"

func Create() error {
	f, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte("erDiagram\n"))
	if err != nil {
		return err
	}

	return nil
}

func Delete() error {
	return os.Remove(FILE_NAME)
}

func CreateTable(tableName model.Table, column model.Column) error {
	tmpl := "\n" + tableName.String() + " {\n"
	for k, v := range column {
		if strings.Contains(v.String(), "int") {
			tmpl += fmt.Sprintf("\t%v %v\n", "int", k)
		} else if strings.Contains(v.String(), "float") {
			tmpl += fmt.Sprintf("\t%v %v\n", "float", k)
		} else if strings.Contains(v.String(), "char") {
			tmpl += fmt.Sprintf("\t%v %v\n", "string", k)
		} else {
			tmpl += fmt.Sprintf("\t%v %v\n", v, k)
		}
	}
	tmpl += "}\n\n"

	f, err := os.OpenFile(FILE_NAME, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(tmpl))
	if err != nil {
		return err
	}

	return nil
}

func CreateReference(tableName model.Table, reference model.MetaReference) error {
	tmpl := tableName.String() + " ||--|{ " + reference.TableName + ": \"\"" + "\n"

	f, err := os.OpenFile(FILE_NAME, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(tmpl))
	if err != nil {
		return err
	}

	return nil
}
