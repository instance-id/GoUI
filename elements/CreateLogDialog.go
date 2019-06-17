package elements

import (
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
	term "github.com/nsf/termbox-go"
)

func CreateLogDialog(logTitle string) *LogDialog {
	logDialog := LogViewer
	cw, ch := term.Size()

	logDialog.View = ui.AddWindow(cw/2-75, ch/2-16, ui.AutoSize, ui.AutoSize, logTitle)
	ui.WindowManager().BeginUpdate()
	defer ui.WindowManager().EndUpdate()
	logDialog.View.SetGaps(1, ui.KeepValue)
	logDialog.View.SetModal(true)
	logDialog.View.SetPack(ui.Vertical)

	logDialog.Frame = NewFramedWindowInput(logDialog.View, "", nil)
	logDialog.Frame.SetBackColor(236)
	logDialog.Log = ui.CreateTextView(logDialog.Frame, 145, 25, 1)
	ui.ActivateControl(logDialog.Frame, logDialog.Log)
	autoScroll := ui.CreateCheckBox(logDialog.Frame, ui.AutoSize, TxtAutoScrollChk, ui.Fixed)
	autoScroll.SetState(1)
	autoScroll.OnChange(func(i int) {
		logDialog.Log.SetAutoScroll(func() bool {
			var result bool
			switch i {
			case 0:
				result = false
			case 1:
				result = true
			}
			return result
		}())
	})

	logDialog.Log.SetBackColor(238)
	logDialog.Log.SetTextColor(250)

	//RefreshLog(logDialog)

	// --- Buttons -------------------------------------------------------
	// --- Add new asset to data map ----------------------
	frm1 := ui.CreateFrame(logDialog.View, 16, 4, ui.BorderNone, ui.Fixed)
	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, 1)

	// --- Close the dialog -------------------------------
	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, 1)
	btn3 := ui.CreateButton(frm1, ui.AutoSize, ui.AutoSize, TxtCloseBtn, ui.Fixed)
	btn3.OnClick(func(ev ui.Event) {
		logDialog.result = DialogButton3
		logDialog.edtResult = ""
		logDialog.value = -1
		ui.WindowManager().DestroyWindow(logDialog.View)
		if logDialog.onClose != nil {
			logDialog.onClose()
		}
	})

	logDialog.View.OnClose(func(ev ui.Event) bool {
		if logDialog.result == DialogAlive {
			logDialog.result = DialogClosed
			if ev.X != 1 {
				ui.WindowManager().DestroyWindow(logDialog.View)
			}
			if logDialog.onClose != nil {
				logDialog.onClose()
			}
		}

		return true
	})
	return logDialog
}

// OnClose sets the callback that is called when the
// dialog is closed
func (d *LogDialog) OnClose(fn func()) {
	ui.WindowManager().BeginUpdate()
	defer ui.WindowManager().EndUpdate()
	d.onClose = fn
}

// Result returns what button closed the dialog.
// See DialogButton constants. It can equal DialogAlive
// that means that the dialog is still visible and a
// user still does not click any button
func (d *LogDialog) Result() int {
	return d.result
}

// Value returns the number of the selected item or
// -1 if nothing is selected or the dialog is cancelled
func (d *LogDialog) Value() int {
	return d.value
}

// EditResult returns the text from editfield
func (d *LogDialog) EditResult() string {
	return d.edtResult
}
