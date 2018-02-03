package types

import "text/template"

type Build struct {
	// Returns filename of the file generated from code.tpl
	Outfile string
	// FilesToBuilt returns a list of files that one should write to the build folder
	// Returns a list of []string{file name, file content}
	FilesBuilt [][]string
	// Folder name of tech pack
	RecipePath string
}

// Provider is the interface every tech pack plugin must implement
type Plugin interface {
	// Used to set up default endpoints and such
	PreCreate() error
	// Builds files that must be written into the build folder and deployed
	Build(fm *template.FuncMap) (*Build, error)
}
