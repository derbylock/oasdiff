package checker

import (
	"fmt"

	"github.com/tufin/oasdiff/diff"
)

func RequestParameterEnumValueRemovedCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap) []BackwardCompatibilityError {
	result := make([]BackwardCompatibilityError, 0)
	if diffReport.PathsDiff == nil {
		return result
	}
	for path, pathItem := range diffReport.PathsDiff.Modified {
		if pathItem.OperationsDiff == nil {
			continue
		}
		for operation, operationItem := range pathItem.OperationsDiff.Modified {
			if operationItem.ParametersDiff == nil {
				continue
			}
			if operationItem.ParametersDiff.Modified == nil {
				continue
			}
			source := (*operationsSources)[operationItem.Revision]
			for paramLocation, paramItems := range operationItem.ParametersDiff.Modified {
				for paramName, paramItem := range paramItems {
					if paramItem.SchemaDiff == nil {
						continue
					}
					enumDiff := paramItem.SchemaDiff.EnumDiff
					if enumDiff == nil || enumDiff.Deleted == nil {
						continue
					}
					for _, enumVal := range enumDiff.Deleted {
						result = append(result, BackwardCompatibilityError{
							Id:        "request-parameter-enum-value-removed",
							Level:     ERR,
							Text:      fmt.Sprintf("removed the enum value %s for the %s request parameter %s", enumVal, ColorizedValue(paramLocation), ColorizedValue(paramName)),
							Operation: operation,
							Path:      path,
							Source:    source,
							ToDo:      "Add to exceptions-list.md",
						})
					}
				}
			}
		}
	}
	return result
}
