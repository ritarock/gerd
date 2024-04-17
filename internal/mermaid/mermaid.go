package mermaid

import (
	"fmt"
	"os"

	"github.com/ritarock/gerd/internal/model"
)

func Create(fileName string) error {
	f, err := os.Create(fileName)
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

func Delete(fileName string) error {
	return os.Remove(fileName)
}

func CreateTable(tableName model.Table, column model.Column, fileName string) error {
	tmpl := "\n" + tableName.String() + " {\n"
	for k, v := range column {
		tmpl += fmt.Sprintf("\t%v %v\n", v, k)
	}
	tmpl += "}\n\n"

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
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

func CreateReference(tableName model.Table, reference model.MetaReference, fileName string) error {
	tmpl := tableName.String() + " ||--|{ " + reference.TableName + ": \"\"" + "\n"

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0666)
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
