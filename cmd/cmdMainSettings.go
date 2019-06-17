package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandMainSettings() {

	// --- Disable ---------------------------------------------
	FrmVerifier.SetVisible(false)
	FrmVerifier.SetActive(false)
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
	BtnRunVerifier.SetTitle(TxtRunVerifier)
	BtnMainSettings.SetTitle(TxtMainSettings + TxtActive)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnDatabaseSettings.SetTitle(TxtDatabaseSettings)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------

	// --- Buttons ---------------------------------------------
	BtnLogLevel.SetActive(false)
	BtnMainSettingsSave.SetActive(false)
}
