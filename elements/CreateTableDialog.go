package elements

import (
	"fmt"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
	term "github.com/nsf/termbox-go"
)

var tmpAssetData []*AssetDetails

type dbCache struct {
	firstRow int        // previous first visible row
	rowCount int        // previous visible row count
	data     [][]string // cache - contains at least 'rowCount' rows from DB
}

const columnInTable = 7

func (d *dbCache) preload(firstRow, rowCount int) {
	if firstRow == d.firstRow && rowCount == d.rowCount {
		// fast path: view area is the same, return immediately
		return
	}

	d.data = make([][]string, rowCount, rowCount)
	for i := 0; i < rowCount; i++ {
		absIndex := firstRow + i
		d.data[i] = make([]string, columnInTable, columnInTable)
		d.data[i][0] = tmpAssetData[absIndex].AssetCode
		d.data[i][1] = tmpAssetData[absIndex].AssetName
		d.data[i][2] = tmpAssetData[absIndex].AssetApiKey
		d.data[i][3] = tmpAssetData[absIndex].AssetRole
		d.data[i][4] = tmpAssetData[absIndex].AssetVersion
		d.data[i][5] = tmpAssetData[absIndex].AssetReplaced
		d.data[i][6] = tmpAssetData[absIndex].ReplaceDate
	}

	// do not forget to save the last values
	d.firstRow = firstRow
	d.rowCount = rowCount
}

func (d *dbCache) value(row, col int) string {
	rowId := row - d.firstRow
	if rowId >= len(d.data) {
		return ""
	}
	rowValues := d.data[rowId]
	if col >= len(rowValues) {
		return ""
	}
	return rowValues[col]
}

func (d *dbCache) NewValue(row, col int, newText string) {
	d.data[row][col] = newText
}

// --- Window type for data table ----------------------------------------
func CreateTableDialog(btn *ui.ButtonNoShadow, tableTitle string) {
	tableDialog := new(TableDialog)

	// --- Obtain terminal overall size ----------------------------------
	cw, ch := term.Size()

	// --- Save current values to temp value until saved -----------------
	tmpAssetData = AssetDetail

	// --- Create new popup window for table data ------------------------
	tableDialog.View = ui.AddWindow(cw/2-75, ch/2-16, ui.AutoSize, ui.AutoSize, TxtAssetDetails)
	ui.WindowManager().BeginUpdate()
	defer ui.WindowManager().EndUpdate()
	tableDialog.View.SetGaps(1, ui.KeepValue)
	tableDialog.View.SetModal(true)
	tableDialog.View.SetPack(ui.Vertical)

	tableDialog.Frame = NewFramedWindowInput(tableDialog.View, tableTitle, nil)

	// --- Create data table ---------------------------------------------
	td := ui.CreateTableView(tableDialog.Frame, 145, 15, 1)
	ui.ActivateControl(tableDialog.Frame, td)

	cache := &dbCache{firstRow: -1}
	rowCount := len(AssetDetail)
	td.SetShowLines(true)
	td.SetShowRowNumber(true)
	td.SetRowCount(rowCount)

	cols := []ui.Column{
		ui.Column{Title: "Asset Code", Width: 5, Alignment: ui.AlignLeft},
		ui.Column{Title: "Asset Name", Width: 50, Alignment: ui.AlignLeft},
		ui.Column{Title: "Asset APIKey", Width: 30, Alignment: ui.AlignLeft},
		ui.Column{Title: "Asset RoleId", Width: 20, Alignment: ui.AlignLeft},
		ui.Column{Title: "Version", Width: 7, Alignment: ui.AlignLeft},
		ui.Column{Title: "Replaced?", Width: 10, Alignment: ui.AlignLeft},
		ui.Column{Title: "Replace Date", Width: 12, Alignment: ui.AlignLeft},
	}
	td.SetColumns(cols)

	td.OnBeforeDraw(func(col, row, colCnt, rowCnt int) {
		cache.preload(row, rowCnt)
		l, t, w, h := td.VisibleArea()
		tableDialog.Frame.SetTitle(fmt.Sprintf("Caching: %d:%d - %dx%d", l, t, w, h))
	})
	td.OnDrawCell(func(info *ui.ColumnDrawInfo) {
		info.Text = cache.value(info.Row, info.Col)
	})

	td.OnAction(func(ev ui.TableEvent) {
		btns := []string{"Close", "Dismiss"}
		var action string
		switch ev.Action {
		case ui.TableActionSort:
			action = "Sort table"
		case ui.TableActionEdit:
			c := ev.Col
			r := ev.Row
			var newInfo = ui.ColumnDrawInfo{Row: r, Col: c}
			var editVal = TableEdit{Row: r, Col: c}

			(func(info *ui.ColumnDrawInfo) {
				editVal.OldVal = cache.value(info.Row, info.Col)
			})(&newInfo)

			dlg := ui.CreateEditDialog(fmt.Sprintf("Editing value: %s", editVal.OldVal), "New value", editVal.OldVal)
			dlg.OnClose(func() {
				switch dlg.Result() {
				case ui.DialogButton1:
					editVal.NewVal = dlg.EditResult()
					cache.NewValue(editVal.Row, editVal.Col, editVal.NewVal)
					ui.PutEvent(ui.Event{Type: ui.EventRedraw})
				}
			})
			return
		case ui.TableActionNew:
			action = "Add new row"
		case ui.TableActionDelete:
			action = "Delete row"
		default:
			action = "Unknown action"
		}

		dlg := ui.CreateConfirmationDialog(
			"<c:blue>"+action,
			"Click any button or press <c:yellow>SPACE<c:> to close the dialog",
			btns, ui.DialogButton1)
		dlg.OnClose(func() {})
	})

	btnFrame := ui.CreateFrame(tableDialog.Frame, 1, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(1, 1)
	textFrame := ui.CreateFrame(btnFrame, 1, 1, ui.BorderNone, ui.Fixed)
	textFrame.SetPack(ui.Vertical)
	ui.CreateLabel(textFrame, ui.AutoSize, ui.AutoSize, "_____________", ui.Fixed)
	ui.CreateLabel(textFrame, ui.AutoSize, ui.AutoSize, "Instructions: Use arrow keys or pageup/down to navigate the fields.", ui.Fixed)
	ui.CreateLabel(textFrame, ui.AutoSize, ui.AutoSize, "Highlight the field you would like to edit. Press F2 or Space to edit the field ", ui.Fixed)
	ui.CreateLabel(textFrame, ui.AutoSize, ui.AutoSize, "Simply press ok when completed. - Don't forget to save! -", ui.Fixed)

	// --- Window Controls -----------------------------------------------
	ui.CreateFrame(btnFrame, 1, 1, ui.BorderNone, 1)
	BtnSave := ui.CreateButton(btnFrame, 15, 1, " Save ", ui.Fixed)
	BtnSave.OnClick(func(ev ui.Event) {
		AssetDetail = tmpAssetData
		btn.SetEnabled(true)
	})
	BtnClose := ui.CreateButton(btnFrame, 15, 1, " Close ", ui.Fixed)
	BtnClose.OnClick(func(ev ui.Event) {
		ui.WindowManager().DestroyWindow(tableDialog.View)
		btn.SetEnabled(true)
	})
}
