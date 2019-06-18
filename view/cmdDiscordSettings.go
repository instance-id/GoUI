package view

import (
	. "github.com/instance-id/GoUI/text"
)

func CommandDiscordSettings() {

	// --- Disable ---------------------------------------------
	FrmVerifier.SetVisible(false)
	FrmVerifier.SetActive(false)
	FrmMainSettings.SetVisible(false)
	FrmMainSettings.SetActive(false)
	FrmDatabaseSettings.SetVisible(false)
	FrmDatabaseSettings.SetActive(false)
	FrmPlugins.SetVisible(false)
	FrmPlugins.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmDiscordSettings.SetVisible(true)
	FrmDiscordSettings.SetActive(true)

	// --- Modify ----------------------------------------------
	BtnRunVerifier.SetTitle(TxtRunVerifier)
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings + TxtActive)
	BtnDatabaseSettings.SetTitle(TxtDatabaseSettings)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------
	//FrmDiscordSettings.SetActive(true)

	// --- Buttons ---------------------------------------------

}
