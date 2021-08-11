package SqlParser

import (
	"path/filepath"
	"strings"
)

type SqlObjectType string

const (
	Table     SqlObjectType = "Table"
	View      SqlObjectType = "View"
	Function  SqlObjectType = "Function"
	Procedure SqlObjectType = "Procedure"
	Other     SqlObjectType = "Other"
)

type SqlObject struct {
	Path string
	Code string
}

// get type
func (o *SqlObject) Type() SqlObjectType {
	switch {
	case strings.Contains(o.Path, "Tables/"):
		return Table
	case strings.Contains(o.Path, "Views/"):
		return View
	case strings.Contains(o.Path, "Programmability/Functions/"):
		return Function
	case strings.Contains(o.Path, "Programmability/Procedures/"):
		return Procedure
	default:
		return Other
	}
}

// get name
func (o *SqlObject) Name() string {
	base := filepath.Base(o.Path)
	if pos := strings.LastIndexByte(base, '.'); pos != -1 {
		return base[:pos]
	}
	return base
}

type SqlObjects struct {
	Objects []SqlObject
}
