package mermaid

import (
	"fmt"
	"os"
	"strings"
)

func CreateFile() error {
	f, err := os.Create("mermaid.md")
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

func CreateTableInfo(tableName string, tableInfo map[string]string) error {
	tmpl := "\n" + tableName + " {\n"
	for k, v := range tableInfo {
		if strings.Contains(v, "int") {
			tmpl += fmt.Sprintf("\t%v %v\n", "int", k)
		} else if strings.Contains(v, "float") {
			tmpl += fmt.Sprintf("\t%v %v\n", "float", k)
		} else if strings.Contains(v, "char") {
			tmpl += fmt.Sprintf("\t%v %v\n", "string", k)
		} else {
			tmpl += fmt.Sprintf("\t%v %v\n", v, k)
		}
	}
	tmpl += "}\n\n"

	f, err := os.OpenFile("mermaid.md", os.O_WRONLY|os.O_APPEND, 0666)
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

func CreateReferenceInfo(tableName, referenceTableName string) error {
	tmpl := tableName + " ||--|{ " + referenceTableName + ": \"\"" + "\n"

	f, err := os.OpenFile("mermaid.md", os.O_WRONLY|os.O_APPEND, 0666)
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
