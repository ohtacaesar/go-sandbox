package excelize

import "testing"

func TestData(t *testing.T) {
	f := openTestDataAndSetCleanup(t)
	sheetName := f.GetSheetList()[0]
	t.Log(f.GetCellFormula(sheetName, "A1"))
}
