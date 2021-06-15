package test

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	New,
)
