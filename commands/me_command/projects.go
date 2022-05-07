package me_command

import (
	"flag"
	"fmt"
	"main/intrapi"
)

type FtCommandProject struct {
	flags *flag.FlagSet
	arg   bool
}

func (cmd *FtCommandProject) GetCommand() string {
	return "projects"
}

func (cmd *FtCommandProject) GetAlias() []string {
	return []string{"p"}
}

func (cmd *FtCommandProject) GetFlags() *flag.FlagSet {
	return cmd.flags
}

func (cmd *FtCommandProject) GetDescription() string {
	return "Get all my current projects"
}

func (cmd *FtCommandProject) Init() {
}

func (cmd *FtCommandProject) DefaultOutput() {
	fmt.Println("You should create a default output for this!")
}

func (cmd *FtCommandProject) Handler(args []string) {
	projects := intrapi.GetMeProjects()
	fmt.Println(projects)

	// Doing something here!
}
