package asm8

// Package represents a package node.
type pkg struct {
	Path    string // package import path
	Files   []*file
	Imports map[string]*PkgImport
}

// NewPackage creates an empty package node.
func newPkg(path string) *pkg {
	ret := new(pkg)
	ret.Path = path
	return ret
}

// AddFile adds a file into the package.
func (p *pkg) AddFile(f *file) {
	p.Files = append(p.Files, f)
}
