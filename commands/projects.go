package commands

import (
	"flag"
	"fmt"
	"log"
	"main/intrapi"
	"strconv"
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
	cmd.flags.Parse(args[1:])

	if len(cmd.flags.Args()) == 1 {
		id, err := strconv.Atoi(cmd.flags.Arg(0))

		var project intrapi.Project

		if err != nil {
			project, err = intrapi.GetProjectFromName(cmd.flags.Arg(0))
			if err != nil {
				log.Fatal(err)
			}
			id = project.ID
		}

		project = intrapi.GetProject(id)

		if project.Name == "" {
			log.Fatal("No project found.")
		}

		fmt.Printf("%d %s\n", project.ID, project.Name)

		return
	}

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

	fmt.Printf("Projects you are looking for:\n\n")
	fmt.Println(projects)
	fmt.Println("\nUse the following command to get more information about a project")
	fmt.Printf("ftcli projects {id}\n\n")
}
