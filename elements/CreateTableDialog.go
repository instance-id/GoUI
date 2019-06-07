package elements

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/text"
)

func CreateTableDialog(btn *ui.Button) {
	WindowAssetCodes = ui.AddWindow(3, 3, 50, 10, TxtAssetCodes)
	bch := ui.CreateTableView(WindowAssetCodes, 25, 12, 1)
	ui.ActivateControl(WindowAssetCodes, bch)
	WindowAssetCodes.SetGaps(1, ui.KeepValue)
	WindowAssetCodes.SetModal(true)
	ui.CreateTableView(WindowAssetCodes, ui.AutoSize, ui.AutoSize, 1)

	BtnTheme = ui.CreateButton(WindowAssetCodes, 15, 10, "Close", 1)
	BtnTheme.OnClick(func(ev ui.Event) {
		ui.WindowManager().DestroyWindow(WindowAssetCodes)
		btn.SetEnabled(true)

	})
}
