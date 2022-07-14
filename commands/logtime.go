package commands

import (
	"flag"
	"fmt"
)

type FtCommandLogtime struct {
	flags *flag.FlagSet
	arg   bool
}

func (cmd *FtCommandLogtime) GetCommand() string {
	return "logtime"
}

func (cmd *FtCommandLogtime) GetAlias() []string {
	return []string{}
}

func (cmd *FtCommandLogtime) GetFlags() *flag.FlagSet {
	return cmd.flags
}

func (cmd *FtCommandLogtime) GetDescription() string {
	return "Get the school logtime"
}

func (cmd *FtCommandLogtime) Init() {
	cmd.arg = false

	cmd.flags = flag.NewFlagSet(cmd.GetCommand(), flag.ExitOnError)
	cmd.flags.BoolVar(&cmd.arg, "arg", false, "Get the arg value (false by default)")
}

func (cmd *FtCommandLogtime) DefaultOutput() {
	fmt.Println("You should create a default output for this!")
}

func (cmd *FtCommandLogtime) Handler(args []string) {
	if len(args) <= 0 {
		cmd.DefaultOutput()
		return
	}

	// Doing something here!
}
