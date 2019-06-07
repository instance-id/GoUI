package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandPlugins() {

	// --- Disable ---------------------------------------------
	FrmMainSettings.SetVisible(false)
	FrmMainSettings.SetActive(false)
	FrmDiscordSettings.SetVisible(false)
	FrmDiscordSettings.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmPlugins.SetVisible(true)

	// --- Modify ----------------------------------------------
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnPlugins.SetTitle(TxtPlugins + TxtActive)

	// --- Activate --------------------------------------------
	FrmPlugins.SetActive(true)
}
