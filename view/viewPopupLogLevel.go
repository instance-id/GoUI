package view

import (
	. "github.com/instance-id/GoUI/elements"
	ui "github.com/instance-id/clui"
)

//func CreateViewLogLevel() {
//	frmLogLevel := ui.CreateFrame(FrameMenu, 8, 1, ui.BorderNone, ui.Fixed)
//	frmLogLevel.SetGaps(1, ui.KeepValue)
//	frmLogLevel.SetVisible(false)
//	frmLogLevel.SetActive(false)
//}

func SelectLogLevel(btn *ui.ButtonNoShadow) {

	dlgType := ui.SelectDialogRadio

	curr := -1
	for i, lLevel := range Log.LogLevel {
		if lLevel == Log.CurrentLogLevel {
			curr = i
			break
		}
	}

	selDlg := ui.CreateSelectDialog("Choose log level", Log.LogLevel, curr, dlgType)
	selDlg.OnClose(func() {
		switch selDlg.Result() {
		case ui.DialogButton1:
			idx := selDlg.Value()
			if idx != -1 {
				Log.CurrentLogLevel = Log.LogLevel[idx]
			}
		}

		btn.SetEnabled(true)
		// ask the composer to repaint all windows
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	})

}
