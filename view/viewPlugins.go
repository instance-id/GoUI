package view

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
	. "github.com/instance-id/GoUI/utils"
	ui "github.com/instance-id/clui"
)

var pendingPlugins = struct {
	tmp1 bool
	tmp2 bool
}{
	tmp1: false,
	tmp2: false,
}

func CreateViewPlugins() {

	// --- Plugins Frame -------------------------------------------------
	FrmPlugins = ui.CreateFrame(FrameContent, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.Fixed)
	FrmPlugins.SetTitle("FrameTop")
	FrmPlugins.SetPack(ui.Horizontal)

	// --- Plugins Content -----------------------------------------------
	pluginsFrame := ui.CreateFrame(FrmPlugins, 130, ui.AutoSize, ui.BorderThin, ui.AutoSize)
	pluginsFrame.SetPaddings(2, 2)
	pluginsFrame.SetTitle(TxtPlugins)
	pluginsFrame.SetPack(ui.Vertical)

	ui.CreateLabel(pluginsFrame, ui.AutoSize, ui.AutoSize, TxtPluginsLabel, ui.Fixed)

	// --- Window Control ------------------------------------------------
	btnFrame := ui.CreateFrame(pluginsFrame, 10, 1, ui.BorderNone, ui.Fixed)
	btnFrame.SetPaddings(2, 2)
	var params = FramedInputParams{Orientation: ui.Vertical, Width: 25, Height: 4, Scale: ui.Fixed}
	saveSettings := NewFramedInput(btnFrame, TxtSaveDesc, &params)
	BtnPluginsSave = ui.CreateButton(saveSettings, ui.AutoSize, ui.AutoSize, TxtSaveBtn, ui.Fixed)
	BtnPluginsSave.SetAlign(ui.AlignLeft)
	BtnPluginsSave.SetShadowType(ui.ShadowHalf)
	BtnPluginsSave.OnClick(func(ev ui.Event) {

	})

	// --- Setup Complete ------------------------------------------------
	FrmPlugins.SetVisible(false)
}

func AnimateButton() {
	aData := new(AnimateData)
	aData.Button = BtnPluginsSave
	aData.AnimateBtnNShdw(3000, 100, 10, 0)

}

func SavePendingPlugins() {
	if pendingPlugins.tmp1 || pendingPlugins.tmp2 == true {
		BtnPluginsSave.SetTitle(TxtSavePendingBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	} else {
		BtnPluginsSave.SetTitle(TxtSaveBtn)
		ui.PutEvent(ui.Event{Type: ui.EventRedraw})
	}
}
