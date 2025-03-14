package types

type Config struct {
	SelectedEnvironment string                  `json:"selectedEnvironment" yaml:"selectedEnvironment"`
	Environments        map[string]*Environment `json:"environments"        yaml:"environments"`
}

type Environment struct {
	URL          string           `json:"url"          yaml:"url"`
	ShortName    string           `json:"shortName"    yaml:"shortName"`
	Description  string           `json:"description"  yaml:"description"`
	SelectedUser string           `json:"selectedUser" yaml:"selectedUser"`
	Users        map[string]*User `json:"users"        yaml:"users"`
}

type User struct {
	Slug  string `json:"slug"  yaml:"slug"`
	Token string `json:"token" yaml:"token"`
}
