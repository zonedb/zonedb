package build

import (
	"testing"
)

func TestAppendRune(t *testing.T) {
	table := &CodeTable{CodeRange{'0', '9'}}

	table.AppendRune('-')
	CompareTable(t, *table, "09--", false)

	table.AppendRune(',')
	CompareTable(t, *table, "09--,,", false)

	table.AppendRune('_')
	CompareTable(t, *table, "09--,,__", false)
}

func TestAppendRuneCompress(t *testing.T) {
	table := &CodeTable{CodeRange{'0', '9'}}

	table.AppendRune('-')
	CompareTable(t, *table, "--09", true)

	table.AppendRune(',')
	CompareTable(t, *table, ",-09", true)

	table.AppendRune('_')
	CompareTable(t, *table, ",-09__", true)
}

func TestAppendRange(t *testing.T) {
	table := &CodeTable{}

	table.AppendRange('0', '9')
	CompareTable(t, *table, "09", false)

	// Range being added is inside of existing range
	table.AppendRange('1', '2')
	CompareTable(t, *table, "0912", false)

	// Starts before beginning, finishes in middle
	table.AppendRange('/', '1')
	CompareTable(t, *table, "0912/1", false)

	// Starts in middle, finishes after end
	table.AppendRange('8', ';')
	CompareTable(t, *table, "0912/18;", false)

	// Range being added completely
	table.AppendRange(' ', '@')
	CompareTable(t, *table, "0912/18; @", false)
}

func TestAppendRangeCompress(t *testing.T) {
	table := &CodeTable{}

	table.AppendRange('0', '9')
	CompareTable(t, *table, "09", false)
	CompareTable(t, *table, "09", true)

	// Range being added is inside of existing range
	table.AppendRange('1', '2')
	CompareTable(t, *table, "0912", false)
	CompareTable(t, *table, "09", true)

	// Starts before beginning, finishes in middle
	table.AppendRange('/', '1')
	CompareTable(t, *table, "0912/1", false)
	CompareTable(t, *table, "/9", true)

	// Starts in middle, finishes after end
	table.AppendRange('8', ';')
	CompareTable(t, *table, "0912/18;", false)
	CompareTable(t, *table, "/;", true)

	// Range being added completely
	table.AppendRange(' ', '@')
	CompareTable(t, *table, "0912/18; @", false)
	CompareTable(t, *table, " @", true)
}

func CompareTable(t *testing.T, table CodeTable, expected string, compress bool) {
	var tmp CodeTable
	if compress {
		tmp = make(CodeTable, len(table))
		copy(tmp, table)
		tmp.Compress()
	} else {
		tmp = table
	}
	out, _ := tmp.MarshalText()
	outStr := string(out)
	if outStr != expected {
		t.Errorf(`Expected %q, got %q`, expected, outStr)
	}
}
