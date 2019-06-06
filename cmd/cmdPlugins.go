package cmd

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandPlugins() {

	// --- Disable ---------------------------------------------
	FrmMainSettings.SetVisible(false)
	FrmDiscordSettings.SetVisible(false)

	// --- Enable ----------------------------------------------
	FrmPlugins.SetVisible(true)

	// --- Modify ----------------------------------------------
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnPlugins.SetTitle(TxtPlugins + TxtActive)

	// --- Activate --------------------------------------------
	ui.ActivateControl(WindowMain, Btn3)
	//Btn3.SetVisible(true)
}
