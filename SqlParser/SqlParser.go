package SqlParser

import (
	"fmt"
	"strings"
)

// sql parser
type SqlParser struct {
	sqlObjects *SqlObjects
	outputPath string
}

// new sql parser
func NewSqlParser(sqlObjects *SqlObjects, outputPath string) *SqlParser {
	return &SqlParser{
		sqlObjects: sqlObjects,
		outputPath: outputPath,
	}
}

// generate sql script
func (p *SqlParser) Parse(commitInfo string) string {
	sqlScript := fmt.Sprintf("-- %v\n", commitInfo)
	for _, object := range p.sqlObjects.Objects {
		sqlScript += fmt.Sprintf("-- file: %v\n", string(object.Path))
	}

	sqlScript += "SET XACT_ABORT ON\nBEGIN TRAN\n"
	for _, object := range p.sqlObjects.Objects {
		supported := true

		switch {
		case object.Type() == View:
			sqlScript += fmt.Sprintf("IF OBJECT_ID('%v') IS NOT NULL\n\tDROP VIEW %v\nGO\n", object.Name(), object.Name())
		case object.Type() == Function:
			sqlScript += fmt.Sprintf("IF OBJECT_ID('%v') IS NOT NULL\n\tDROP FUNCTION %v\nGO\n", object.Name(), object.Name())
		case object.Type() == Procedure:
			sqlScript += fmt.Sprintf("IF OBJECT_ID('%v') IS NOT NULL\n\tDROP PROCEDURE %v\nGO\n", object.Name(), object.Name())
		default:
			sqlScript += fmt.Sprintf("-- %v - object scripting not supported !!!\n", object.Name())
			supported = false
		}
		if supported {
			sqlScript += fmt.Sprintf("%v\nGO\n", strings.TrimSuffix(strings.TrimSpace(object.Code), "\nGO"))
		}
	}
	sqlScript += "COMMIT TRAN\nGO\n"

	return sqlScript
}
