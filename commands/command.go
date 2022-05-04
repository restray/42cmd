package commands

type FtCommand interface {
	GetCommand() string
	GetDescription() string
	GetAlias() []string
	Init()
	Handler(args []string)
}
