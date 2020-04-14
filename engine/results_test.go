package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResults(t *testing.T) {
	assert := assert.New(t)

	t.Run("SortResultsByScore", func(t *testing.T) {
		results := []*QueryResult{
			&QueryResult{
				File:       "/file",
				Score:      1,
				FirstMatch: "",
			},
			&QueryResult{
				File:       "/file",
				Score:      3,
				FirstMatch: "",
			},
			&QueryResult{
				File:       "/file",
				Score:      9,
				FirstMatch: "",
			},
			&QueryResult{
				File:       "/file",
				Score:      2,
				FirstMatch: "",
			},
		}

		expected := []*QueryResult{
			&QueryResult{
				File:       "/file",
				Score:      9,
				FirstMatch: "",
			},
			&QueryResult{
				File:       "/file",
				Score:      3,
				FirstMatch: "",
			},
			&QueryResult{
				File:       "/file",
				Score:      2,
				FirstMatch: "",
			},
			&QueryResult{
				File:       "/file",
				Score:      1,
				FirstMatch: "",
			},
		}

		sorted := SortResultsByScore(results)

		assert.Equal(expected, sorted, "should be equal")
	})
}
