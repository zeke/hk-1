package main

import (
	"strconv"
	"strings"
	"time"
)

// See https://github.com/heroku/api-doc#apps
type App struct {
	Id     string
	Name   string
	Stack  struct {
		Id   string
		Name string
	}
	WebURL string `json:"web_url"`
	GitURL string `json:"git_url"`
	Owner  struct {
		Id    string
		Email string
	}
	Region struct {
		Id   string
		Name string
	}
	RepoSize    *int       `json:"repo_size"`
	SlugSize    *int       `json:"slug_size"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	ReleasedAt  *time.Time `json:"released_at"`
	Maintenance bool

	BuildpackProvidedDescription *NullString `json:"buildpack_provided_description"`
}

type Dyno struct {
	Name    string
	Id      string
	Type    string
	Command string
	AppName string `json:"app_name"`
	Release struct {
		Id      string
		Version int
	}
	Size      string
	State     string
	AttachURL *string   `json:"attach_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Dyno) Age() time.Duration {
	return time.Now().Sub(d.UpdatedAt)
}

func (d *Dyno) Seq() int {
	i, _ := strconv.Atoi(strings.TrimPrefix(d.Name, d.Type+"."))
	return i
}

type Release struct {
	Id   string
	User struct {
		Id    string
		Email string
	}
	Slug struct {
		Id string
	}
	Description string
	Version     int
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Who    string // same as User.Email or abbreviated
	Commit string // deduced from Description, if possible
}

type Addon struct {
	Id   string
	Plan struct {
		Id   string
		Name string
	}
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Resource struct {
	URL   *NullString
	Price struct {
		Unit  string
		Cents int
	}
	TermsOfService    bool   `json:"terms_of_service"`
	GroupDescription  string `json:"group_description"`
	Configured        bool
	State             string
	SSOURL            *NullString `json:"sso_url"`
	Attachable        bool
	Description       string
	ConsumesDynoHours bool
	Selective         bool
	Beta              bool
	Name              string
	Slug              string
}

type Attachment struct {
	ConfigVar string `json:"config_var"`
	App       struct {
		Owner string
		Id    string
		Name  string
	}
	Resource struct {
		Name       string
		Type       string
		Id         string
		Value      string
		SSOURL     *NullString `json:"sso_url"`
		BillingApp struct {
			Name  string
			Id    string
			Owner string
		} `json:"billing_app"`
	}
}

type LogSession struct {
	LogplexURL string `json:"logplex_url"`
	CreatedAt  time.Time
}

type NullString string

func (s *NullString) String() string {
	if s == nil {
		return "(null)"
	}
	return string(*s)
}
