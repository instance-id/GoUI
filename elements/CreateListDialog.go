package elements

//
//import (
//	"fmt"
//	. "github.com/instance-id/GoUI/text"
//	. "github.com/instance-id/GoUI/utils"
//	ui "github.com/instance-id/clui"
//	term "github.com/nsf/termbox-go"
//)
//
//// --- Populate list ------------------------------------------------
//func RefreshAssetList(list *EditDialog) {
//	list.list.Clear()
//
//	for k, _ := range *Asset {
//		list.list.AddItem(k)
//	}
//	ui.PutEvent(ui.Event{Type: ui.EventRedraw})
//
//}
//
//func CreateListDialog(listTitle string) *EditDialog {
//	listDialog := new(EditDialog)
//
//	cw, ch := term.Size()
//
//	listDialog.View = ui.AddWindow(cw/2-24, ch/2-16, ui.AutoSize, ui.AutoSize, TxtAssetCodes)
//	ui.WindowManager().BeginUpdate()
//	defer ui.WindowManager().EndUpdate()
//	listDialog.View.SetGaps(1, ui.KeepValue)
//	listDialog.View.SetModal(true)
//	listDialog.View.SetPack(ui.Vertical)
//
//	listDialog.Frame = NewFramedWindowInput(listDialog.View, listTitle, nil)
//	listDialog.list = ui.CreateListBox(listDialog.Frame, 25, 12, 1)
//	listDialog.list.SetBackColor(term.ColorBlack)
//	listDialog.list.SetTextColor(term.ColorWhite)
//	ui.ActivateControl(listDialog.Frame, listDialog.list)
//	RefreshAssetList(listDialog)
//
//	// --- Buttons -------------------------------------------------------
//	// --- Add new asset to data map ----------------------
//	frm1 := ui.CreateFrame(listDialog.View, 16, 4, ui.BorderNone, ui.Fixed)
//	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, 1)
//	btn0 := ui.CreateButton(frm1, ui.AutoSize, ui.AutoSize, "Add New", ui.Fixed)
//	btn0.OnClick(func(ev ui.Event) {
//		listDialog.result = DialogButton1
//		listDialog.editDialog = ui.CreateEditDialog(fmt.Sprintf("Add New"), fmt.Sprintf("Type new Asset"), "")
//		listDialog.editDialog.View.SetSize(30, 15)
//		listDialog.editDialog.View.SetScale(1)
//		listDialog.editDialog.OnClose(func() {
//			switch listDialog.editDialog.Result() {
//			case ui.DialogButton1:
//				RefreshAssetList(listDialog)
//				btn0.SetActive(false)
//				ui.PutEvent(ui.Event{Type: ui.EventRedraw})
//			case ui.DialogButton2:
//				btn0.SetActive(false)
//
//			}
//		})
//	})
//
//	// --- Edit current asset key in map ------------------
//	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, 1)
//	btn1 := ui.CreateButton(frm1, ui.AutoSize, ui.AutoSize, "Edit", ui.Fixed)
//	btn1.OnClick(func(ev ui.Event) {
//		listDialog.result = DialogButton2
//		listDialog.editDialog = ui.CreateEditDialog("Edit Asset", "Edit Asset", listDialog.list.SelectedItemText())
//		listDialog.editDialog.OnClose(func() {
//			switch listDialog.editDialog.Result() {
//			case ui.DialogButton1:
//				newText := listDialog.editDialog.EditResult()
//				tmpVal := AssetData.AssetPackages[listDialog.list.SelectedItemText()]
//				delete(AssetData.AssetPackages, listDialog.list.SelectedItemText())
//				RefreshAssetList(listDialog)
//				AssetData.AssetPackages[newText] = tmpVal
//				RefreshAssetList(listDialog)
//				btn1.SetActive(false)
//				ui.PutEvent(ui.Event{Type: ui.EventRedraw})
//			case ui.DialogButton2:
//				btn1.SetActive(false)
//			}
//
//		})
//	})
//
//	// --- Delete asset key from map ----------------------
//	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, 1)
//	btn2 := ui.CreateButton(frm1, ui.AutoSize, ui.AutoSize, "Remove", ui.Fixed)
//	btn2.OnClick(func(ev ui.Event) {
//		listDialog.result = DialogButton3
//		btns := []string{"Remove", "Cancel"}
//		item := listDialog.list.SelectedItemText()
//		listDialog.confirm = ui.CreateConfirmationDialog("Remove Item", fmt.Sprintf("Would you like to remove %s?", item), btns, 1)
//		listDialog.confirm.OnClose(func() {
//			switch listDialog.confirm.Result() {
//			case 0:
//				break
//			case 1:
//				delete(AssetData.AssetPackages, listDialog.list.SelectedItemText())
//				RefreshAssetList(listDialog)
//				btn2.SetActive(false)
//			}
//			listDialog.result = -1
//		})
//
//		if listDialog.onClose != nil {
//			listDialog.onClose()
//		}
//	})
//
//	// --- Close the dialog -------------------------------
//	ui.CreateFrame(frm1, 1, 1, ui.BorderNone, 1)
//	btn3 := ui.CreateButton(frm1, ui.AutoSize, ui.AutoSize, "Close", ui.Fixed)
//	btn3.OnClick(func(ev ui.Event) {
//		listDialog.result = DialogButton3
//		listDialog.edtResult = ""
//		listDialog.value = -1
//		ui.WindowManager().DestroyWindow(listDialog.View)
//		btn3.SetActive(false)
//		if listDialog.onClose != nil {
//			listDialog.onClose()
//		}
//	})
//
//	listDialog.View.OnClose(func(ev ui.Event) bool {
//		if listDialog.result == DialogAlive {
//			listDialog.result = DialogClosed
//			if ev.X != 1 {
//				ui.WindowManager().DestroyWindow(listDialog.View)
//			}
//			if listDialog.onClose != nil {
//				listDialog.onClose()
//			}
//		}
//
//		return true
//	})
//	BtnAssetCodes.SetActive(false)
//
//	return listDialog
//}
//
//// OnClose sets the callback that is called when the
//// dialog is closed
//func (d *EditDialog) OnClose(fn func()) {
//	ui.WindowManager().BeginUpdate()
//	defer ui.WindowManager().EndUpdate()
//	d.onClose = fn
//}
//
//// Result returns what button closed the dialog.
//// See DialogButton constants. It can equal DialogAlive
//// that means that the dialog is still visible and a
//// user still does not click any button
//func (d *EditDialog) Result() int {
//	return d.result
//}
//
//// Value returns the number of the selected item or
//// -1 if nothing is selected or the dialog is cancelled
//func (d *EditDialog) Value() int {
//	return d.value
//}
//
//// EditResult returns the text from editfield
//func (d *EditDialog) EditResult() string {
//	return d.edtResult
//}
