package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandMainSettings() {

	// --- Disable ---------------------------------------------
	FrmDiscordSettings.SetVisible(false)
	FrmDiscordSettings.SetActive(false)
	FrmPlugins.SetVisible(false)
	FrmPlugins.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmMainSettings.SetVisible(true)
	FrmMainSettings.SetActive(true)

	// --- Modify ----------------------------------------------
	BtnMainSettings.SetTitle(TxtMainSettings + TxtActive)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------
}
