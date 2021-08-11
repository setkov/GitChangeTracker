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
	sqlScript := fmt.Sprintf("-- %v\r\n", commitInfo)
	for _, object := range p.sqlObjects.Objects {
		sqlScript += fmt.Sprintf("-- file: %v\r\n", string(object.Path))
	}

	sqlScript += "SET XACT_ABORT ON\r\nBEGIN TRAN\r\n"
	for _, object := range p.sqlObjects.Objects {
		supported := true

		switch {
		case object.Type() == View:
			sqlScript += fmt.Sprintf("IF OBJECT_ID('%v') IS NOT NULL\r\n\tDROP VIEW %v\r\nGO\r\n", object.Name(), object.Name())
		case object.Type() == Function:
			sqlScript += fmt.Sprintf("IF OBJECT_ID('%v') IS NOT NULL\r\n\tDROP FUNCTION %v\r\nGO\r\n", object.Name(), object.Name())
		case object.Type() == Procedure:
			sqlScript += fmt.Sprintf("IF OBJECT_ID('%v') IS NOT NULL\r\n\tDROP PROCEDURE %v\r\nGO\r\n", object.Name(), object.Name())
		default:
			sqlScript += fmt.Sprintf("-- %v - object scripting not supported !!!\r\n", object.Name())
			supported = false
		}
		if supported {
			sqlScript += fmt.Sprintf("%v\r\nGO\r\n", strings.TrimSuffix(strings.TrimSpace(object.Code), "\r\nGO"))
		}
	}
	sqlScript += "COMMIT TRAN\r\nGO\r\n"

	return sqlScript
}
