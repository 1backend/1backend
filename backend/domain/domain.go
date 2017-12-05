package domain

import (
	"time"

	shortid "github.com/ventu-io/go-shortid"
)

var Sid *shortid.Shortid

func init() {
	var err error
	Sid, err = shortid.New(1, shortid.DefaultABC, 2719)
	if err != nil {
		panic(err)
	}
}

type Project struct {
	Id string
	// Password for infrastructural elements like mysql, redis etc
	InfraPassword string `json:"-"`
	// Author name. Could be a user id, but it's not due to cruft.
	Author string
	Name   string
	Stars  int
	// Semver version of the project.
	Version string
	ReadMe  string
	// The port of the running instance for internal use.
	// Might have to extend this contept to support multiple instances.
	Port int
	// Private repos can only be read by their owner (no organisations yet)
	Public bool
	// Source code is visible to those who can view the project
	OpenSource bool
	// Short description of the project)
	Description string
	// Eg. go, nodejs, docker
	Mode string
	// Different setups inside the same tech (see language)
	Recipe string
	// Secrets as envars for open source applications with private keys
	Secrets string
	// eg. requires for nodejs or github imports for go apps
	Imports string
	// eg. package.json for nodejs apps
	Packages string
	// JSON that defines the types belonging to this project
	Types string
	// Source contains source code if the project mode/recipe does not use the code from the endpoints
	Source string
	// Github ssh for private code
	Endpoints    []Endpoint
	Builds       []Build
	Dependencies []Dependency
	CreatedAt    time.Time
	UpdatedAt    time.Time
	// Very stupid, I know
	Starrers []User `gorm:"many2many:stars;"`
}

type Dependency struct {
	Id        string
	ProjectId string
	Config    string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Endpoint struct {
	Id          string
	Url         string
	Description string
	Method      string
	Code        string
	// JSON that defines the input structure
	Input string
	// JSON that defines the output structure
	Output    string
	ProjectId string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Build struct {
	Id        string
	Output    string
	Success   bool
	ProjectId string
	// Version of the project at the time of the build
	Version    string
	InProgress bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type User struct {
	Id         string
	Password   string `json:"-"`
	Nick       string
	Name       string
	Email      string
	AvatarLink string
	Tokens     []Token
	Premium    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// For logins
type AccessToken struct {
	Id        string
	Token     string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// service access token
type Token struct {
	Id          string
	Token       string
	UserId      string
	Name        string // eg. "test", "xyz-app" etc. max 32 chars
	Description string
	Quota       int64 `gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Issue struct {
	Id           string
	ProjectId    string
	Title        string
	UserId       string
	User         User
	Comments     []Comment
	CommentCount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Comment struct {
	Id        string
	IssueId   string
	Content   string
	UserId    string
	User      User
	Issue     Issue
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Star struct {
	Id        string
	UserId    string
	ProjectId string
}

type Charge struct {
	Id          string
	Amount      int
	UserId      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Post struct {
	Id        string
	Title     string
	Slug      string
	Content   string
	UserId    string
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reset struct {
	Id        string
	Secret    string
	UserId    string
	CreatedAt time.Time
	Used      bool
}
