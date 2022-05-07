package intrapi

import (
	"encoding/json"
	"fmt"
	"time"
)

type MeProjects []MeProject
type MeProject struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	Parent      interface{}   `json:"parent"`
	Children    []interface{} `json:"children"`
	Objectives  []string      `json:"objectives"`
	Tier        int           `json:"tier"`
	Attachments []interface{} `json:"attachments"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Exam        bool          `json:"exam"`
	Cursus      []struct {
		ID        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		Slug      string    `json:"slug"`
	} `json:"cursus"`
	Campus []struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		TimeZone string `json:"time_zone"`
		Language struct {
			ID         int       `json:"id"`
			Name       string    `json:"name"`
			Identifier string    `json:"identifier"`
			CreatedAt  time.Time `json:"created_at"`
			UpdatedAt  time.Time `json:"updated_at"`
		} `json:"language"`
		UsersCount  int `json:"users_count"`
		VogsphereID int `json:"vogsphere_id"`
	} `json:"campus"`
	Skills []struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"skills"`
	Videos []interface{} `json:"videos"`
	Tags   []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Kind string `json:"kind"`
	} `json:"tags"`
	ProjectSessions []struct {
		ID               int         `json:"id"`
		Solo             bool        `json:"solo"`
		BeginAt          interface{} `json:"begin_at"`
		EndAt            interface{} `json:"end_at"`
		EstimateTime     int         `json:"estimate_time"`
		DurationDays     interface{} `json:"duration_days"`
		TerminatingAfter interface{} `json:"terminating_after"`
		ProjectID        int         `json:"project_id"`
		CampusID         interface{} `json:"campus_id"`
		CursusID         interface{} `json:"cursus_id"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		MaxPeople        interface{} `json:"max_people"`
		IsSubscriptable  bool        `json:"is_subscriptable"`
		Scales           []struct {
			ID               int  `json:"id"`
			CorrectionNumber int  `json:"correction_number"`
			IsPrimary        bool `json:"is_primary"`
		} `json:"scales"`
		Uploads []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"uploads"`
		TeamBehaviour string `json:"team_behaviour"`
	} `json:"project_sessions"`
}

func GetProjects() MeProjects {
	result := makeAPIReq("/me/projects")

	var projects MeProjects
	json.Unmarshal(result, &projects)
	return projects
}

func (projects MeProjects) String() string {
	ret := ""

	for _, v := range projects {
		ret += fmt.Sprintln(v)
	}

	return ret
}

func (projectu MeProject) String() string {
	return fmt.Sprintf("ID: %d | Project: %s", projectu.ID, projectu.Name)
}
