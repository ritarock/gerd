package action

import (
	database "github.com/ritarock/gerd/internal/db"
	"github.com/ritarock/gerd/internal/mermaid"
)

func Run(db, address, user, password string) error {
	dbx := database.NewDbx()
	if err := dbx.Connect(db, address, user, password); err != nil {
		return err
	}

	tables, err := dbx.GetTables()
	if err != nil {
		return err
	}

	mermaid.Create()
	for _, table := range tables {
		column, err := dbx.GetDescribe(table)
		if err != nil {
			return err
		}
		if err := mermaid.CreateTable(table, column); err != nil {
			return err
		}
	}

	for _, table := range tables {
		references, err := dbx.GetReferences(table)
		if err != nil {
			return err
		}
		if len(references) == 0 {
			continue
		}
		for _, reference := range references {
			if err := mermaid.CreateReference(table, reference); err != nil {
				return err
			}
		}
	}

	return nil
}
