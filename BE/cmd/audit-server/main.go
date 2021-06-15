package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"github.com/kiem-toan/cmd/audit-server/build"
	"github.com/kiem-toan/cmd/audit-server/config"
	_ "github.com/lib/pq"
)

func main() {
	config.InitFlags()
	config.ParseFlags()
	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("Error in loading config: ",err)
	}
	app ,err := build.InitApp(cfg)
	db := app.Db
	res, err := db.Db.Exec("insert into test values ('2102101', 'trường đẹp trai1')")
	pp.Println(res.RowsAffected())
	if err != nil {
		panic(err)
	}
}

