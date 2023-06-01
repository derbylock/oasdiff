package checker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/checker"
	"github.com/tufin/oasdiff/diff"
)

func TestIgnore(t *testing.T) {
	s1 := l(t, 1)
	s2 := l(t, 3)

	d, osm, err := diff.GetWithOperationsSourcesMap(&diff.Config{}, &s1, &s2)
	require.NoError(t, err)
	errs := checker.CheckBackwardCompatibility(checker.DefaultChecks(), d, osm)
	require.Equal(t, 6, len(errs))

	errs, err = checker.ProcessIgnoredBackwardCompatibilityErrors(checker.ERR, errs, "../data/ignore-err.txt")
	require.NoError(t, err)
	require.Equal(t, 5, len(errs))
}

func TestIgnorePrefixedErrs(t *testing.T) {
	errs := []checker.BackwardCompatibilityError{
		{
			Id:        "err1",
			Text:      "err1-text",
			Comment:   "",
			Level:     checker.ERR,
			Operation: "GET",
			Path:      "/objects",
			Source:    "",
		},
		{
			Id:        "err1",
			Text:      "err1-text",
			Comment:   "",
			Level:     checker.ERR,
			Operation: "GET",
			Path:      "/objects/{object_id}",
			Source:    "",
		},
	}
	errs, err := checker.ProcessIgnoredBackwardCompatibilityErrors(checker.ERR, errs, "../data/ignore-err-prefixes.txt")
	require.NoError(t, err)
	require.Equal(t, 0, len(errs))
}
