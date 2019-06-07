package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandDiscordSettings() {

	// --- Disable ---------------------------------------------
	FrmMainSettings.SetVisible(false)
	FrmMainSettings.SetActive(false)
	FrmPlugins.SetVisible(false)
	FrmPlugins.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmDiscordSettings.SetVisible(true)
	//FrmDiscordSettings.SetActive(true)

	// --- Modify ----------------------------------------------
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings + TxtActive)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------
	//FrmDiscordSettings.SetActive(true)

}
