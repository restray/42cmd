package commands

import (
	"flag"
	"fmt"
)

type FtCommandUser struct {
	flags *flag.FlagSet
	arg   bool
}

func (me *FtCommandUser) GetCommand() string {
	return "user"
}

func (me *FtCommandUser) GetAlias() []string {
	return []string{"u", "profile"}
}

func (me *FtCommandUser) GetFlags() *flag.FlagSet {
	return me.flags
}

func (me *FtCommandUser) GetDescription() string {
	return "Give an user"
}

func (me *FtCommandUser) Init() {
	me.arg = false

	me.flags = flag.NewFlagSet(me.GetCommand(), flag.ExitOnError)
	me.flags.BoolVar(&me.arg, "arg", false, "Get the arg value (false by default)")
}

func (me *FtCommandUser) DefaultOutput() {
	fmt.Println("You should create a default output for this!")
}

func (me *FtCommandUser) Handler(args []string) {
	if len(args) <= 0 {
		me.DefaultOutput()
		return
	}

	// Doing something here!
}
