package view

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

const columnInTable = 7

func (d *dbCache) preload(firstRow, rowCount int, tmpAssetData AssetContainer) {

	if tmpAssetData.AD == nil {
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
		d.data[i][0] = tmpAssetData.AD[absIndex].AssetCode
		d.data[i][1] = tmpAssetData.AD[absIndex].AssetName
		d.data[i][2] = tmpAssetData.AD[absIndex].AssetApiKey
		d.data[i][3] = tmpAssetData.AD[absIndex].AssetRole
		d.data[i][4] = tmpAssetData.AD[absIndex].AssetReplaced
		d.data[i][5] = tmpAssetData.AD[absIndex].AssetReplacement
		d.data[i][6] = tmpAssetData.AD[absIndex].ReplaceDate
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
	var tmpAssetData AssetContainer
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

	// --- Load table data from data container ---------------------------
	Asset = LoadTableData()

	//fmt.Printf("Test table data: %s", Asset.AD[0].AssetCode)

	//--- Save current values to temp value or create new ---------------
	if Asset.AD == nil {
		tmpAssetData = AssetContainer{AD: []*AssetDetails{{"", "", "", "", "", "", ""}}}
	} else {
		tmpAssetData = Asset
		if tmpAssetData.AD[0].AssetCode == "" {
			restrictor = 0
		} else {
			restrictor = 1
		}
	}

	cache := &dbCache{firstRow: -1}
	var rowCount int
	rowCount = len(tmpAssetData.AD)

	td.SetShowLines(true)
	td.SetShowRowNumber(true)
	td.SetRowCount(rowCount)

	cols := []ui.Column{
		{Title: "Asset Code", Width: 5, Alignment: ui.AlignLeft},
		{Title: "Asset Name", Width: 50, Alignment: ui.AlignLeft},
		{Title: "Asset APIKey", Width: 30, Alignment: ui.AlignLeft},
		{Title: "Asset RoleId", Width: 20, Alignment: ui.AlignLeft},
		{Title: "Replaced?", Width: 10, Alignment: ui.AlignLeft},
		{Title: "Replacement", Width: 5, Alignment: ui.AlignLeft},
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
					tmpAssetData.AddNewAsset(&details)
					func(info *ui.ColumnDrawInfo) {
						info.Text = cache.AddNewValue(len(tmpAssetData.AD)-1, 0, editVal.NewVal)
					}(&newInfo)
					td.SetRowCount(len(tmpAssetData.AD))
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

	// --- Populate container and write config ---------------------------
	BtnSave.OnClick(func(ev ui.Event) {
		Asset = tmpAssetData
		SaveData(Asset)
		_, err := Cntnrs.Wtr.SetConfig()
		if err != nil {
			ui.CreateAlertDialog(ErrCouldNotSaveCfg, fmt.Sprintf("Error Could not save config: %s", err), TxtCloseBtn)
		}

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

func (d *dbCache) CreateNewData(tmpAssetData AssetContainer) *dbCache {
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
			tmpAssetData.AD = data
		}
	})
	return d
}

// --- Update temp data and table display data ---------------------------
func (d *dbCache) UpdateData(row int, col int, data string, tmpAssetData AssetContainer) []*AssetDetails {

	switch col {
	case 0:
		tmpAssetData.AD[row].AssetCode = data
		d.data[row][col] = tmpAssetData.AD[row].AssetCode
	case 1:
		tmpAssetData.AD[row].AssetName = data
		d.data[row][col] = tmpAssetData.AD[row].AssetName
	case 2:
		tmpAssetData.AD[row].AssetApiKey = data
		d.data[row][col] = tmpAssetData.AD[row].AssetApiKey
	case 3:
		tmpAssetData.AD[row].AssetRole = data
		d.data[row][col] = tmpAssetData.AD[row].AssetRole
	case 4:
		tmpAssetData.AD[row].AssetReplaced = data
		d.data[row][col] = tmpAssetData.AD[row].AssetReplaced
	case 5:
		tmpAssetData.AD[row].AssetReplacement = data
		d.data[row][col] = tmpAssetData.AD[row].AssetReplacement
	case 6:
		tmpAssetData.AD[row].ReplaceDate = data
		d.data[row][col] = tmpAssetData.AD[row].ReplaceDate
	}

	return tmpAssetData.AD
}

// --- Copy data back to container to be written to config ---------------
func SaveData(assetData AssetContainer) []*AssetDetails {
	var assetCodes []string
	for d := range assetData.AD {
		if assetData.AD[d].AssetCode != "" {
			assetCodes = append(assetCodes, assetData.AD[d].AssetCode)
			Cntnrs.Dac.Assets.Packages[assetData.AD[d].AssetCode] = assetData.AD[d].AssetName
			Cntnrs.Dac.Assets.ApiKeys[assetData.AD[d].AssetCode] = assetData.AD[d].AssetApiKey
			Cntnrs.Dac.Discord.Roles[assetData.AD[d].AssetCode] = assetData.AD[d].AssetRole
			Cntnrs.Dac.Assets.AssetReplaced[assetData.AD[d].AssetCode] = assetData.AD[d].AssetReplaced
			Cntnrs.Dac.Assets.AssetReplacement[assetData.AD[d].AssetCode] = assetData.AD[d].AssetReplacement
			Cntnrs.Dac.Assets.ReplaceDate[assetData.AD[d].AssetCode] = assetData.AD[d].ReplaceDate
		}

		if _, ok := Cntnrs.Dac.Assets.Packages[""]; ok {
			delete(Cntnrs.Dac.Assets.Packages, "")
		}
		if _, ok := Cntnrs.Dac.Assets.ApiKeys[""]; ok {
			delete(Cntnrs.Dac.Assets.ApiKeys, "")
		}
		if _, ok := Cntnrs.Dac.Discord.Roles[""]; ok {
			delete(Cntnrs.Dac.Discord.Roles, "")
		}
		if _, ok := Cntnrs.Dac.Assets.AssetReplaced[""]; ok {
			delete(Cntnrs.Dac.Assets.AssetReplaced, "")
		}
		if _, ok := Cntnrs.Dac.Assets.AssetReplacement[""]; ok {
			delete(Cntnrs.Dac.Assets.AssetReplacement, "")
		}
		if _, ok := Cntnrs.Dac.Assets.ReplaceDate[""]; ok {
			delete(Cntnrs.Dac.Assets.ReplaceDate, "")
		}

	}
	Cntnrs.Dac.Assets.AssetCodes = assetCodes
	return assetData.AD
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// ---  ------------------------------------------------
