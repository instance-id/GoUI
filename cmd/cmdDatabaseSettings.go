package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandDatabaseSettings() {

	// --- Disable ---------------------------------------------
	FrmVerifier.SetVisible(false)
	FrmVerifier.SetActive(false)
	FrmMainSettings.SetVisible(false)
	FrmMainSettings.SetActive(false)
	FrmDiscordSettings.SetVisible(false)
	FrmDiscordSettings.SetActive(false)
	FrmPlugins.SetVisible(false)
	FrmPlugins.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmDatabaseSettings.SetVisible(true)
	FrmDatabaseSettings.SetActive(true)

	// --- Modify ----------------------------------------------
	BtnRunVerifier.SetTitle(TxtRunVerifier)
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnDatabaseSettings.SetTitle(TxtDatabaseSettings + TxtActive)
	BtnDatabaseSettings.SetActive(true)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------

}
