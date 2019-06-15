package elements

import (
	"fmt"
	. "github.com/instance-id/GoUI/components"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
	term "github.com/nsf/termbox-go"
)

type dbCache struct {
	firstRow int        // previous first visible row
	rowCount int        // previous visible row count
	data     [][]string // cache - contains at least 'rowCount' rows from DB
}

type assetContain struct {
	assetContainer *AssetContainer
}

const columnInTable = 7

func (d *dbCache) preload(firstRow, rowCount int, tmpAssetData *assetContain) {

	if tmpAssetData.assetContainer.AD == nil {
		return
	}

	if firstRow == d.firstRow && rowCount == d.rowCount {
		// fast path: view area is the same, return immediately
		return
	}

	d.data = make([][]string, rowCount, rowCount)
	for i := 0; i < rowCount; i++ {
		absIndex := firstRow + i
		d.data[i] = make([]string, columnInTable, columnInTable)
		d.data[i][0] = tmpAssetData.assetContainer.AD[absIndex].AssetCode
		d.data[i][1] = tmpAssetData.assetContainer.AD[absIndex].AssetName
		d.data[i][2] = tmpAssetData.assetContainer.AD[absIndex].AssetApiKey
		d.data[i][3] = tmpAssetData.assetContainer.AD[absIndex].AssetRole
		d.data[i][4] = tmpAssetData.assetContainer.AD[absIndex].AssetVersion
		d.data[i][5] = tmpAssetData.assetContainer.AD[absIndex].AssetReplaced
		d.data[i][6] = tmpAssetData.assetContainer.AD[absIndex].ReplaceDate
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

func (d *dbCache) AddNewValue(row, col int, newAsset string) string {
	data := []string{newAsset, "", "", "", "", "", ""}
	d.data = append(d.data, data)
	return newAsset
}

// --- Window type for data table ----------------------------------------
func CreateTableDialog(btn *ui.Button) *ui.TableView {
	var tmpAssetData *assetContain
	var restrictor = 0

	tableDialog := new(TableDialog)

	// --- Obtain terminal overall size ----------------------------------
	cw, ch := term.Size()

	// --- Create new popup window for table data ------------------------
	tableDialog.View = ui.AddWindow(cw/2-75, ch/2-16, ui.AutoSize, ui.AutoSize, TxtAssetDetails)
	ui.WindowManager().BeginUpdate()
	defer ui.WindowManager().EndUpdate()
	tableDialog.View.SetGaps(1, ui.KeepValue)
	tableDialog.View.SetModal(true)
	tableDialog.View.SetPack(ui.Vertical)

	tableDialog.Frame = NewFramedWindowInput(tableDialog.View, "", nil)

	// --- Create data table ---------------------------------------------
	td := ui.CreateTableView(tableDialog.Frame, 145, 15, 1)
	ui.ActivateControl(tableDialog.Frame, td)

	// --- Save current values to temp value or create new if nonexistent -----------------
	if Asset == nil {
		tmpAssetData.assetContainer = &AssetContainer{AD: []AssetDetails{{"", "", "", "", "", "", ""}}}
	} else {
		tmpAssetData.assetContainer = Asset
		restrictor = 1
	}

	cache := &dbCache{firstRow: -1}
	var rowCount int
	rowCount = len(tmpAssetData.assetContainer.AD)

	td.SetShowLines(true)
	td.SetShowRowNumber(true)
	td.SetRowCount(rowCount)

	cols := []ui.Column{
		{Title: "Asset Code", Width: 5, Alignment: ui.AlignLeft},
		{Title: "Asset Name", Width: 50, Alignment: ui.AlignLeft},
		{Title: "Asset APIKey", Width: 30, Alignment: ui.AlignLeft},
		{Title: "Asset RoleId", Width: 20, Alignment: ui.AlignLeft},
		{Title: "Version", Width: 7, Alignment: ui.AlignLeft},
		{Title: "Replaced?", Width: 10, Alignment: ui.AlignLeft},
		{Title: "Replace Date", Width: 12, Alignment: ui.AlignLeft},
	}
	td.SetColumns(cols)

	td.OnBeforeDraw(func(col, row, colCnt, rowCnt int) {
		cache.preload(row, rowCnt, tmpAssetData)
		l, t, w, h := td.VisibleArea()
		tableDialog.Frame.SetTitle(fmt.Sprintf("Caching: %d:%d - %dx%d", l, t, w, h))
	})
	td.OnDrawCell(func(info *ui.ColumnDrawInfo) {
		info.Text = cache.value(info.Row, info.Col)
	})

	var newInfo = ui.ColumnDrawInfo{Row: 0, Col: 0}

	td.OnActive(func(active bool) {
		if (func(info *ui.ColumnDrawInfo) string {
			return cache.value(info.Row, info.Col)
		})(&newInfo) == "" {
			if restrictor == 0 {
				cache.CreateNewData(tmpAssetData)
				restrictor = 1
			}
		}
	})

	td.OnAction(func(ev ui.TableEvent) {
		// btns := []string{TxtApplyBtn, TxtCancelBtn}
		//var action string
		switch ev.Action {
		case ui.TableActionSort:
			//action = "Sort table"
		case ui.TableActionEdit:
			c := ev.Col
			r := ev.Row
			var newInfo = ui.ColumnDrawInfo{Row: r, Col: c}
			var editVal = TableEdit{Row: r, Col: c}

			(func(info *ui.ColumnDrawInfo) {
				editVal.OldVal = cache.value(info.Row, info.Col)
			})(&newInfo)

			dlg := CreateEditableDialog(fmt.Sprintf("%s: %s", TxtEditing, editVal.OldVal), editVal.OldVal)
			dlg.View.SetSize(35, 10)
			dlg.View.BaseControl.SetSize(35, 10)
			dlg.OnClose(func() {
				switch dlg.Result() {
				case ui.DialogButton1:
					editVal.NewVal = dlg.EditResult()
					cache.UpdateData(editVal.Row, editVal.Col, editVal.NewVal, tmpAssetData)
					ui.PutEvent(ui.Event{Type: ui.EventRedraw})
				}
			})
			return
		case ui.TableActionNew:
			c := ev.Col
			r := ev.Row
			var editVal = TableEdit{Row: r, Col: c}
			var newInfo = ui.ColumnDrawInfo{Row: r, Col: c}
			dlg := CreateEditableDialog(fmt.Sprintf("%s:", TxtNewAssetCodeValue), "")
			dlg.View.SetSize(35, 10)
			dlg.View.BaseControl.SetSize(35, 10)
			dlg.OnClose(func() {
				switch dlg.Result() {
				case ui.DialogButton1:
					editVal.NewVal = dlg.EditResult()
					details := AssetDetails{AssetCode: editVal.NewVal}
					tmpAssetData.AddNewAsset(details)
					func(info *ui.ColumnDrawInfo) {
						info.Text = cache.AddNewValue(len(tmpAssetData.assetContainer.AD)-1, 0, editVal.NewVal)
					}(&newInfo)
					td.SetRowCount(len(tmpAssetData.assetContainer.AD))
					ui.PutEvent(ui.Event{Type: ui.EventRedraw})

				}
				ui.PutEvent(ui.Event{Type: ui.EventRedraw})

			})
		case ui.TableActionDelete:
			//action = "Delete row"
		default:
			//action = "Unknown action"
		}
	})

	btnFrame := ui.CreateFrame(tableDialog.Frame, 1, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(1, 1)
	textFrame := ui.CreateFrame(btnFrame, 1, 1, ui.BorderNone, ui.Fixed)
	textFrame.SetPack(ui.Vertical)

	ui.CreateLabel(textFrame, ui.AutoSize, ui.AutoSize, TxtInstructs1, ui.Fixed)
	ui.CreateLabel(textFrame, ui.AutoSize, ui.AutoSize, TxtInstructs2, ui.Fixed)
	ui.CreateLabel(textFrame, ui.AutoSize, ui.AutoSize, TxtInstructs3, ui.Fixed)

	// --- Window Controls -----------------------------------------------
	ui.CreateFrame(btnFrame, 1, 1, ui.BorderNone, 1)
	BtnSave := ui.CreateButton(btnFrame, 15, 1, TxtSaveBtn, ui.Fixed)
	BtnSave.SetShadowType(ui.ShadowHalf)
	BtnSave.OnClick(func(ev ui.Event) {
		Asset = tmpAssetData.assetContainer
		btn.SetEnabled(true)
	})
	BtnClose := ui.CreateButton(btnFrame, 15, 1, TxtCloseBtn, ui.Fixed)
	BtnClose.SetShadowType(ui.ShadowHalf)
	BtnClose.OnClick(func(ev ui.Event) {
		ui.WindowManager().DestroyWindow(tableDialog.View)
		btn.SetEnabled(true)
	})

	BtnSave.SetActive(false)
	BtnClose.SetActive(false)
	return td
}

func (d *dbCache) CreateNewData(tmpAssetData *assetContain) *dbCache {
	var editVal = TableEdit{Row: 0, Col: 0}

	dlg := CreateEditableDialog(fmt.Sprintf("%s:", TxtNewAssetCodeValue), "")
	dlg.View.SetSize(35, 10)
	dlg.View.SetActive(true)
	dlg.View.SetEnabled(true)
	dlg.OnClose(func() {
		switch dlg.Result() {
		case ui.DialogButton1:
			editVal.NewVal = dlg.EditResult()
			var newInfo = ui.ColumnDrawInfo{Row: 0, Col: 0}
			data := d.UpdateData(newInfo.Row, newInfo.Col, editVal.NewVal, tmpAssetData)
			tmpAssetData.assetContainer.AD = data
		}
	})
	return d
}

func (a *assetContain) AddNewAsset(asset AssetDetails) []AssetDetails {
	a.assetContainer.AD = append(a.assetContainer.AD, asset)
	return a.assetContainer.AD
}

func (d *dbCache) UpdateData(row int, col int, data string, tmpAssetData *assetContain) []AssetDetails {

	switch col {
	case 0:
		tmpAssetData.assetContainer.AD[row].AssetCode = data
		d.data[row][col] = tmpAssetData.assetContainer.AD[row].AssetCode
	case 1:
		tmpAssetData.assetContainer.AD[row].AssetName = data
		d.data[row][col] = tmpAssetData.assetContainer.AD[row].AssetName
	case 2:
		tmpAssetData.assetContainer.AD[row].AssetApiKey = data
		d.data[row][col] = tmpAssetData.assetContainer.AD[row].AssetApiKey
	case 3:
		tmpAssetData.assetContainer.AD[row].AssetRole = data
		d.data[row][col] = tmpAssetData.assetContainer.AD[row].AssetRole
	case 4:
		tmpAssetData.assetContainer.AD[row].AssetVersion = data
		d.data[row][col] = tmpAssetData.assetContainer.AD[row].AssetVersion
	case 5:
		tmpAssetData.assetContainer.AD[row].AssetReplaced = data
		d.data[row][col] = tmpAssetData.assetContainer.AD[row].AssetReplaced
	case 6:
		tmpAssetData.assetContainer.AD[row].ReplaceDate = data
		d.data[row][col] = tmpAssetData.assetContainer.AD[row].ReplaceDate
	}

	return tmpAssetData.assetContainer.AD
}
