// Copyright 2013 Apcera Inc. All rights reserved.

package termtables

import (
	"testing"
)

func TestCreateTableHTML(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<tr><th>Name</th><th>Value</th></tr>
</thead>
<tbody>
<tr><td>hey</td><td>you</td></tr>
<tr><td>ken</td><td>1234</td></tr>
<tr><td>derek</td><td>3.14</td></tr>
<tr><td>derek too</td><td>3.15</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()

	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableWithHeaderHTML(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<caption>Example</caption>
<tr><th>Name</th><th>Value</th></tr>
</thead>
<tbody>
<tr><td>hey</td><td>you</td></tr>
<tr><td>ken</td><td>1234</td></tr>
<tr><td>derek</td><td>3.14</td></tr>
<tr><td>derek too</td><td>3.15</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()

	table.AddTitle("Example")
	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableTitleWidthAdjustsHTML(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<caption>Example My Foo Bar&#39;d Test</caption>
<tr><th>Name</th><th>Value</th></tr>
</thead>
<tbody>
<tr><td>hey</td><td>you</td></tr>
<tr><td>ken</td><td>1234</td></tr>
<tr><td>derek</td><td>3.14</td></tr>
<tr><td>derek too</td><td>3.15</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()

	table.AddTitle("Example My Foo Bar'd Test")
	table.AddHeaderRow("Name", "Value")
	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableWithNoHeadersHTML(t *testing.T) {
	expected := trim(`
<table class="termtable">
<tbody>
<tr><td>hey</td><td>you</td></tr>
<tr><td>ken</td><td>1234</td></tr>
<tr><td>derek</td><td>3.14</td></tr>
<tr><td>derek too</td><td>3.15</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()

	table.AddRow("hey", "you")
	table.AddRow("ken", 1234)
	table.AddRow("derek", 3.14)
	table.AddRow("derek too", 3.1456788)

	checkRendersTo(t, table, expected)
}

func TestTableUnicodeWidthsHTML(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<tr><th>Name</th><th>Cost</th></tr>
</thead>
<tbody>
<tr><td>Currency</td><td>¤10</td></tr>
<tr><td>US Dollar</td><td>$30</td></tr>
<tr><td>Euro</td><td>€27</td></tr>
<tr><td>Thai</td><td>฿70</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()
	table.AddHeaderRow("Name", "Cost")
	table.AddRow("Currency", "¤10")
	table.AddRow("US Dollar", "$30")
	table.AddRow("Euro", "€27")
	table.AddRow("Thai", "฿70")

	checkRendersTo(t, table, expected)
}

func TestTableWithAlignment(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<tr><th>Foo</th><th>Bar</th></tr>
</thead>
<tbody>
<tr><td>humpty</td><td>dumpty</td></tr>
<tr><td align='right'>r</td><td>&lt;- on right</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()
	table.AddHeaderRow("Foo", "Bar")
	table.AddRow("humpty", "dumpty")
	table.AddRow(CreateCell("r", &CellStyle{Alignment: AlignRight}), "<- on right")

	checkRendersTo(t, table, expected)
}

func TestTableAfterSetAlign(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<tr><th>Alphabetical</th><th>Num</th></tr>
</thead>
<tbody>
<tr><td align='right'>alfa</td><td>1</td></tr>
<tr><td align='right'>bravo</td><td>2</td></tr>
<tr><td align='right'>charlie</td><td>3</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()
	table.AddHeaderRow("Alphabetical", "Num")
	table.AddRow("alfa", 1)
	table.AddRow("bravo", 2)
	table.AddRow("charlie", 3)
	table.SetAlign(AlignRight, 1)

	checkRendersTo(t, table, expected)
}

func TestTableWithAltTitleStyle(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<tr><th style="text-align: center" colspan="3">Metasyntactic</th></tr>
<tr><th>Foo</th><th>Bar</th><th>Baz</th></tr>
</thead>
<tbody>
<tr><td>a</td><td>b</td><td>c</td></tr>
<tr><td>α</td><td>β</td><td>γ</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()
	table.SetHTMLStyleTitle(TitleAsThSpan)
	table.AddTitle("Metasyntactic")
	table.AddHeaderRow("Foo", "Bar", "Baz")
	table.AddRow("a", "b", "c")
	table.AddRow("α", "β", "γ")

	checkRendersTo(t, table, expected)
}

func TestTableWithMultipleHeaderRows(t *testing.T) {
	expected := trim(`
<table class="termtable">
<thead>
<tr><th style="text-align: center" colspan="2">Metasyntactic</th></tr>
<tr><th>Foo</th><th>Bar</th></tr>
<tr><th>Baz</th><th>Quux</th></tr>
</thead>
<tbody>
<tr><td>a</td><td>b</td></tr>
<tr><td>α</td><td>β</td></tr>
</tbody>
</table>
`)

	table := CreateTable()
	table.SetModeHTML()
	table.SetHTMLStyleTitle(TitleAsThSpan)
	table.AddTitle("Metasyntactic")
	table.AddHeaderRow("Foo", "Bar")
	table.AddHeaderRow("Baz", "Quux")
	table.AddRow("a", "b")
	table.AddRow("α", "β")

	checkRendersTo(t, table, expected)
}
