package main

import (
	"database/sql"

	"github.com/lemmego/migration"
)

func init() {
	migration.GetMigrator().AddMigration(&migration.Migration{
		Version: "20230222004736",
		Up:      mig_20230222004736_create_orgs_table_up,
		Down:    mig_20230222004736_create_orgs_table_down,
	})
}

func mig_20230222004736_create_orgs_table_up(tx *sql.Tx) error {
	schema := migration.NewSchema().Create("orgs", func(t *migration.Table) error {
		t.BigIncrements("id").Primary()
		t.String("name", 255)
		t.String("subdomain", 255)
		t.String("email", 255)
		t.DateTime("created_at").Default("now()")
		t.DateTime("updated_at").Default("now()")
		t.Unique("subdomain", "email")
		return nil
	}).Build()

	_, err := tx.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}

func mig_20230222004736_create_orgs_table_down(tx *sql.Tx) error {
	schema := migration.NewSchema().Drop("orgs").Build()
	_, err := tx.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
