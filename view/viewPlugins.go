package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
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

	// --- Setup Complete ------------------------------------------------
	FrmPlugins.SetVisible(false)
}
