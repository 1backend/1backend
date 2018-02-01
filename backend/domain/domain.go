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

// A Project is a Go, Node, TypeScript etc. application
type Project struct {
	Id string
	// Password for infrastructural elements like mysql, redis etc
	InfraPassword string `json:"-"`
	// A secret value used to map projects to namespaces
	CallerId string `json:"-"`
	// Author name.
	Author string
	// Project name - make it nice, because it appears everywhere
	Name string
	// Project namespace. Defaults to `$AUTHOR:default`
	Namespace string
	Stars     int
	// Semver version of the project.
	Version string
	// A readme that appears on the front page
	ReadMe string
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
	Issues       []Issue
	CreatedAt    time.Time
	UpdatedAt    time.Time
	// Very stupid, I know
	Starrers []User `gorm:"many2many:stars;"`
}

// A Dependency is an infrastructural element (eg. SQL server)
type Dependency struct {
	Id        string
	ProjectId string
	Config    string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// An Endpoint is a lambda as a HTTP endpoint
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

// A Build is pretty much what the name says: a verification of the project when it's saved.
type Build struct {
	Id        string
	Success   bool
	ProjectId string
	// Version of the project at the time of the build
	Version    string
	InProgress bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Build steps could be hierarchised somehow but we are going to ignore that step for now.
// Perhaps we should implement it with some kind of elegant hack rather than trying to save a tree structure in SQL.
type BuildStep struct {
	Id        string
	Title     string // ie. "Setting up Mysql"
	Output    string
	Success   bool
	BuildId   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// A User is a registered user of a 1Backend installation
type User struct {
	Id         string
	Password   string `json:"-"`
	Nick       string
	Name       string
	Email      string
	AvatarLink string
	Level      int
	Tokens     []Token
	Posts      []Post
	Premium    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// An AccessToken is used for logins.
// It is used to call 1Backend daemon endpoints (not services running on 1Backend).
type AccessToken struct {
	Id        string
	Token     string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// A Token or a service access token is like a service key
// used to call services running on 1Backend (not the 1Backend daemon itself).
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

// An Issue is just like a GitHub repo issue, except it's for 1Backend services.
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

// A Comment represents one comment out of the many that can belong to an `Issue`
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

// A Star is used to express interest in a project (just like GitHub starring).
type Star struct {
	Id        string
	UserId    string
	ProjectId string
}

// A Charge is just a bookkeeping entitity - 1Backend saves one when a user pays for platform usage
// (if the given 1Backend installation is configured as a cloud provider)
type Charge struct {
	Id          string
	Amount      int
	UserId      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// A Post is a blog post that users can publish under their profile.
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

// A Reset is used for password resets.
type Reset struct {
	Id        string
	Secret    string
	UserId    string
	CreatedAt time.Time
	Used      bool
}

// non-database types

// CodeSections is used by infra-packs (dependency plugins)
// to generate appropriate code snippets for a project
type CodeSections struct {
	Readme  string
	Globals string
	Imports string
}
