package view

import (
	. "github.com/instance-id/GoUI/elements"
	ui "github.com/instance-id/clui"
)

func CreateViewPopupTheme() {
	frmTheme := ui.CreateFrame(FrameMenu, 8, 1, ui.BorderNone, ui.Fixed)
	frmTheme.SetGaps(1, ui.KeepValue)
	frmTheme.SetVisible(false)
	frmTheme.SetActive(false)
}

func ChangeTheme(btn *ui.ButtonNoShadow) {
	items := ui.ThemeNames()
	dlgType := ui.SelectDialogRadio

	curr := -1
	for i, tName := range items {
		if tName == ui.CurrentTheme() {
			curr = i
			break
		}
	}

	selDlg := ui.CreateSelectDialog("Choose a theme", items, curr, dlgType)
	selDlg.OnClose(func() {
		switch selDlg.Result() {
		case ui.DialogButton1:
			idx := selDlg.Value()
			if idx != -1 {
				ui.SetCurrentTheme(items[idx])
			}
		}

		btn.SetEnabled(true)
		// ask the composer to repaint all windows
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	})
}
