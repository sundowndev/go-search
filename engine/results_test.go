package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResults(t *testing.T) {
	assert := assert.New(t)

	t.Run("SortResultsByScore", func(t *testing.T) {
		results := []*QueryResult{
			{
				File:       "/file",
				Score:      1,
				FirstMatch: "",
			},
			{
				File:       "/file",
				Score:      3,
				FirstMatch: "",
			},
			{
				File:       "/file",
				Score:      9,
				FirstMatch: "",
			},
			{
				File:       "/file",
				Score:      2,
				FirstMatch: "",
			},
		}

		expected := []*QueryResult{
			{
				File:       "/file",
				Score:      9,
				FirstMatch: "",
			},
			{
				File:       "/file",
				Score:      3,
				FirstMatch: "",
			},
			{
				File:       "/file",
				Score:      2,
				FirstMatch: "",
			},
			{
				File:       "/file",
				Score:      1,
				FirstMatch: "",
			},
		}

		sorted := SortResultsByScore(results)

		assert.Equal(expected, sorted, "should be equal")
	})
}
