package view

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CreateViewPlugins() {

	// --- Plugins Frame -------------------------------------------------
	FrmPlugins = ui.CreateFrame(WindowMain, ui.AutoSize, ui.AutoSize, ui.BorderNone, ui.AutoSize)
	FrmPlugins.SetTitle("FrameTop")
	FrmPlugins.SetPack(ui.Horizontal)

	// --- Plugins Content -----------------------------------------------
	Btn3 = ui.CreateButton(FrmPlugins, ui.AutoSize, ui.AutoSize, TxtPlugins, 1)
}
