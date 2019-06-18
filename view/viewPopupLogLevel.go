package view

import (
	. "github.com/instance-id/GoUI/components"
	. "github.com/instance-id/GoUI/text"
	ui "github.com/instance-id/clui"
)

func SelectLogLevel(btn *ui.Button, currentLogLevel int) int {

	logLevel := ui.CreateSelectDialog(TxtDbProvider, Log.LogLevel, currentLogLevel, ui.SelectDialogList)
	logLevel.View.SetTitle(Log.LogLevel[currentLogLevel])
	logLevel.OnClose(func() {
		switch logLevel.Result() {
		case ui.DialogButton1:
			logResult := logLevel.Value()
			Log.CurrentLogLevel = logResult
			btn.SetTitle(TxtLogLevelBtn + Log.LogLevel[Log.CurrentLogLevel])
			if Log.CurrentLogLevel != Cntnrs.Dac.System.FileLogLevel {
				pendingMain.logLevel = true
				SavePendingMainSettings()
			} else {
				pendingMain.logLevel = false
				SavePendingMainSettings()
			}
		}
		btn.SetEnabled(true)
	})
	return Log.CurrentLogLevel
}
