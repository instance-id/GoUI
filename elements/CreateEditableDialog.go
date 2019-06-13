package elements

import (
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
	term "github.com/nsf/termbox-go"
)

func CreateEditableDialog(editTitle string, oldVal string) *EditableDialog {
	editableDialog := new(EditableDialog)

	cw, ch := term.Size()

	editableDialog.View = ui.AddWindow(cw/2-28, ch/2-8, ui.AutoSize, ui.AutoSize, TxtAssetCodes)
	ui.WindowManager().BeginUpdate()
	defer ui.WindowManager().EndUpdate()
	editableDialog.View.SetGaps(1, ui.KeepValue)
	editableDialog.View.SetModal(true)
	editableDialog.View.SetPack(ui.Vertical)

	var params = FramedInputParams{Orientation: ui.Vertical, PadX: 2, PadY: 2}
	editableDialog.Frame = NewFramedWindowInput(editableDialog.View, editTitle, &params)
	editableDialog.edit = ui.CreateEditField(editableDialog.Frame, 50, oldVal, ui.Fixed)
	ui.ActivateControl(editableDialog.Frame, editableDialog.edit)

	// --- Buttons -------------------------------------------------------
	// --- Edit currently selected entry ------------------
	frm1 := ui.CreateFrame(editableDialog.View, 16, 4, ui.BorderNone, ui.Fixed)
	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, ui.Fixed)
	btn0 := ui.CreateButton(frm1, ui.AutoSize, ui.AutoSize, TxtApplyBtn, ui.Fixed)
	btn0.OnClick(func(ev ui.Event) {
		editableDialog.result = DialogButton1
		editableDialog.edtResult = editableDialog.edit.Title()
		ui.WindowManager().DestroyWindow(editableDialog.View)

		if editableDialog.onClose != nil {
			editableDialog.onClose()
		}
	})

	// --- Close the dialog -------------------------------
	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, ui.Fixed)
	btn1 := ui.CreateButton(frm1, ui.AutoSize, ui.AutoSize, TxtCancelBtn, ui.Fixed)
	btn1.OnClick(func(ev ui.Event) {
		editableDialog.result = DialogButton2
		editableDialog.edtResult = oldVal
		editableDialog.value = -1
		ui.WindowManager().DestroyWindow(editableDialog.View)
		if editableDialog.onClose != nil {
			editableDialog.onClose()
		}
	})

	editableDialog.View.OnClose(func(ev ui.Event) bool {
		if editableDialog.result == DialogAlive {
			editableDialog.result = DialogClosed
			if ev.X != 1 {
				ui.WindowManager().DestroyWindow(editableDialog.View)
			}
			if editableDialog.onClose != nil {
				editableDialog.onClose()
			}
		}

		return true
	})

	btn0.SetActive(false)
	btn1.SetActive(false)

	return editableDialog
}

// OnClose sets the callback that is called when the
// dialog is closed
func (d *EditableDialog) OnClose(fn func()) {
	ui.WindowManager().BeginUpdate()
	defer ui.WindowManager().EndUpdate()
	d.onClose = fn
}

// Result returns what button closed the dialog.
// See DialogButton constants. It can equal DialogAlive
// that means that the dialog is still visible and a
// user still does not click any button
func (d *EditableDialog) Result() int {
	return d.result
}

// Value returns the number of the selected item or
// -1 if nothing is selected or the dialog is cancelled
func (d *EditableDialog) Value() int {
	return d.value
}

// EditResult returns the text from editfield
func (d *EditableDialog) EditResult() string {
	return d.edtResult
}
