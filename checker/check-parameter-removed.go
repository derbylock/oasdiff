package checker

import (
	"fmt"

	"github.com/tufin/oasdiff/diff"
)

func ParameterRemovedCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap) []BackwardCompatibilityError {
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
			for paramLocation, paramItems := range operationItem.ParametersDiff.Deleted {
				for _, paramName := range paramItems {
					source := (*operationsSources)[operationItem.Revision]
					result = append(result, BackwardCompatibilityError{
						Id:        "parameter-removed",
						Level:     WARN,
						Text:      fmt.Sprintf("deleted the %s request parameter %s", ColorizedValue(paramLocation), ColorizedValue(paramName)),
						Operation: operation,
						Path:      path,
						Source:    source,
						ToDo:      "Add to exceptions-list.md",
					})
				}
			}
		}
	}
	return result
}
