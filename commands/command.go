package commands

import (
	"flag"
)

type FtCommand interface {
	GetCommand() string
	GetDescription() string
	GetAlias() []string
	GetFlags() *flag.FlagSet
	Init()
	Handler(args []string)
}
