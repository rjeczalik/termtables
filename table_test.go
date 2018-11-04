// Copyright 2012-2013 Apcera Inc. All rights reserved.
package termtables

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	expected := trim(`
+-----------+-------+
| Name      | Value |
+-----------+-------+
| hey       | you   |
| ken       | 1234  |
| derek     | 3.14  |
| derek too | 3.15  |
| escaping  | rox%% |
+-----------+-------+
`)

	table := CreateTable()

	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)
	table.AddRow("escaping", "rox%%")

	checkRendersTo(t, table, expected)
}

func TestStyleResets(t *testing.T) {
	expected := trim(`
+-----------+-------+
| Name      | Value |
+-----------+-------+
| hey       | you   |
| ken       | 1234  |
| derek     | 3.14  |
| derek too | 3.15  |
+-----------+-------+
`)

	table := CreateTable()
	table.UTF8Box()
	table.Style.setAsciiBoxStyle()

	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableWithHeader(t *testing.T) {
	expected := trim(`
+-------------------+
|      Example      |
+-----------+-------+
| Name      | Value |
+-----------+-------+
| hey       | you   |
| ken       | 1234  |
| derek     | 3.14  |
| derek too | 3.15  |
+-----------+-------+
`)

	table := CreateTable()

	table.AddTitle("Example")
	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

// TestTableWithHeaderMultipleTimes ensures that printing a table with headers
// multiple times continues to render correctly.
func TestTableWithHeaderMultipleTimes(t *testing.T) {
	expected := trim(`
+-------------------+
|      Example      |
+-----------+-------+
| Name      | Value |
+-----------+-------+
| hey       | you   |
| ken       | 1234  |
| derek     | 3.14  |
| derek too | 3.15  |
+-----------+-------+
`)

	table := CreateTable()

	table.AddTitle("Example")
	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
	checkRendersTo(t, table, expected)
}

func TestTableTitleWidthAdjusts(t *testing.T) {
	expected := trim(`
+---------------------------+
| Example My Foo Bar'd Test |
+-----------+---------------+
| Name      | Value         |
+-----------+---------------+
| hey       | you           |
| ken       | 1234          |
| derek     | 3.14          |
| derek too | 3.15          |
+-----------+---------------+
`)

	table := CreateTable()

	table.AddTitle("Example My Foo Bar'd Test")
	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableHeaderWidthAdjusts(t *testing.T) {
	expected := trim(`
+---------------+---------------------+
| Slightly Long | More than 2 columns |
+---------------+---------------------+
| a             | b                   |
+---------------+---------------------+
`)

	table := CreateTable()

	table.AddHeaderRow("Slightly Long", "More than 2 columns")
	table.AddRow("a", "b")

	checkRendersTo(t, table, expected)
}

func TestTableWithNoHeaders(t *testing.T) {
	expected := trim(`
+-----------+------+
| hey       | you  |
| ken       | 1234 |
| derek     | 3.14 |
| derek too | 3.15 |
+-----------+------+
`)

	table := CreateTable()

	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableUnicodeWidths(t *testing.T) {
	expected := trim(`
+-----------+------+
| Name      | Cost |
+-----------+------+
| Currency  | ¤10  |
| US Dollar | $30  |
| Euro      | €27  |
| Thai      | ฿70  |
+-----------+------+
`)

	table := CreateTable()
	table.AddHeaderRow("Name", "Cost")
	table.AddRow("Currency", "¤10")
	table.AddRow("US Dollar", "$30")
	table.AddRow("Euro", "€27")
	table.AddRow("Thai", "฿70")

	checkRendersTo(t, table, expected)
}

func TestTableInUTF8(t *testing.T) {
	expected := trim(`
╭───────────────────╮
│      Example      │
├───────────┬───────┤
│ Name      │ Value │
├───────────┼───────┤
│ hey       │ you   │
│ ken       │ 1234  │
│ derek     │ 3.14  │
│ derek too │ 3.15  │
│ escaping  │ rox%% │
╰───────────┴───────╯
`)

	table := CreateTable()
	table.UTF8Box()

	table.AddTitle("Example")
	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)
	table.AddRow("escaping", "rox%%")

	checkRendersTo(t, table, expected)
}

func TestTableInUTF8WithMultipleHeaderRows(t *testing.T) {
	expected := trim(`
╭───────────────────╮
│      Example      │
├───────────┬───────┤
│ Name      │ Value │
│ Sub       │ Sub   │
├───────────┼───────┤
│ hey       │ you   │
│ ken       │ 1234  │
│ derek     │ 3.14  │
│ derek too │ 3.15  │
│ escaping  │ rox%% │
╰───────────┴───────╯
`)

	table := CreateTable()
	table.UTF8Box()

	table.AddTitle("Example")
	table.AddHeaderRow("Name", "Value")
	table.AddHeaderRow("Sub", "Sub")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)
	table.AddRow("escaping", "rox%%")

	checkRendersTo(t, table, expected)
}

func TestTableInUTF8WithColspan(t *testing.T) {
	expected := trim(`
╭──────────────╮
│   Example    │
├──────┬───────┤
│ Name │ Value │
│ Sub          │
├──────┼───────┤
│ hey  │ you   │
╰──────┴───────╯
`)

	table := CreateTable()
	table.UTF8Box()

	table.AddTitle("Example")
	table.AddHeaderRow("Name", "Value")
	table.AddHeaderRow(CreateCell("Sub", &CellStyle{ColSpan: 2}))
	table.AddRow("hey", "you")

	checkRendersTo(t, table, expected)
}

func TestTableUnicodeUTF8AndSGR(t *testing.T) {
	// at present, this mostly just tests that alignment still works
	expected := "" +
		"╭───────────────────────╮\n" +
		"│       \033[1mFanciness\033[0m       │\n" +
		"├──────────┬────────────┤\n" +
		"│ \033[31mred\033[0m      │ \033[32mgreen\033[0m      │\n" +
		"├──────────┼────────────┤\n" +
		"│ plain    │ text       │\n" +
		"│ Καλημέρα │ κόσμε      │\n" +
		"│ \033[1mvery\033[0m     │ \033[4munderlined\033[0m │\n" +
		"│ a\033[1mb\033[0mc      │ \033[45mmagenta\033[0m    │\n" +
		"│ \033[31m→\033[0m        │ \033[32m←\033[0m          │\n" +
		"╰──────────┴────────────╯\n"

	sgred := func(in string, sgrPm string) string {
		return "\033[" + sgrPm + "m" + in + "\033[0m"
	}
	bold := func(in string) string { return sgred(in, "1") }

	table := CreateTable()
	table.UTF8Box()

	table.AddTitle(bold("Fanciness"))
	table.AddHeaderRow(sgred("red", "31"), sgred("green", "32"))
	table.AddRow("plain", "text")
	table.AddRow("Καλημέρα", "κόσμε") // from http://plan9.bell-labs.com/sys/doc/utf.html
	table.AddRow(bold("very"), sgred("underlined", "4"))
	table.AddRow("a"+bold("b")+"c", sgred("magenta", "45"))
	table.AddRow(sgred("→", "31"), sgred("←", "32"))
	// TODO: in future, if we start detecting presence of SGR sequences, we
	// should ensure that the SGR reset is done at the end of the cell content,
	// so that SGR doesn't "bleed across" (cells or rows).  We would then add
	// tests for that here.
	//
	// Of course, at that point, we'd also want to support automatic HTML
	// styling conversion too, so would need a test for that also.
	checkRendersTo(t, table, expected)
}

func TestTableInMarkdown(t *testing.T) {
	expected := trim(`
Table: Example

| Name  | Value |
| ----- | ----- |
| hey   | you   |
| a &#x7c; b | esc   |
| esc   | rox%% |
`)

	table := CreateTable()
	table.SetModeMarkdown()

	table.AddTitle("Example")
	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("a | b", "esc")
	table.AddRow("esc", "rox%%")

	checkRendersTo(t, table, expected)
}

func TestTitleUnicodeWidths(t *testing.T) {
	expected := trim(`
+-------+
| ← 5 → |
+---+---+
| a | b |
| c | d |
| e | 3 |
+---+---+
`)

	// minimum width for a table of two columns is 9 characters, given
	// one space of padding, and non-empty tables.

	table := CreateTable()

	// We have 4 characters down for left and right columns and padding, so
	// a width of 5 for us should match the minimum per the columns

	// 5 characters; each arrow is three octets in UTF-8, giving 9 bytes
	// so, same in character-count-width, longer in bytes
	table.AddTitle("← 5 →")

	// a single character per cell, here; use ASCII characters
	table.AddRow("a", "b")
	table.AddRow("c", "d")
	table.AddRow("e", 3)

	checkRendersTo(t, table, expected)
}

// We identified two error conditions wherein length wrapping would not correctly
// wrap width when, for instance, in a two-column table, the longest row in the
// right-hand column was not the same as the longest row in the left-hand column.
// This tests that we correctly accumulate the maximum width across all rows of
// the termtable and adjust width accordingly.
func TestTableWidthHandling(t *testing.T) {
	expected := trim(`
+-----------------------------------------+
|        Example... to Fix My Test        |
+-----------------+-----------------------+
| hey foo bar baz | you                   |
| ken             | you should write code |
| derek           | 3.14                  |
| derek too       | 3.15                  |
+-----------------+-----------------------+
`)

	table := CreateTable()

	table.AddTitle("Example... to Fix My Test")
	table.AddRow("hey foo bar baz", "you")
	table.AddRow("ken", "you should write code")
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableWidthHandling_SecondErrorCondition(t *testing.T) {
	expected := trim(`
+----------------------------------------+
|       Example... to Fix My Test        |
+-----------------+----------------------+
| hey foo bar baz | you                  |
| ken             | you should sell cod! |
| derek           | 3.14                 |
| derek too       | 3.15                 |
+-----------------+----------------------+
`)

	table := CreateTable()

	table.AddTitle("Example... to Fix My Test")
	table.AddRow("hey foo bar baz", "you")
	table.AddRow("ken", "you should sell cod!")
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableAlignPostsetting(t *testing.T) {
	expected := trim(`
+-----------+-------+
| Name      | Value |
+-----------+-------+
|       hey | you   |
|       ken | 1234  |
|     derek | 3.14  |
| derek too | 3.15  |
|  escaping | rox%% |
+-----------+-------+
`)

	table := CreateTable()

	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)
	table.AddRow("escaping", "rox%%")

	table.SetAlign(AlignRight, 1)

	checkRendersTo(t, table, expected)
}

func TestTableMissingCells(t *testing.T) {
	expected := trim(`
+----------+---------+---------+
| Name     | Value 1 | Value 2 |
+----------+---------+---------+
| hey      | you     | person  |
| ken      | 1234    |
| escaping | rox%s%% |
+----------+---------+---------+
`)
	// FIXME: missing extra cells there

	table := CreateTable()

	table.AddHeaderRow("Name", "Value 1", "Value 2")
	table.AddRow("hey", "you", "person")
	table.AddRow("ken", 1234)
	table.AddRow("escaping", "rox%s%%")

	checkRendersTo(t, table, expected)
}

// We don't yet support combining characters, double-width characters or
// anything to do with estimating a tty-style "character width" for what in
// Unicode is a grapheme cluster.  This disabled test shows what we want
// to support, but don't yet.
func TestTableWithCombiningChars(t *testing.T) {
	expected := trim(`
+------+---+
| noel | 1 |
| noël | 2 |
| noël | 3 |
+------+---+
`)

	table := CreateTable()

	table.AddRow("noel", "1")
	table.AddRow("noe\u0308l", "2") // LATIN SMALL LETTER E  +  COMBINING DIAERESIS
	table.AddRow("noël", "3")       // Hex EB; LATIN SMALL LETTER E WITH DIAERESIS

	checkRendersTo(t, table, expected)
}

// another unicode length issue
func TestTableWithFullwidthChars(t *testing.T) {
	expected := trim(`
+----------+------------+
| wide     | not really |
| ｗｉｄｅ | fullwidth  |
+----------+------------+
`)

	table := CreateTable()
	table.AddRow("wide", "not really")
	table.AddRow("ｗｉｄｅ", "fullwidth") // FULLWIDTH LATIN SMALL LETTER <X>

	checkRendersTo(t, table, expected)
}

// Tests CJK characters using examples given in issue #33. The examples may not
// look like they line up but you can visually confirm its accuracy with a
// fmt.Print.
func TestCJKChars(t *testing.T) {
	expected := trim(`
+-------+---------+----------+
| KeyID | ValueID | ValueCN  |
+-------+---------+----------+
| 8     | 51      | 精钢     |
| 8     | 52      | 鳄鱼皮   |
| 8     | 53      | 镀金皮带 |
| 8     | 54      | 精钢     |
+-------+---------+----------+
`)

	table := CreateTable()
	table.AddHeaderRow("KeyID", "ValueID", "ValueCN")
	table.AddRow("8", 51, "精钢")
	table.AddRow("8", 52, "鳄鱼皮")
	table.AddRow("8", 53, "镀金皮带")
	table.AddRow("8", 54, "精钢")
	checkRendersTo(t, table, expected)

	expected2 := trim(`
+--------------------+----------------------+
| field              | value                |
+--------------------+----------------------+
| GoodsPropertyKeyID | 9                    |
| MerchantAccountID  | 0                    |
| GoodsCategoryCode  | 100001               |
| NameCN             | 机芯类型             |
| NameJP             | ムーブメントのタイプ |
+--------------------+----------------------+
`)
	table = CreateTable()
	table.AddHeaderRow("field", "value")
	table.AddRow("GoodsPropertyKeyID", 9)
	table.AddRow("MerchantAccountID", 0)
	table.AddRow("GoodsCategoryCode", 100001)
	table.AddRow("NameCN", "机芯类型")
	table.AddRow("NameJP", "ムーブメントのタイプ")
	checkRendersTo(t, table, expected2)
}

func TestTableMultipleAddHeaderRow(t *testing.T) {
	expected := trim(`
+--------------+--------+
| First column | Second |
| Sub          | Sub    |
+--------------+--------+
| 2            | 3      |
+--------------+--------+
`)

	table := CreateTable()
	table.AddHeaderRow("First column", "Second")
	table.AddHeaderRow("Sub", "Sub")
	table.AddRow(2, 3)

	checkRendersTo(t, table, expected)
}

func createTestTable() *Table {
	table := CreateTable()
	header := []interface{}{}
	for i := 0; i < 50; i++ {
		header = append(header, "First Column")
	}
	table.AddHeaderRow(header...)
	for i := 0; i < 3000; i++ {
		row := []interface{}{}
		for i := 0; i < 50; i++ {
			row = append(row, "First row value")
		}
		table.AddRow(row...)
	}
	return table
}

func BenchmarkTableRenderTerminal(b *testing.B) {
	table := createTestTable()
	table.SetModeTerminal()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table.Render()
	}
}

func BenchmarkTableRenderMarkdown(b *testing.B) {
	table := createTestTable()
	table.SetModeMarkdown()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table.Render()
	}
}

func BenchmarkTableRenderHTML(b *testing.B) {
	table := createTestTable()
	table.SetModeHTML()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table.Render()
	}
}
