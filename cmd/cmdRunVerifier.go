package cmd

import (
	. "github.com/instance-id/GoUI/elements"
	. "github.com/instance-id/GoUI/text"
)

func CommandRunVerifier() {

	// --- Disable ---------------------------------------------
	FrmMainSettings.SetVisible(false)
	FrmMainSettings.SetActive(false)
	FrmDiscordSettings.SetVisible(false)
	FrmDiscordSettings.SetActive(false)
	FrmDatabaseSettings.SetVisible(false)
	FrmDatabaseSettings.SetActive(false)
	FrmPlugins.SetVisible(false)
	FrmPlugins.SetActive(false)

	// --- Enable ----------------------------------------------
	FrmVerifier.SetVisible(true)
	FrmVerifier.SetActive(true)

	// --- Modify ----------------------------------------------
	BtnRunVerifier.SetTitle(TxtRunVerifier + TxtActive)
	BtnMainSettings.SetTitle(TxtMainSettings)
	BtnDiscordSettings.SetTitle(TxtDiscordSettings)
	BtnDatabaseSettings.SetTitle(TxtDatabaseSettings)
	BtnPlugins.SetTitle(TxtPlugins)

	// --- Activate --------------------------------------------

	// --- Buttons ---------------------------------------------
	BtnVerifierStart.SetActive(false)
	BtnVerifierRestart.SetActive(false)
	BtnVerifierStop.SetActive(false)
}
