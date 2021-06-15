package test

import "github.com/kiem-toan/pkg/database"

type TestStr struct {
	Db database.Database
}

func New(db *database.Database) *TestStr {
	return &TestStr{
		Db: *db,
	}
}