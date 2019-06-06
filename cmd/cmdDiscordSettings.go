package cmd

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandDiscordSettings() {

	// --- Disable ---------------------------------------------
	FrmMainSettings.SetVisible(false)
	FrmPlugins.SetVisible(false)

	// --- Enable ----------------------------------------------
	FrmDiscordSettings.SetVisible(true)

	// --- Modify ----------------------------------------------
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings + TxtActive)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------
	ui.ActivateControl(WindowMain, Btn2)
	//Btn2.SetVisible(true)
}
