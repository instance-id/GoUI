package cmd

import (
	ui "github.com/VladimirMarkelov/clui"
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandMainSettings() {

	// --- Disable ---------------------------------------------
	FrmDiscordSettings.SetVisible(false)
	FrmPlugins.SetVisible(false)

	// --- Enable ----------------------------------------------
	FrmMainSettings.SetVisible(true)

	// --- Modify ----------------------------------------------
	BtnMainSettings.SetTitle(TxtMainSettings + TxtActive)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------
	ui.ActivateControl(WindowMain, Btn1)
}
