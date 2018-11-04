// Copyright 2012-2013 Apcera Inc. All rights reserved.
package termtables

import (
	"strings"
	"testing"
)

func trim(s string) string {
	return strings.TrimPrefix(s, "\n")
}

func checkRendersTo(t *testing.T, table *Table, expected string) {
	output := table.Render()
	if output != expected {
		t.Fatal(actualVersusExpected(output, expected))
	}
}

func actualVersusExpected(actual, expected string) string {
	return "Output didn't match expected\n\n" +
		"Actual:\n\n" +
		actual + "\n" +
		"Expected:\n\n" +
		expected
}
