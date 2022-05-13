package action

import (
	"github.com/ritarock/gerd/internal/db"
	"github.com/ritarock/gerd/internal/mermaid"
)

func Run(user, passwd, addr, dbName string) error {
	client, err := db.Connection(user, passwd, addr, dbName)
	if err != nil {
		return err
	}
	defer client.Close()

	tables, err := db.GetShowTables(client)
	if err != nil {
		return err
	}

	mermaid.CreateFile()

	set := map[string]map[string]string{}
	for _, table := range tables {
		info, err := db.GetDescribeTable(client, table)
		if err != nil {
			panic(err)
		}
		set[table] = info
	}

	for k, v := range set {
		mermaid.CreateTableInfo(k, v)
	}

	for _, table := range tables {
		referenceTables := db.GetReferences(client, table)
		if len(referenceTables) == 0 {
			continue
		}
		for _, reference := range referenceTables {
			if err := mermaid.CreateReferenceInfo(table, reference.TableName); err != nil {
				return err
			}
		}
	}

	return nil
}
