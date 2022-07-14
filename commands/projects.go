package commands

import (
	"flag"
	"fmt"
	"main/intrapi"
)

type FtCommandProject struct {
	flags      *flag.FlagSet
	inProgress bool
	available  bool
	finished   bool
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
	cmd.flags = flag.NewFlagSet("projects", flag.ExitOnError)
	cmd.flags.BoolVar(&cmd.available, "available", false, "Show available projects")
	cmd.flags.BoolVar(&cmd.inProgress, "in-progress", true, "Show in progress projects")
	cmd.flags.BoolVar(&cmd.finished, "finished", false, "Show finished projects")
}

func (cmd *FtCommandProject) DefaultOutput() {
	fmt.Println("You should create a default output for this!")
}

func (cmd *FtCommandProject) Handler(args []string) {
	cmd.flags.Parse(args)

	project_searched := make([]intrapi.ProjectStatus, 0)

	// if cmd.available {
	// 	project_searched = append(project_searched, intrapi.PROJECT_AVAILABLE)
	// }
	if cmd.inProgress {
		project_searched = append(project_searched, intrapi.PROJECT_IN_PROGRESS)
	}
	if cmd.finished {
		project_searched = append(project_searched, intrapi.PROJECT_FINISHED)
	}

	projects := intrapi.GetMeProjects(project_searched, intrapi.CURSUS_42CURSUS)

	fmt.Println("Projects you are looking for:\n")
	fmt.Println(projects)
	fmt.Println("\nUse the following command to get more information about a project")
	fmt.Println("ftcli me project {id}\n")
}
