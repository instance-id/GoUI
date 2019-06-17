package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandPlugins() {

	// --- Disable ---------------------------------------------
	FrmVerifier.SetVisible(false)
	FrmVerifier.SetActive(false)
	FrmMainSettings.SetVisible(false)
	FrmMainSettings.SetActive(false)
	FrmDiscordSettings.SetVisible(false)
	FrmDiscordSettings.SetActive(false)
	FrmDatabaseSettings.SetVisible(false)
	FrmDatabaseSettings.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmPlugins.SetVisible(true)
	FrmPlugins.SetActive(true)

	// --- Modify ----------------------------------------------
	BtnRunVerifier.SetTitle(TxtRunVerifier)
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnDatabaseSettings.SetTitle(TxtDatabaseSettings)
	BtnPlugins.SetTitle(TxtPlugins + TxtActive)

	// --- Activate --------------------------------------------

}
