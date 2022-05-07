package intrapi

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"
)

type markedTime struct {
	*time.Time
}

type ProjectsUsers []ProjectUsers
type ProjectUsers struct {
	ID            int       `json:"id"`
	Occurrence    int       `json:"occurrence"`
	FinalMark     int       `json:"final_mark"`
	Status        string    `json:"status"`
	Validated     bool      `json:"validated?"`
	CurrentTeamID int       `json:"current_team_id"`
	Project       Project   `json:"project"`
	CursusIds     []int     `json:"cursus_ids"`
	MarkedAt      time.Time `json:"marked_at"`
	Marked        bool      `json:"marked"`
	RetriableAt   time.Time `json:"retriable_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (t markedTime) String() string {
	if t.IsZero() {
		return "InProgress"
	}
	return t.Format("01/02/2006")
}

func GetMeProjects() ProjectsUsers {
	result := makeAPIReq("/me")

	var user User42
	json.Unmarshal(result, &user)
	return user.CurrentProjects
}

func GetProjectsUsers() ProjectsUsers {
	result := makeAPIReq("/me/projects")

	var projects ProjectsUsers
	json.Unmarshal(result, &projects)
	return projects
}

func (projects ProjectsUsers) String() string {
	ret := ""

	sort.Slice(projects, func(i, j int) bool {
		if projects[i].MarkedAt.IsZero() {
			return true
		}
		return projects[i].MarkedAt.After(projects[j].MarkedAt)
	})

	for _, v := range projects {
		ret += fmt.Sprintln(v)
	}

	return ret
}

func (projectu ProjectUsers) String() string {
	return fmt.Sprintf(" %d - %s - %s (ProjectID: %d)", projectu.ID, markedTime{&projectu.MarkedAt}, projectu.Project.Name, projectu.Project.ID)
}
