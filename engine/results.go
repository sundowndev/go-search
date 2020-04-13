package engine

import (
	"fmt"
	"os"
	"text/tabwriter"
)

const tabPadding = 3
const tabPaddingChar = ' '

// QueryResult defines the structure of a query result
type QueryResult struct {
	File       string
	Count      float64
	FirstMatch string
}

// ShowResults display results in a table
func ShowResults(results []*QueryResult) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, tabPadding, tabPaddingChar, tabwriter.AlignRight|tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintln(w, "File\tCount\tFirst match\t")
	for _, v := range results {
		fmt.Fprintf(w, "%v\t%v\t%v\t\n", v.File, v.Count, v.FirstMatch)
	}
}
