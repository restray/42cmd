package intrapi

import (
	"encoding/json"
	"fmt"
	"main/commands/utils"
	"time"
)

type User42 struct {
	ID              int           `json:"id"`
	Email           string        `json:"email"`
	Login           string        `json:"login"`
	FirstName       string        `json:"first_name"`
	LastName        string        `json:"last_name"`
	UsualFullName   string        `json:"usual_full_name"`
	UsualFirstName  interface{}   `json:"usual_first_name"`
	URL             string        `json:"url"`
	Phone           string        `json:"phone"`
	Displayname     string        `json:"displayname"`
	ImageURL        string        `json:"image_url"`
	NewImageURL     string        `json:"new_image_url"`
	Staff           bool          `json:"staff?"`
	CorrectionPoint int           `json:"correction_point"`
	PoolMonth       string        `json:"pool_month"`
	PoolYear        string        `json:"pool_year"`
	Location        interface{}   `json:"location"`
	Wallet          int           `json:"wallet"`
	AnonymizeDate   time.Time     `json:"anonymize_date"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	Alumni          bool          `json:"alumni"`
	IsLaunched      bool          `json:"is_launched?"`
	Groups          []interface{} `json:"groups"`
	CursusUsers     []struct {
		Grade  interface{} `json:"grade"`
		Level  float64     `json:"level"`
		Skills []struct {
			ID    int     `json:"id"`
			Name  string  `json:"name"`
			Level float64 `json:"level"`
		} `json:"skills"`
		BlackholedAt interface{} `json:"blackholed_at"`
		ID           int         `json:"id"`
		BeginAt      time.Time   `json:"begin_at"`
		EndAt        time.Time   `json:"end_at"`
		CursusID     int         `json:"cursus_id"`
		HasCoalition bool        `json:"has_coalition"`
		CreatedAt    time.Time   `json:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at"`
		User         struct {
			ID              int         `json:"id"`
			Email           string      `json:"email"`
			Login           string      `json:"login"`
			FirstName       string      `json:"first_name"`
			LastName        string      `json:"last_name"`
			UsualFullName   string      `json:"usual_full_name"`
			UsualFirstName  interface{} `json:"usual_first_name"`
			URL             string      `json:"url"`
			Phone           string      `json:"phone"`
			Displayname     string      `json:"displayname"`
			ImageURL        string      `json:"image_url"`
			NewImageURL     string      `json:"new_image_url"`
			Staff           bool        `json:"staff?"`
			CorrectionPoint int         `json:"correction_point"`
			PoolMonth       string      `json:"pool_month"`
			PoolYear        string      `json:"pool_year"`
			Location        interface{} `json:"location"`
			Wallet          int         `json:"wallet"`
			AnonymizeDate   time.Time   `json:"anonymize_date"`
			CreatedAt       time.Time   `json:"created_at"`
			UpdatedAt       time.Time   `json:"updated_at"`
			Alumni          bool        `json:"alumni"`
			IsLaunched      bool        `json:"is_launched?"`
		} `json:"user"`
		Cursus struct {
			ID        int       `json:"id"`
			CreatedAt time.Time `json:"created_at"`
			Name      string    `json:"name"`
			Slug      string    `json:"slug"`
		} `json:"cursus"`
		Launcher interface{} `json:"launcher"`
	} `json:"cursus_users"`
	CurrentProjects ProjectsUsers `json:"projects_users"`
	LanguagesUsers  []struct {
		ID         int       `json:"id"`
		LanguageID int       `json:"language_id"`
		UserID     int       `json:"user_id"`
		Position   int       `json:"position"`
		CreatedAt  time.Time `json:"created_at"`
	} `json:"languages_users"`
	Achievements []struct {
		ID           int         `json:"id"`
		Name         string      `json:"name"`
		Description  string      `json:"description"`
		Tier         string      `json:"tier"`
		Kind         string      `json:"kind"`
		Visible      bool        `json:"visible"`
		Image        string      `json:"image"`
		NbrOfSuccess interface{} `json:"nbr_of_success"`
		UsersURL     string      `json:"users_url"`
	} `json:"achievements"`
	Titles          []interface{} `json:"titles"`
	TitlesUsers     []interface{} `json:"titles_users"`
	Partnerships    []interface{} `json:"partnerships"`
	Patroned        []interface{} `json:"patroned"`
	Patroning       []interface{} `json:"patroning"`
	ExpertisesUsers []interface{} `json:"expertises_users"`
	Roles           []interface{} `json:"roles"`
	Campus          []struct {
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
		UsersCount         int    `json:"users_count"`
		VogsphereID        int    `json:"vogsphere_id"`
		Country            string `json:"country"`
		Address            string `json:"address"`
		Zip                string `json:"zip"`
		City               string `json:"city"`
		Website            string `json:"website"`
		Facebook           string `json:"facebook"`
		Twitter            string `json:"twitter"`
		Active             bool   `json:"active"`
		EmailExtension     string `json:"email_extension"`
		DefaultHiddenPhone bool   `json:"default_hidden_phone"`
	} `json:"campus"`
	CampusUsers []struct {
		ID        int       `json:"id"`
		UserID    int       `json:"user_id"`
		CampusID  int       `json:"campus_id"`
		IsPrimary bool      `json:"is_primary"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"campus_users"`
	Launcher interface{} `json:"launcher"`
}

func (u User42) formatLocation() string {
	if u.Location != nil {
		return fmt.Sprint(u.Location)
	}
	return "Not at school"
}

func (u User42) String() string {
	output := ""
	output += fmt.Sprintf("\n    ██╗  ██╗██████╗\n")
	output += fmt.Sprintf("    ██║  ██║╚════██╗	%s <%s>\n", u.UsualFullName, u.Login)
	output += fmt.Sprintf("    ███████║ █████╔╝	%s\n", u.formatLocation())
	output += fmt.Sprintf("    ╚════██║██╔═══╝ 	%d Correction points\n", u.CorrectionPoint)
	output += fmt.Sprintf("         ██║███████╗	lvl %.0f : lvl 0 [%s] lvl 21\n", u.CursusUsers[1].Level, utils.LoadingBar(0, 21, int(u.CursusUsers[1].Level)))
	output += fmt.Sprintf("         ╚═╝╚══════╝\n")
	return output
}

func GetMe() User42 {
	result := makeAPIReq("/me")

	var user User42
	json.Unmarshal(result, &user)
	return user
}

func GetUserByName(name string) *User42 {
	result := makeAPIReq("/users/" + name)

	user := new(User42)
	json.Unmarshal(result, user)
	if user.ID > 0 {
		return user
	}
	return nil
}
