package view

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

func CreateViewPlugins() {

	// --- Plugins Frame -------------------------------------------------
	FrmPlugins = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmPlugins.SetTitle("FrameTop")
	FrmPlugins.SetPack(ui.Horizontal)

	// --- Plugins Content -----------------------------------------------
	pluginsFrame := ui.CreateFrame(FrmPlugins, 100, ui.AutoSize, ui.BorderThin, ui.AutoSize)
	pluginsFrame.SetPaddings(2, 2)
	pluginsFrame.SetTitle(TxtPlugins)
	pluginsFrame.SetPack(ui.Vertical)

	// --- Window Control ------------------------------------------------
	btnFrame := ui.CreateFrame(pluginsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)

	var params = FramedInputParams{Orientation: ui.Vertical, Width: 190, Height: 4, Scale: ui.Fixed}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	BtnMainSettingsSave = ui.CreateButton_NoShadow(saveSettings, ui.AutoSize, ui.AutoSize, TxtSave, ui.Fixed)
	BtnMainSettingsSave.SetAlign(ui.AlignLeft)

	BtnMainSettingsSave.OnClick(func(ev ui.Event) {

	})

	// --- Setup Complete ------------------------------------------------
	FrmPlugins.SetVisible(false)
}

func AnimateButton() {
	aData := new(AnimateData)
	aData.Button = BtnMainSettingsSave
	aData.AnimateBtnNShdw(3000, 100, 10, 0)

}
