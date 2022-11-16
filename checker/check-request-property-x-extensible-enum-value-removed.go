package checker

import (
	"encoding/json"
	"fmt"

	"github.com/tufin/oasdiff/diff"
	"golang.org/x/exp/slices"
)

func RequestPropertyXExtensibleEnumValueRemovedCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap) []BackwardCompatibilityError {
	result := make([]BackwardCompatibilityError, 0)
	if diffReport.PathsDiff == nil {
		return result
	}
	for path, pathItem := range diffReport.PathsDiff.Modified {
		if pathItem.OperationsDiff == nil {
			continue
		}
		for operation, operationItem := range pathItem.OperationsDiff.Modified {
			if operationItem.RequestBodyDiff == nil ||
				operationItem.RequestBodyDiff.ContentDiff == nil ||
				operationItem.RequestBodyDiff.ContentDiff.MediaTypeModified == nil {
				continue
			}
			source := (*operationsSources)[operationItem.Revision]

			modifiedMediaTypes := operationItem.RequestBodyDiff.ContentDiff.MediaTypeModified
			for _, mediaTypeDiff := range modifiedMediaTypes {
				CheckModifiedPropertiesDiff(
					mediaTypeDiff.SchemaDiff,
					func(propertyPath string, propertyName string, propertyDiff *diff.SchemaDiff, parent *diff.SchemaDiff) {
						if propertyDiff.ExtensionsDiff == nil {
							return
						}
						if propertyDiff.ExtensionsDiff.Modified == nil {
							return
						}
						if propertyDiff.ExtensionsDiff.Modified["x-extensible-enum"] == nil {
							return
						}
						from, ok := propertyDiff.ExtensionsDiff.Modified["x-extensible-enum"].From.(json.RawMessage)
						if !ok {
							return
						}
						to, ok := propertyDiff.ExtensionsDiff.Modified["x-extensible-enum"].To.(json.RawMessage)
						if !ok {
							return
						}
						var fromSlice []string
						if err := json.Unmarshal(from, &fromSlice); err != nil {
							result = append(result, BackwardCompatibilityError{
								Id:        "unparseable-from-x-extensible-enum",
								Level:     ERR,
								Text:      fmt.Sprintf("unparseable x-extensible-enum of the request property %s", ColorizedValue(propertyFullName(propertyPath, propertyName))),
								Operation: operation,
								Path:      path,
								Source:    source,
								ToDo:      "Add to exceptions-list.md",
							})
							return
						}
						var toSlice []string
						if err := json.Unmarshal(to, &toSlice); err != nil {
							result = append(result, BackwardCompatibilityError{
								Id:        "unparseable-to-x-extensible-enum",
								Level:     ERR,
								Text:      fmt.Sprintf("unparseable x-extensible-enum of the request property %s", ColorizedValue(propertyFullName(propertyPath, propertyName))),
								Operation: operation,
								Path:      path,
								Source:    source,
								ToDo:      "Add to exceptions-list.md",
							})
							return
						}

						deletedVals := make([]string, 0)
						for _, fromVal := range fromSlice {
							if !slices.Contains(toSlice, fromVal) {
								deletedVals = append(deletedVals, fromVal)
							}
						}

						if propertyDiff.Revision.Value.ReadOnly {
							return
						}
						for _, enumVal := range deletedVals {
							result = append(result, BackwardCompatibilityError{
								Id:        "request-property-x-extensible-enum-value-removed",
								Level:     ERR,
								Text:      fmt.Sprintf("removed the x-extensible-enum value '%s' of the request property %s", enumVal, ColorizedValue(propertyFullName(propertyPath, propertyName))),
								Operation: operation,
								Path:      path,
								Source:    source,
								ToDo:      "Add to exceptions-list.md",
							})
						}
					})
			}
		}
	}
	return result
}
