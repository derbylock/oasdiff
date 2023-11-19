package checker_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tufin/oasdiff/checker"
)

var componentChange = checker.ComponentChange{
	Id:              "id",
	Text:            "text",
	Comment:         "comment",
	Level:           checker.ERR,
	Source:          "source",
	Component:       "component",
	SourceFile:      "sourceFile",
	SourceLine:      1,
	SourceLineEnd:   2,
	SourceColumn:    3,
	SourceColumnEnd: 4,
}

func TestComponentChange(t *testing.T) {
	require.Equal(t, "comment", componentChange.GetComment())
	require.Equal(t, "", componentChange.GetOperationId())
	require.Equal(t, "sourceFile", componentChange.GetSourceFile())
	require.Equal(t, 1, componentChange.GetSourceLine())
	require.Equal(t, 2, componentChange.GetSourceLineEnd())
	require.Equal(t, 3, componentChange.GetSourceColumn())
	require.Equal(t, 4, componentChange.GetSourceColumnEnd())
	require.Equal(t, "error, in components/component text [id]. comment", componentChange.LocalizedError(checker.NewDefaultLocalizer()))
	require.Equal(t, "error, in components/component text [id]. comment", componentChange.PrettyErrorText(checker.NewDefaultLocalizer()))
	require.Equal(t, "error, in components/component text [id]. comment", componentChange.Error())
}

func TestComponentChange_MatchIgnore(t *testing.T) {
	require.True(t, componentChange.MatchIgnore("", "error, in components/component text [id]. comment"))
}

func TestComponentChange_PrettyPiped(t *testing.T) {
	piped := true
	save := checker.SetPipedOutput(&piped)
	defer checker.SetPipedOutput(save)
	require.Equal(t, "error, in components/component text [id]. comment", componentChange.PrettyErrorText(checker.NewDefaultLocalizer()))
}