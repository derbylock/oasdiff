package checker

import (
	"fmt"

	"github.com/tufin/oasdiff/diff"
)

func ParameterRemovedCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap, diffBC *BCDiff) []BackwardCompatibilityError {
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
					opDiff := diffBC.AddModifiedOperation(path, operation)
					if opDiff.ParametersDiff == nil {
						opDiff.ParametersDiff = &diff.ParametersDiff{}
					}
					if opDiff.ParametersDiff.Deleted == nil {
						opDiff.ParametersDiff.Deleted = make(diff.ParamNamesByLocation)
					}
					items := opDiff.ParametersDiff.Deleted[paramLocation].ToStringSet()
					items.Add(paramName)
					opDiff.ParametersDiff.Deleted[paramLocation] = items.ToStringList()
				}
			}
		}
	}
	return result
}