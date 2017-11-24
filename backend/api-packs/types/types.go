package types

type EndpointSignature struct {
	Method string
	Path   string
	Input  []Field
	Output Field
}

type Field struct {
	Import Import // empty for local packages
	IsList bool
	Name   string
	Type   string // not recursive
}

type Type struct {
	Name   string
	Fields []Field
}

type Context struct {
	ProjectName string
	Types       map[string]Type
	Imports     []Import
}

type Import struct {
	Author      string
	ProjectName string
}
