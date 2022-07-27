package intrapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Cursus struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Kind      string    `json:"kind"`
}

type Projects []Project
type Project struct {
	ID              int           `json:"id"`
	Name            string        `json:"name"`
	Slug            string        `json:"slug"`
	Parent          interface{}   `json:"parent"`
	Children        []interface{} `json:"children"`
	Attachments     []interface{} `json:"attachments"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Exam            bool          `json:"exam"`
	GitID           int           `json:"git_id"`
	Repository      string        `json:"repository"`
	Recommendation  string        `json:"recommendation"`
	Cursus          []Cursus      `json:"cursus"`
	Videos          []interface{} `json:"videos"`
	ProjectSessions []struct {
		ID               int         `json:"id"`
		Solo             bool        `json:"solo"`
		BeginAt          interface{} `json:"begin_at"`
		EndAt            interface{} `json:"end_at"`
		EstimateTime     string      `json:"estimate_time"`
		Difficulty       int         `json:"difficulty"`
		Objectives       []string    `json:"objectives"`
		Description      string      `json:"description"`
		DurationDays     interface{} `json:"duration_days"`
		TerminatingAfter interface{} `json:"terminating_after"`
		ProjectID        int         `json:"project_id"`
		CampusID         int         `json:"campus_id"`
		CursusID         int         `json:"cursus_id"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		MaxPeople        interface{} `json:"max_people"`
		IsSubscriptable  bool        `json:"is_subscriptable"`
		Scales           []struct {
			ID               int  `json:"id"`
			CorrectionNumber int  `json:"correction_number"`
			IsPrimary        bool `json:"is_primary"`
		} `json:"scales"`
		Uploads       []interface{} `json:"uploads"`
		TeamBehaviour string        `json:"team_behaviour"`
		Commit        string        `json:"commit"`
	} `json:"project_sessions"`
}

func GetProject(id int) Project {
	result := makeAPIReq(fmt.Sprintf("/projects/%d", id))

	var project Project
	json.Unmarshal(result, &project)

	return project
}

func GetProjectFromName(name string) (Project, error) {
	projects := GetMeProjects(nil, CURSUS_42CURSUS)

	for _, project := range projects {
		if strings.ToLower(project.Project.Name) == strings.ToLower(name) {
			return project.Project, nil
		}
	}
	return Project{}, errors.New("No project found with this name.")
}

func (p *Projects) toInterface() []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(*p))
	for i, d := range *p {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

func (p *Project) GetSubject() []interface{} {
	fmt.Println(p.ProjectSessions[0])
	result := makeAPIReq(fmt.Sprintf("/attachments"))
	fmt.Println(string(result))

	var attachments []interface{}
	json.Unmarshal(result, &attachments)

	fmt.Println(attachments)
	return attachments
}
