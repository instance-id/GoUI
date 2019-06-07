package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CreateViewPopupAssetCodes() {
	WindowAssetCodes = ui.AddWindow(10, 5, 1, 1, TxtAssetCodes)
	WindowAssetCodes.SetGaps(1, ui.KeepValue)
	WindowAssetCodes.SetVisible(false)
}

func ChangeAssetCodes(btn *ui.Button) {

	WindowAssetCodes = ui.AddWindow(3, 3, 50, 10, TxtAssetCodes)
	WindowAssetCodes.SetGaps(1, ui.KeepValue)
	WindowAssetCodes.Modal()
	ui.CreateTableView(WindowAssetCodes, ui.AutoSize, ui.AutoSize, 1)

	BtnTheme = ui.CreateButton(WindowAssetCodes, 15, 10, "Close", 1)
	BtnTheme.OnClick(func(ev ui.Event) {
		ui.WindowManager().DestroyWindow(WindowAssetCodes)
		btn.SetEnabled(true)

	})
}
