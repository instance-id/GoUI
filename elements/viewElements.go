package elements

import (
	ui "github.com/instance-id/clui"
)

// --- Application Elements ----------------------------------------------
var (
	WindowMain *ui.Window

	FrameMenu           *ui.Frame
	FrameMainMenu       *ui.Frame
	FrameContent        *ui.Frame
	FrmVerifier         *ui.Frame
	FrmMainSettings     *ui.Frame
	FrmDiscordSettings  *ui.Frame
	FrmDatabaseSettings *ui.Frame
	FrmPlugins          *ui.Frame

	// --- Menu buttons -------------------------
	BtnRunVerifier      *ui.Button
	BtnMainSettings     *ui.Button
	BtnDiscordSettings  *ui.Button
	BtnDatabaseSettings *ui.Button
	BtnPlugins          *ui.Button
	BtnLogs             *ui.Button
	BtnQuit             *ui.Button

	// --- Verifier control buttons -------------
	BtnVerifierStart   *ui.Button
	BtnVerifierRestart *ui.Button
	BtnVerifierStop    *ui.Button

	// --- Main Settings buttons ----------------
	BtnLogLevel         *ui.Button
	BtnMainSettingsSave *ui.Button

	// --- Discord settings buttons -------------
	BtnAssetDetails        *ui.Button
	BtnDiscordSettingsSave *ui.Button

	// --- Database Settings buttons ------------
	BtnDatabaseProvider     *ui.Button
	BtnDatabaseSettingsSave *ui.Button

	// --- Plugins buttons ----------------------
	BtnPluginsSave *ui.Button

	// --- Arbitrary ------------------------------------------------
)
