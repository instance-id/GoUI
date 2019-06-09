package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandMainSettings() {

	// --- Disable ---------------------------------------------
	FrmDiscordSettings.SetVisible(false)
	FrmDiscordSettings.SetActive(false)
	FrmDatabaseSettings.SetVisible(false)
	FrmDatabaseSettings.SetActive(false)
	FrmPlugins.SetVisible(false)
	FrmPlugins.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmMainSettings.SetVisible(true)
	FrmMainSettings.SetActive(true)

	// --- Modify ----------------------------------------------
	BtnMainSettings.SetTitle(TxtMainSettings + TxtActive)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnDatabaseSettings.SetTitle(TxtDatabaseSettings)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------
}
