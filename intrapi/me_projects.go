package intrapi

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

type ProjectStatus int

const (
	PROJECT_IN_PROGRESS ProjectStatus = iota
	PROJECT_FINISHED
)

var projectStatus = map[ProjectStatus]string{
	PROJECT_IN_PROGRESS: "in_progress",
	PROJECT_FINISHED:    "finished",
}

type MeProjects []MeProject
type MeProject struct {
	ID            int         `json:"id"`
	Occurrence    int         `json:"occurrence"`
	FinalMark     interface{} `json:"final_mark"`
	Status        string      `json:"status"`
	Validated     interface{} `json:"validated?"`
	CurrentTeamID int         `json:"current_team_id"`
	Project       Project     `json:"project"`
	CursusIds     []int       `json:"cursus_ids"`
	User          struct {
		ID    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"user"`
	Teams []struct {
		ID            int         `json:"id"`
		Name          string      `json:"name"`
		URL           string      `json:"url"`
		FinalMark     interface{} `json:"final_mark"`
		ProjectID     int         `json:"project_id"`
		CreatedAt     time.Time   `json:"created_at"`
		UpdatedAt     time.Time   `json:"updated_at"`
		Status        string      `json:"status"`
		TerminatingAt interface{} `json:"terminating_at"`
		Users         []struct {
			ID             int    `json:"id"`
			Login          string `json:"login"`
			URL            string `json:"url"`
			Leader         bool   `json:"leader"`
			Occurrence     int    `json:"occurrence"`
			Validated      bool   `json:"validated"`
			ProjectsUserID int    `json:"projects_user_id"`
		} `json:"users"`
		Locked           bool        `json:"locked?"`
		Validated        interface{} `json:"validated?"`
		Closed           bool        `json:"closed?"`
		RepoURL          interface{} `json:"repo_url"`
		RepoUUID         string      `json:"repo_uuid"`
		LockedAt         time.Time   `json:"locked_at"`
		ClosedAt         interface{} `json:"closed_at"`
		ProjectSessionID int         `json:"project_session_id"`
	} `json:"teams"`
}

func GetProjects() MeProjects {
	result := makeAPIReq("/me/projects")

	var projects MeProjects
	json.Unmarshal(result, &projects)
	return projects
}

func formatProjectStatus(status []ProjectStatus) string {
	ret := ""
	for _, v := range status {
		if len(ret) > 0 {
			ret += ","
		}
		ret += projectStatus[v]
	}
	return ret
}

type MeProjectsParams struct {
	Status []ProjectStatus
	Cursus cursusId
}

func GetMeProjects(status []ProjectStatus, cursus cursusId) MeProjects {
	user := GetMe()

	query := fmt.Sprintf("cursus_id=%d", cursus-1)
	if len(status) > 0 {
		first := true
		query += "&filter[status]="
		for _, stat := range status {
			if !first {
				query += ","
			}
			query += projectStatus[stat]
			first = false
		}
	}
	result := makeAPIReq(fmt.Sprintf("/users/%d/projects_users?%s", user.ID, query))

	var projects MeProjects
	if err := json.Unmarshal(result, &projects); err != nil {
		log.Fatal("Error on parsing JSON", err)
	}

	return projects
}

func getProjectsForCursus(projects MeProjects, cursus_searched cursusId) MeProjects {
	sortedProjects := make(MeProjects, 0)
	for _, v := range projects {
		for _, cursus := range v.CursusIds {
			if cursus == int(cursus_searched)-1 {
				sortedProjects = append(sortedProjects, v)
			}
		}
	}
	return sortedProjects
}

func (projects MeProjects) String() string {
	tableString := &strings.Builder{}

	table := tablewriter.NewWriter(tableString)
	table.SetBorder(false)
	table.SetCenterSeparator("|")
	table.SetHeader([]string{"Id", "Team Id", "Project", "Status", "Grade"})

	for _, v := range projects {
		mark := v.FinalMark
		if mark == nil {
			mark = "-"
		}

		table.Append([]string{fmt.Sprint(v.Project.ID), fmt.Sprint(v.ID), v.Project.Name, v.Status, fmt.Sprint(mark)})
	}

	table.Render()

	return tableString.String()
}

func (projectu MeProject) String() string {
	return fmt.Sprintf("ID: %d | Project: %s (%s) | Cursus id: %d", projectu.ID, projectu.Project.Name, projectu.Status, projectu.CursusIds[0])
}
