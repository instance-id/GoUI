package elements

//
//import (
//	. "github.com/VladimirMarkelov/clui"
//	term "github.com/nsf/termbox-go"
//)
//
//// SelectDialog allows to user to select an item from
//// the list. Items can be displayed in ListBox or in
//// RadioGroup.
//// The dialog is modal, so a user cannot interact other
//// Views until the user closes the dialog
//type TableWindow struct {
//	View      *Window
//	result    int
//	value     int
//	edtResult string
//	rg        *RadioGroup
//	list      *ListBox
//	edit      *EditField
//	typ       SelectDialogType
//	onClose   func()
//}
//
//// ------------------------ Selection Dialog ---------------------
//
//// NewSelectDialog creates new dialog to select an item from list.
//// c is a composer that manages the dialog
//// title is a dialog title
//// items is a list of items to select from
//// selectedItem is the index of the item that is selected after
////  the dialog is created
//// typ is a selection type: ListBox or RadioGroup
//// Returns nil in case of creation process fails, e.g, if item list is empty
//func CreateTableWindow(title string, items []string, selectedItem int, typ SelectDialogType) *SelectDialog {
//	dlg := new(TableWindow)
//
//	if len(items) == 0 {
//		// Item list must contain at least 1 item
//		return nil
//	}
//
//	cw, ch := term.Size()
//
//	dlg.typ = typ
//	dlg.View = AddWindow(cw/2-12, ch/2-8, 20, 10, title)
//	WindowManager().BeginUpdate()
//	defer WindowManager().EndUpdate()
//	dlg.View.SetModal(true)
//	dlg.View.SetPack(Vertical)
//
//	if typ == SelectDialogList {
//		fList := CreateFrame(dlg.View, 1, 1, BorderNone, 1)
//		fList.SetPaddings(1, 1)
//		fList.SetGaps(0, 0)
//		dlg.list = CreateListBox(fList, 15, 5, 1)
//		for _, item := range items {
//			dlg.list.AddItem(item)
//		}
//		if selectedItem >= 0 && selectedItem < len(items) {
//			dlg.list.SelectItem(selectedItem)
//		}
//	} else if typ == SelectDialogEdit {
//		CreateFrame(dlg.View, 1, 1, BorderNone, Fixed)
//		lb := CreateLabel(dlg.View, 10, 3, items[0], 1)
//		lb.SetMultiline(true)
//
//		fWidth, _ := dlg.View.Size()
//		dlg.edit = CreateEditField(dlg.View, fWidth-2, items[1], AutoSize)
//		CreateFrame(dlg.View, 1, 1, BorderNone, Fixed)
//
//		dlg.edit.OnKeyPress(func(key term.Key, r rune) bool {
//			var input string
//			if key == term.KeyEnter {
//				input = dlg.edit.Title()
//				dlg.edtResult = input
//				dlg.value = -1
//				dlg.result = DialogButton1
//
//				WindowManager().DestroyWindow(dlg.View)
//				if dlg.onClose != nil {
//					dlg.onClose()
//				}
//			}
//			// returning false so that other keypresses work as usual
//			return false
//		})
//	} else {
//		fRadio := CreateFrame(dlg.View, 1, 1, BorderNone, Fixed)
//		fRadio.SetPaddings(1, 1)
//		fRadio.SetGaps(0, 0)
//		fRadio.SetPack(Vertical)
//		dlg.rg = CreateRadioGroup()
//		for _, item := range items {
//			r := CreateRadio(fRadio, AutoSize, item, Fixed)
//			dlg.rg.AddItem(r)
//		}
//		if selectedItem >= 0 && selectedItem < len(items) {
//			dlg.rg.SetSelected(selectedItem)
//		}
//	}
//
//	frm1 := CreateFrame(dlg.View, 16, 4, BorderNone, Fixed)
//	CreateFrame(frm1, 1, 1, BorderNone, 1)
//	btn1 := NewCreateButton(frm1, AutoSize, AutoSize, "OK", Fixed)
//	btn1.OnClick(func(ev Event) {
//		dlg.result = DialogButton1
//		if dlg.typ == SelectDialogList {
//			dlg.value = dlg.list.SelectedItem()
//		} else if dlg.typ == SelectDialogEdit {
//			dlg.edtResult = dlg.edit.Title()
//			dlg.value = -1
//		} else {
//			dlg.value = dlg.rg.Selected()
//		}
//		WindowManager().DestroyWindow(dlg.View)
//		if dlg.onClose != nil {
//			dlg.onClose()
//		}
//	})
//
//	CreateFrame(frm1, 1, 1, BorderNone, 1)
//	btn2 := NewCreateButton(frm1, AutoSize, AutoSize, "Cancel", Fixed)
//	btn2.OnClick(func(ev Event) {
//		dlg.result = DialogButton2
//		dlg.edtResult = ""
//		dlg.value = -1
//		WindowManager().DestroyWindow(dlg.View)
//		if dlg.onClose != nil {
//			dlg.onClose()
//		}
//	})
//
//	if typ == SelectDialogEdit {
//		ActivateControl(dlg.View, dlg.edit)
//	} else {
//		ActivateControl(dlg.View, btn2)
//	}
//	CreateFrame(frm1, 1, 1, BorderNone, 1)
//
//	dlg.View.OnClose(func(ev Event) bool {
//		if dlg.result == DialogAlive {
//			dlg.result = DialogClosed
//			if ev.X != 1 {
//				WindowManager().DestroyWindow(dlg.View)
//			}
//			if dlg.onClose != nil {
//				dlg.onClose()
//			}
//		}
//
//		return true
//	})
//
//	return dlg
//}
//
//// OnClose sets the callback that is called when the
//// dialog is closed
//func (d *SelectDialog) OnClose(fn func()) {
//	WindowManager().BeginUpdate()
//	defer WindowManager().EndUpdate()
//	d.onClose = fn
//}
//
//// Result returns what button closed the dialog.
//// See DialogButton constants. It can equal DialogAlive
//// that means that the dialog is still visible and a
//// user still does not click any button
//func (d *SelectDialog) Result() int {
//	return d.result
//}
//
//// Value returns the number of the selected item or
//// -1 if nothing is selected or the dialog is cancelled
//func (d *SelectDialog) Value() int {
//	return d.value
//}
//
//// EditResult returns the text from editfield
//func (d *SelectDialog) EditResult() string {
//	return d.edtResult
//}
