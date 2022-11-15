package checker

import (
	"fmt"

	"github.com/tufin/oasdiff/diff"
	"golang.org/x/exp/slices"
)

func ResponseRequiredPropertyBecameNonWriteOnlyCheck(diffReport *diff.Diff, operationsSources *diff.OperationsSourcesMap) []BackwardCompatibilityError {
	result := make([]BackwardCompatibilityError, 0)
	if diffReport.PathsDiff == nil {
		return result
	}
	for path, pathItem := range diffReport.PathsDiff.Modified {
		if pathItem.OperationsDiff == nil {
			continue
		}
		for operation, operationItem := range pathItem.OperationsDiff.Modified {
			source := (*operationsSources)[operationItem.Revision]

			if operationItem.ResponsesDiff == nil {
				continue
			}

			for responseStatus, responseDiff := range operationItem.ResponsesDiff.Modified {
				if responseDiff.ContentDiff == nil ||
					responseDiff.ContentDiff.MediaTypeModified == nil {
					continue
				}

				modifiedMediaTypes := responseDiff.ContentDiff.MediaTypeModified
				for _, mediaTypeDiff := range modifiedMediaTypes {
					if mediaTypeDiff.SchemaDiff == nil {
						continue
					}

					CheckModifiedPropertiesDiff(
						mediaTypeDiff.SchemaDiff,
						func(propertyPath string, propertyName string, propertyDiff *diff.SchemaDiff, parent *diff.SchemaDiff) {
							writeOnlyDiff := propertyDiff.WriteOnlyDiff
							if writeOnlyDiff == nil {
								return
							}
							if writeOnlyDiff.From != true {
								return
							}
							if parent.Revision.Value.Properties[propertyName] == nil {
								// removed properties processed by the ResponseRequiredPropertyRemovedCheck check
								return
							}
							if !slices.Contains(parent.Base.Value.Required, propertyName) {
								// removed properties processed by the ResponseRequiredPropertyRemovedCheck check
								return
							}

							result = append(result, BackwardCompatibilityError{
								Id:        "response-required-property-became-not-write-only",
								Level:     WARN,
								Text:      fmt.Sprintf("the response required property %s became not write-only for the status %s", ColorizedValue(propertyFullName(propertyPath, propertyName)), ColorizedValue(responseStatus)),
								Comment:   "It is valid if the property was always returned before the specification has been changed",
								Operation: operation,
								Path:      path,
								Source:    source,
								ToDo:      "Add to exceptions-list.md",
							})
						})
				}

			}
		}
	}
	return result
}
