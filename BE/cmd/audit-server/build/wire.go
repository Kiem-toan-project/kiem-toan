// +build wireinject

package build

import (
	"github.com/google/wire"
	"github.com/kiem-toan/cmd/audit-server/config"
	"github.com/kiem-toan/pkg/database"
	"github.com/kiem-toan/pkg/test"
)

func InitApp(cfg config.Config) (*App, error){
	wire.Build(
		test.Provider,
		database.New,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}
type App struct {
	Test *test.TestStr
	Db *database.Database
}